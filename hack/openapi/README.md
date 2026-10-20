# hack/openapi

Post-processor for the OpenAPI 3.1 documents that `protoc-gen-openapiv3` (grpc-gateway) writes to `bin/openapi-raw`. `make generate` runs it and writes the result to `generated/openapi`.

```bash
go run ./hack/openapi -in bin/openapi-raw -out generated/openapi [-version 1.0.0]
```

## What it does

1. **SOL 013 responses.** The generator has no annotation for responses, so the REST conventions of ETSI GS NFV-SOL 013 are added by method name:
   - `Create*` answers 201 with a `Location` header.
   - `Delete*` answers 204 with no body.
   - Every operation gets a `default` error response of type `application/problem+json` (`ProblemDetails`, SOL 013 clause 6) and the `Version` request and response header (clause 9.4).
2. **`oneof` as named variants.** The generator encodes a proto `oneof` with a "none set" guard that makes generated clients reject valid input. Each `oneof` becomes a `oneOf` of single-member variants (`ImageSourceHttp`, `ImageSourceRegistry`, ...) with `additionalProperties: false`, so clients accept exactly one member.
3. **Client bundle.** One document per API is kept as published. All of them are also merged into `kubevim.openapi.json`, with package prefixes stripped from schema names (`kubenfv.kubevim.compute.v1.Compute` becomes `Compute`) and service prefixes from operation ids (`ComputeService_CreateCompute` becomes `CreateCompute`). The Python client is generated from the bundle.

Security schemes and scopes are not added here; they come from the `openapiv3_document` and `openapiv3_operation` annotations in the protos.

The schema only describes the REST behaviour. The kube-vim gateway is what returns these status codes, errors and headers.

## Tests

`go test ./hack/...` runs the tool on `testdata/raw`, real `protoc-gen-openapiv3` v2.32.0 output for draft compute and image APIs, and checks each rule above, the bundle and deterministic output.

## When to change it

Add a rule here only when the generator cannot express something and the rule applies by convention to every API. A rule becomes unnecessary when grpc-gateway can express it (response annotations, a client-safe `oneof` encoding); remove it then.
