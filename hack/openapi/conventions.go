package main

import (
	"fmt"
	"strings"
)

var httpMethods = []string{"get", "put", "post", "delete", "patch", "head", "options", "trace"}

// applyConventions adds the ETSI GS NFV-SOL 013 behaviour protoc-gen-openapiv3
// cannot express: 201 with Location on Create, 204 on Delete, problem+json
// errors (clause 6) and the Version header (clause 9.4) on every operation.
func applyConventions(doc map[string]any) {
	comps := child(doc, "components")
	child(comps, "schemas")["ProblemDetails"] = problemDetailsSchema()
	child(comps, "parameters")["Version"] = map[string]any{
		"name":        "Version",
		"in":          "header",
		"required":    true,
		"description": "API version requested by the consumer (SOL 013 9.4).",
		"schema":      map[string]any{"type": "string"},
	}
	for _, path := range sortedKeys(child(doc, "paths")) {
		item := child(doc, "paths")[path].(map[string]any)
		for _, method := range httpMethods {
			if op, ok := item[method].(map[string]any); ok {
				applyOperation(op)
			}
		}
	}
}

func applyOperation(op map[string]any) {
	responses := child(op, "responses")
	ok, found := responses["200"].(map[string]any)
	if !found {
		ok = map[string]any{"description": "OK"}
	}
	delete(responses, "200")
	operationID, _ := op["operationId"].(string)
	switch method := operationID[strings.LastIndex(operationID, "_")+1:]; {
	case strings.HasPrefix(method, "Create"):
		child(ok, "headers")["Location"] = map[string]any{
			"description": "URI of the created resource.",
			"schema":      map[string]any{"type": "string", "format": "uri"},
		}
		responses["201"] = ok
	case strings.HasPrefix(method, "Delete"):
		responses["204"] = map[string]any{"description": "No Content"}
	default:
		responses["200"] = ok
	}
	responses["default"] = map[string]any{
		"description": "Error (SOL 013 6).",
		"content": map[string]any{
			"application/problem+json": map[string]any{
				"schema": map[string]any{"$ref": "#/components/schemas/ProblemDetails"},
			},
		},
	}
	for _, code := range sortedKeys(responses) {
		child(responses[code].(map[string]any), "headers")["Version"] = map[string]any{
			"description": "API version used (SOL 013 9.4).",
			"schema":      map[string]any{"type": "string"},
		}
	}
	params, _ := op["parameters"].([]any)
	op["parameters"] = append(params, map[string]any{"$ref": "#/components/parameters/Version"})
}

// problemDetailsSchema is ProblemDetails as SOL 013 6.3 profiles RFC 9457.
func problemDetailsSchema() map[string]any {
	return map[string]any{
		"type":        "object",
		"description": "Error details (SOL 013 6.3, RFC 9457).",
		"required":    []any{"status", "detail"},
		"properties": map[string]any{
			"type":     map[string]any{"type": "string", "format": "uri"},
			"title":    map[string]any{"type": "string"},
			"status":   map[string]any{"type": "integer"},
			"detail":   map[string]any{"type": "string"},
			"instance": map[string]any{"type": "string", "format": "uri"},
		},
	}
}

// rewriteOneofs replaces protoc-gen-openapiv3's guarded oneof encoding, which
// makes generated clients reject valid input, with a oneOf of named variants
// that each hold exactly one member.
func rewriteOneofs(doc map[string]any) error {
	schemas := child(child(doc, "components"), "schemas")
	for _, name := range sortedKeys(schemas) {
		s, _ := schemas[name].(map[string]any)
		alts, _ := s["oneOf"].([]any)
		props, _ := s["properties"].(map[string]any)
		if len(alts) < 2 || props == nil {
			continue
		}
		if guard, _ := alts[0].(map[string]any); guard["not"] == nil {
			continue
		}
		var variants []any
		for _, alt := range alts[1:] {
			req, _ := alt.(map[string]any)["required"].([]any)
			if len(req) != 1 {
				return fmt.Errorf("schema %s: unexpected oneof alternative %v", name, alt)
			}
			member := req[0].(string)
			variant := name + strings.ToUpper(member[:1]) + member[1:]
			if _, taken := schemas[variant]; taken {
				return fmt.Errorf("schema %s: variant name %s already used", name, variant)
			}
			schemas[variant] = map[string]any{
				"type":                 "object",
				"properties":           map[string]any{member: props[member]},
				"required":             []any{member},
				"additionalProperties": false,
			}
			variants = append(variants, map[string]any{"$ref": "#/components/schemas/" + variant})
		}
		replaced := map[string]any{"oneOf": variants}
		if d, ok := s["description"]; ok {
			replaced["description"] = d
		}
		schemas[name] = replaced
	}
	return nil
}
