package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// testdata/raw holds protoc-gen-openapiv3 v2.32.0 output: a draft compute API
// (AIP methods, a Delete returning Empty, OAuth scopes) and the image API
// (a message that is only a oneof, and oneofs among other fields).

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
	doc := load(t, filepath.Join(process(t), "image/v1/image.openapi.json"))
	src := get(t, doc, "components", "schemas", "kubenfv.kubevim.image.v1.ImageSource").(map[string]any)
	if _, ok := src["properties"]; ok {
		t.Error("ImageSource keeps top-level properties")
	}
	var refs []string
	for _, v := range src["oneOf"].([]any) {
		refs = append(refs, v.(map[string]any)["$ref"].(string))
	}
	var want []string
	for _, m := range []string{"Http", "S3", "Gcs", "Registry", "Pvc", "Snapshot", "Upload", "Blank", "Imageio", "Vddk"} {
		want = append(want, "#/components/schemas/kubenfv.kubevim.image.v1.ImageSource"+m)
	}
	if got := strings.Join(refs, ","); got != strings.Join(want, ",") {
		t.Errorf("oneOf = %s, want %s", got, strings.Join(want, ","))
	}
	v := get(t, doc, "components", "schemas", "kubenfv.kubevim.image.v1.ImageSourceHttp").(map[string]any)
	if v["additionalProperties"] != false || v["required"].([]any)[0] != "http" {
		t.Errorf("variant ImageSourceHttp = %v", v)
	}
}

func TestOneofAmongFields(t *testing.T) {
	doc := load(t, filepath.Join(process(t), "image/v1/image.openapi.json"))
	s := get(t, doc, "components", "schemas", "kubenfv.kubevim.image.v1.HttpSource").(map[string]any)
	if _, ok := s["oneOf"]; ok {
		t.Error("HttpSource keeps the oneOf guard")
	}
	for _, p := range []string{"url", "caBundle", "basicAuth", "credentialsRef"} {
		get(t, s, "properties", p)
	}
	if d := s["description"].(string); !strings.HasSuffix(d, "Only one of basicAuth, credentialsRef may be set.") {
		t.Errorf("description = %q", d)
	}
}

func TestLinkHeader(t *testing.T) {
	doc := load(t, filepath.Join(process(t), "compute/v1/compute_service.openapi.json"))
	get(t, doc, "paths", "/compute/v1/computes", "get", "responses", "200", "headers", "Link")
	if _, ok := get(t, doc, "paths", "/compute/v1/computes/{compute_id}", "get", "responses", "200", "headers").(map[string]any)["Link"]; ok {
		t.Error("Get has a Link header; only paged Lists should")
	}
}

func TestApiVersions(t *testing.T) {
	out := process(t)
	doc := load(t, filepath.Join(out, "compute/v1/compute_service.openapi.json"))
	for path, id := range map[string]string{
		"/compute/api_versions":    "ComputeService_GetApiVersions",
		"/compute/v1/api_versions": "ComputeService_GetApiMajorVersions",
	} {
		op := get(t, doc, "paths", path, "get").(map[string]any)
		if op["operationId"] != id {
			t.Errorf("%s: operationId = %v, want %s", path, op["operationId"], id)
		}
		if op["parameters"] != nil {
			t.Errorf("%s: parameters = %v, want none (no Version header)", path, op["parameters"])
		}
		if ref := get(t, op, "responses", "200", "content", "application/json", "schema", "$ref"); ref != "#/components/schemas/ApiVersionInformation" {
			t.Errorf("%s: response schema = %v", path, ref)
		}
	}
	get(t, doc, "components", "schemas", "ApiVersionInformation", "properties", "apiVersions")
	bundled := load(t, filepath.Join(out, bundleFile))
	if id := get(t, bundled, "paths", "/image/api_versions", "get", "operationId"); id != "Image_GetApiVersions" {
		t.Errorf("bundled operationId = %v, want Image_GetApiVersions", id)
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

func TestOperationIDs(t *testing.T) {
	ops := []map[string]any{
		{"operationId": "ComputeService_CreateCompute"},
		{"operationId": "ComputeService_CreateAffinityGroup"},
		{"operationId": "NetworkService_CreateAffinityGroup"},
	}
	if err := shortenOperationIDs(ops); err != nil {
		t.Fatal(err)
	}
	var got []string
	for _, op := range ops {
		got = append(got, op["operationId"].(string))
	}
	if want := "CreateCompute,Compute_CreateAffinityGroup,Network_CreateAffinityGroup"; strings.Join(got, ",") != want {
		t.Errorf("operation ids = %v, want %s", got, want)
	}
}

func TestDeterministic(t *testing.T) {
	a, b := process(t), process(t)
	for _, rel := range []string{"compute/v1/compute_service.openapi.json", "image/v1/image.openapi.json", bundleFile} {
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
