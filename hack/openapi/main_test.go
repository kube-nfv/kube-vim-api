package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// testdata/raw holds protoc-gen-openapiv3 v2.32.0 output for draft compute and
// image APIs: AIP methods, a Delete returning Empty, a oneof and OAuth scopes.

func process(t *testing.T) string {
	t.Helper()
	out := t.TempDir()
	if err := run("testdata/raw", out, "1.0.0"); err != nil {
		t.Fatal(err)
	}
	return out
}

func load(t *testing.T, path string) map[string]any {
	t.Helper()
	doc, err := readJSON(path)
	if err != nil {
		t.Fatal(err)
	}
	return doc
}

func get(t *testing.T, node any, keys ...string) any {
	t.Helper()
	for _, k := range keys {
		m, ok := node.(map[string]any)
		if !ok {
			t.Fatalf("%s: not an object", strings.Join(keys, "."))
		}
		if node, ok = m[k]; !ok {
			t.Fatalf("missing %q in %s", k, strings.Join(keys, "."))
		}
	}
	return node
}

func responseCodes(t *testing.T, doc map[string]any, path, method string) string {
	t.Helper()
	return strings.Join(sortedKeys(get(t, doc, "paths", path, method, "responses").(map[string]any)), ",")
}

func TestConventions(t *testing.T) {
	doc := load(t, filepath.Join(process(t), "compute/v1/compute_service.openapi.json"))
	for _, tc := range []struct{ path, method, want string }{
		{"/compute/v1/computes", "post", "201,default"},
		{"/compute/v1/computes", "get", "200,default"},
		{"/compute/v1/computes/{compute_id}", "get", "200,default"},
		{"/compute/v1/computes/{compute_id}", "delete", "204,default"},
		{"/compute/v1/computes/{compute_id}/operate", "post", "200,default"},
	} {
		if got := responseCodes(t, doc, tc.path, tc.method); got != tc.want {
			t.Errorf("%s %s responses = %s, want %s", tc.method, tc.path, got, tc.want)
		}
		op := get(t, doc, "paths", tc.path, tc.method).(map[string]any)
		get(t, op, "responses", "default", "content", "application/problem+json")
		get(t, op, "responses", "default", "headers", "Version")
		params := op["parameters"].([]any)
		if ref := params[len(params)-1].(map[string]any)["$ref"]; ref != "#/components/parameters/Version" {
			t.Errorf("%s %s: last parameter is %v, want the Version header", tc.method, tc.path, ref)
		}
	}
	get(t, doc, "paths", "/compute/v1/computes", "post", "responses", "201", "headers", "Location")
	if _, ok := get(t, doc, "paths", "/compute/v1/computes/{compute_id}", "delete", "responses", "204").(map[string]any)["content"]; ok {
		t.Error("204 response has content")
	}
	get(t, doc, "components", "schemas", "ProblemDetails")
}

func TestOneofVariants(t *testing.T) {
	doc := load(t, filepath.Join(process(t), "image/v1/image_service.openapi.json"))
	src := get(t, doc, "components", "schemas", "kubenfv.kubevim.image.v1.ImageSource").(map[string]any)
	if _, ok := src["properties"]; ok {
		t.Error("ImageSource keeps top-level properties")
	}
	var refs []string
	for _, v := range src["oneOf"].([]any) {
		refs = append(refs, v.(map[string]any)["$ref"].(string))
	}
	want := "#/components/schemas/kubenfv.kubevim.image.v1.ImageSourceHttp," +
		"#/components/schemas/kubenfv.kubevim.image.v1.ImageSourceRegistry," +
		"#/components/schemas/kubenfv.kubevim.image.v1.ImageSourceUpload"
	if got := strings.Join(refs, ","); got != want {
		t.Errorf("oneOf = %s, want %s", got, want)
	}
	v := get(t, doc, "components", "schemas", "kubenfv.kubevim.image.v1.ImageSourceHttp").(map[string]any)
	if v["additionalProperties"] != false || v["required"].([]any)[0] != "http" {
		t.Errorf("variant ImageSourceHttp = %v", v)
	}
}

func TestBundle(t *testing.T) {
	out := process(t)
	doc := load(t, filepath.Join(out, bundleFile))
	if got := get(t, doc, "info", "version"); got != "1.0.0" {
		t.Errorf("info.version = %v", got)
	}
	for _, name := range []string{"Compute", "ImageSource", "ImageSourceHttp", "ProblemDetails"} {
		get(t, doc, "components", "schemas", name)
	}
	if id := get(t, doc, "paths", "/compute/v1/computes", "post", "operationId"); id != "CreateCompute" {
		t.Errorf("operationId = %v, want CreateCompute", id)
	}
	data, err := os.ReadFile(filepath.Join(out, bundleFile))
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(data, []byte("kubenfv.kubevim.")) {
		t.Error("bundle still contains package-qualified names")
	}
	scopes := get(t, doc, "components", "securitySchemes", "oauth2", "flows", "clientCredentials", "scopes").(map[string]any)
	if _, ok := scopes["compute:v1:allocate"]; !ok {
		t.Errorf("scopes = %v", scopes)
	}
}

func TestDeterministic(t *testing.T) {
	a, b := process(t), process(t)
	for _, rel := range []string{"compute/v1/compute_service.openapi.json", "image/v1/image_service.openapi.json", bundleFile} {
		x, _ := os.ReadFile(filepath.Join(a, rel))
		y, _ := os.ReadFile(filepath.Join(b, rel))
		if !bytes.Equal(x, y) {
			t.Errorf("%s differs between runs", rel)
		}
	}
}

func TestNoInput(t *testing.T) {
	out := t.TempDir()
	if err := run(filepath.Join(t.TempDir(), "missing"), out, "1.0.0"); err != nil {
		t.Fatal(err)
	}
	if entries, _ := os.ReadDir(out); len(entries) != 0 {
		t.Errorf("wrote %d entries for no input", len(entries))
	}
}
