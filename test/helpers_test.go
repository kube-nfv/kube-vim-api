// Package test pins the wire contract of the generated APIs: golden JSON for
// field names and enum values, and the protovalidate rules of each request.
package test

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"buf.build/go/protovalidate"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type golden = proto.Message

func readFile(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

// checkGolden decodes the file strictly into msg, encodes it again and requires
// the same JSON, so renamed fields or enum values fail.
func checkGolden(t *testing.T, path string, msg proto.Message) {
	t.Helper()
	data := readFile(t, path)
	if err := protojson.Unmarshal([]byte(data), msg); err != nil {
		t.Fatalf("decode: %v", err)
	}
	out, err := protojson.Marshal(msg)
	if err != nil {
		t.Fatalf("encode: %v", err)
	}
	var want, got any
	if err := json.Unmarshal([]byte(data), &want); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(out, &got); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("round trip differs:\nwant %s\n got %s", data, out)
	}
}

func checkValidation(t *testing.T, data string, msg proto.Message, valid bool) {
	t.Helper()
	if err := protojson.Unmarshal([]byte(data), msg); err != nil {
		t.Fatalf("decode: %v", err)
	}
	err := protovalidate.Validate(msg)
	if valid && err != nil {
		t.Errorf("want valid, got %v", err)
	}
	if !valid && err == nil {
		t.Error("want a validation error, got none")
	}
	if !valid && err != nil {
		t.Logf("rejected: %v", err)
	}
}
