package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

// packagePrefix matches the proto package that protoc-gen-openapiv3 puts in
// every component name, e.g. "kubenfv.kubevim.compute.v1.".
var packagePrefix = regexp.MustCompile(`^kubenfv\.kubevim\.\w+\.v\d+\.`)

// bundle merges the per-API documents into one, with component names stripped
// of their proto package and operation ids of their service prefix.
func bundle(docs []map[string]any, version string) (map[string]any, error) {
	paths := map[string]any{}
	comps := map[string]any{}
	var ops []map[string]any
	for _, doc := range docs {
		doc = shortenRefs(doc).(map[string]any)
		for _, path := range sortedKeys(child(doc, "paths")) {
			if _, dup := paths[path]; dup {
				return nil, fmt.Errorf("path %s defined twice", path)
			}
			item := child(doc, "paths")[path].(map[string]any)
			for _, method := range httpMethods {
				if op, ok := item[method].(map[string]any); ok {
					ops = append(ops, op)
				}
			}
			paths[path] = item
		}
		for _, kind := range sortedKeys(child(doc, "components")) {
			src, _ := child(doc, "components")[kind].(map[string]any)
			dst := child(comps, kind)
			for _, name := range sortedKeys(src) {
				short := packagePrefix.ReplaceAllString(name, "")
				prev, exists := dst[short]
				switch {
				case !exists:
					dst[short] = src[name]
				case kind == "securitySchemes":
					if err := mergeScopes(prev, src[name]); err != nil {
						return nil, fmt.Errorf("security scheme %s: %w", short, err)
					}
				case !reflect.DeepEqual(prev, src[name]):
					return nil, fmt.Errorf("%s %s differs between APIs", kind, short)
				}
			}
		}
	}
	if err := shortenOperationIDs(ops); err != nil {
		return nil, err
	}
	// Clients set the Version header once, not per call; the server still requires it.
	if v, ok := child(comps, "parameters")["Version"].(map[string]any); ok {
		v["required"] = false
		v["description"] = v["description"].(string) + " Required by the server; set it once on the client, e.g. as a default header."
	}
	return map[string]any{
		"openapi":    docs[0]["openapi"],
		"info":       map[string]any{"title": "Kube-VIM API", "version": version},
		"paths":      paths,
		"components": comps,
	}, nil
}

// shortenOperationIDs turns "ComputeService_CreateCompute" into "CreateCompute".
// A method name that more than one service uses keeps the service stem
// ("Compute_CreateAffinityGroup"), since operation ids must be unique.
func shortenOperationIDs(ops []map[string]any) error {
	count := map[string]int{}
	for _, op := range ops {
		id, _ := op["operationId"].(string)
		count[id[strings.LastIndex(id, "_")+1:]]++
	}
	seen := map[string]bool{}
	for _, op := range ops {
		id, _ := op["operationId"].(string)
		service, method, _ := strings.Cut(id, "_")
		short := method
		if count[method] > 1 {
			short = strings.TrimSuffix(service, "Service") + "_" + method
		}
		if seen[short] {
			return fmt.Errorf("operation id %s is not unique", short)
		}
		seen[short] = true
		op["operationId"] = short
	}
	return nil
}

// shortenRefs returns a copy of node with package prefixes stripped from $ref targets.
func shortenRefs(node any) any {
	switch n := node.(type) {
	case map[string]any:
		out := make(map[string]any, len(n))
		for k, v := range n {
			if ref, ok := v.(string); ok && k == "$ref" {
				i := strings.LastIndex(ref, "/") + 1
				out[k] = ref[:i] + packagePrefix.ReplaceAllString(ref[i:], "")
				continue
			}
			out[k] = shortenRefs(v)
		}
		return out
	case []any:
		out := make([]any, len(n))
		for i, v := range n {
			out[i] = shortenRefs(v)
		}
		return out
	}
	return node
}

// mergeScopes adds the OAuth 2.0 scopes of src to dst; the schemes must otherwise match.
func mergeScopes(dst, src any) error {
	dstScopes, err := clientCredentialScopes(dst)
	if err != nil {
		return err
	}
	srcScopes, err := clientCredentialScopes(src)
	if err != nil {
		return err
	}
	for k, v := range srcScopes {
		dstScopes[k] = v
	}
	return nil
}

func clientCredentialScopes(scheme any) (map[string]any, error) {
	s, _ := scheme.(map[string]any)
	flows, _ := s["flows"].(map[string]any)
	cc, _ := flows["clientCredentials"].(map[string]any)
	scopes, ok := cc["scopes"].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("not an OAuth 2.0 client credentials scheme")
	}
	return scopes, nil
}
