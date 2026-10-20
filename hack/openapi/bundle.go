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

// bundle merges the per-API documents into one, with component names and
// operation ids stripped of their proto package and service prefixes.
func bundle(docs []map[string]any, version string) (map[string]any, error) {
	paths := map[string]any{}
	comps := map[string]any{}
	for _, doc := range docs {
		doc = shortenRefs(doc).(map[string]any)
		for _, path := range sortedKeys(child(doc, "paths")) {
			if _, dup := paths[path]; dup {
				return nil, fmt.Errorf("path %s defined twice", path)
			}
			item := child(doc, "paths")[path].(map[string]any)
			for _, method := range httpMethods {
				if op, ok := item[method].(map[string]any); ok {
					if id, ok := op["operationId"].(string); ok {
						op["operationId"] = id[strings.LastIndex(id, "_")+1:]
					}
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
	return map[string]any{
		"openapi":    docs[0]["openapi"],
		"info":       map[string]any{"title": "Kube-VIM API", "version": version},
		"paths":      paths,
		"components": comps,
	}, nil
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
