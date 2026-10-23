package main

import (
	"fmt"
	"strings"
)

var httpMethods = []string{"get", "put", "post", "delete", "patch", "head", "options", "trace"}

// applyConventions adds the ETSI GS NFV-SOL 013 behaviour protoc-gen-openapiv3
// cannot express: 201 with Location on Create, 204 on Delete, problem+json
// errors (clause 6), the Version header (clause 9.4) on every operation and the
// api_versions resources (clause 9.3).
func applyConventions(doc map[string]any) error {
	comps := child(doc, "components")
	child(comps, "schemas")["ProblemDetails"] = problemDetailsSchema()
	child(comps, "parameters")["Version"] = map[string]any{
		"name":        "Version",
		"in":          "header",
		"required":    true,
		"description": "API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted.",
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
	return addApiVersions(doc)
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
		if hasParameter(op, "nextpage_opaque_marker") {
			child(ok, "headers")["Link"] = map[string]any{
				"description": "URI of the next page, rel=\"next\", carrying nextpage_opaque_marker; absent on the last page (SOL 013 5.4.2).",
				"schema":      map[string]any{"type": "string"},
			}
		}
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
			"description": "API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4).",
			"schema":      map[string]any{"type": "string"},
		}
	}
	params, _ := op["parameters"].([]any)
	op["parameters"] = append(params, map[string]any{"$ref": "#/components/parameters/Version"})
}

// hasParameter reports whether op takes the named parameter.
func hasParameter(op map[string]any, name string) bool {
	params, _ := op["parameters"].([]any)
	for _, p := range params {
		if m, _ := p.(map[string]any); m["name"] == name {
			return true
		}
	}
	return false
}

// addApiVersions adds GET {apiName}/api_versions and {apiName}/{major}/api_versions
// (SOL 013 9.3), taking the API name, major version and service from the
// document's paths and operation ids. They need a token but no scope, and no
// Version header, since clients call them to find the version to request.
func addApiVersions(doc map[string]any) error {
	paths := child(doc, "paths")
	var apiName, major, service string
	var tags any
	for _, path := range sortedKeys(paths) {
		segs := strings.Split(strings.TrimPrefix(path, "/"), "/")
		if len(segs) < 2 {
			return fmt.Errorf("path %s has no API name and major version", path)
		}
		if apiName == "" {
			apiName, major = segs[0], segs[1]
		}
		if segs[0] != apiName || segs[1] != major {
			return fmt.Errorf("path %s is not under /%s/%s", path, apiName, major)
		}
		for _, method := range httpMethods {
			if op, ok := paths[path].(map[string]any)[method].(map[string]any); ok && service == "" {
				id, _ := op["operationId"].(string)
				service, _, _ = strings.Cut(id, "_")
				tags = op["tags"]
			}
		}
	}
	if service == "" {
		return nil
	}
	child(child(doc, "components"), "schemas")["ApiVersionInformation"] = apiVersionInformationSchema()
	for path, method := range map[string]string{
		"/" + apiName + "/api_versions":               "GetApiVersions",
		"/" + apiName + "/" + major + "/api_versions": "GetApiMajorVersions",
	} {
		paths[path] = map[string]any{"get": map[string]any{
			"operationId": service + "_" + method,
			"tags":        tags,
			"summary":     "API versions this VIM serves (SOL 013 9.3).",
			"security":    []any{map[string]any{"oauth2": []any{}}},
			"responses": map[string]any{
				"200": map[string]any{
					"description": "OK",
					"content": map[string]any{
						"application/json": map[string]any{
							"schema": map[string]any{"$ref": "#/components/schemas/ApiVersionInformation"},
						},
					},
				},
				"default": map[string]any{
					"description": "Error (SOL 013 6).",
					"content": map[string]any{
						"application/problem+json": map[string]any{
							"schema": map[string]any{"$ref": "#/components/schemas/ProblemDetails"},
						},
					},
				},
			},
		}}
	}
	return nil
}

// apiVersionInformationSchema is ApiVersionInformation (SOL 013 9.3.3).
func apiVersionInformationSchema() map[string]any {
	return map[string]any{
		"type":        "object",
		"description": "API versions served under a URI prefix (SOL 013 9.3.3).",
		"required":    []any{"uriPrefix", "apiVersions"},
		"properties": map[string]any{
			"uriPrefix": map[string]any{"type": "string", "description": "URI prefix of the API, e.g. /compute/v1."},
			"apiVersions": map[string]any{
				"type": "array",
				"items": map[string]any{
					"type":     "object",
					"required": []any{"version"},
					"properties": map[string]any{
						"version":        map[string]any{"type": "string", "description": "Version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0."},
						"isDeprecated":   map[string]any{"type": "boolean"},
						"retirementDate": map[string]any{"type": "string", "format": "date-time"},
					},
				},
			},
		},
	}
}

// problemDetailsSchema is ProblemDetails as SOL 013 6.3 profiles RFC 9457.
func problemDetailsSchema() map[string]any {
	return map[string]any{
		"type":        "object",
		"description": "Error details (SOL 013 6.3, RFC 9457).",
		"required":    []any{"status", "detail"},
		"properties": map[string]any{
			"type": map[string]any{
				"type":        "string",
				"format":      "uri",
				"description": "https://kubevim.kubenfv.io/problems/<reason>, the error reason (ErrorReason in the protos) in kebab case, e.g. .../problems/resource-in-use; about:blank for authentication and authorisation errors. Clients should accept unknown types.",
			},
			"title":    map[string]any{"type": "string"},
			"status":   map[string]any{"type": "integer"},
			"detail":   map[string]any{"type": "string"},
			"instance": map[string]any{"type": "string", "format": "uri"},
		},
	}
}

// rewriteOneofs replaces protoc-gen-openapiv3's guarded oneof encoding, which
// makes generated clients reject valid input. A message that is only a oneof
// becomes a oneOf of named variants that each hold exactly one member; a oneof
// among other fields stays as optional properties, with the exclusivity in the
// description and enforced by the server.
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
		var members []string
		for _, alt := range alts[1:] {
			req, _ := alt.(map[string]any)["required"].([]any)
			if len(req) != 1 {
				return fmt.Errorf("schema %s: unexpected oneof alternative %v", name, alt)
			}
			members = append(members, req[0].(string))
		}
		if len(props) > len(members) {
			delete(s, "oneOf")
			note := "Only one of " + strings.Join(members, ", ") + " may be set."
			if d, ok := s["description"].(string); ok && d != "" {
				note = d + " " + note
			}
			s["description"] = note
			continue
		}
		var variants []any
		for _, member := range members {
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
