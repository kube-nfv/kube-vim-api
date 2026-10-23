# hack/openapi

Post-processor for the OpenAPI 3.1 documents that `protoc-gen-openapiv3` (grpc-gateway) writes to `bin/openapi-raw`. `make generate` runs it and writes the result to `generated/openapi`.

```bash
go run ./hack/openapi -in bin/openapi-raw -out generated/openapi [-version 1.0.0]
```

## What it does

1. **SOL 013 responses.** The generator has no annotation for responses, so the REST conventions of ETSI GS NFV-SOL 013 are added by method name:
   - `Create*` answers 201 with a `Location` header.
   - `Delete*` answers 204 with no body.
   - An operation that takes `nextpage_opaque_marker` gets the `Link` response header (clause 5.4.2). Its body is a plain array, set by `response_body` in the protos.
   - Every operation gets a `default` error response of type `application/problem+json` (`ProblemDetails`, SOL 013 clause 6, whose `type` names the `ErrorReason` from `common/v1/errors.proto`) and the `Version` request and response header (clause 9.4).
   - Each API gets `GET /<api>/api_versions` and `GET /<api>/v1/api_versions` (clause 9.3) returning `ApiVersionInformation`, with no `Version` header and a token without a scope. The API name and major version come from the document's paths.
2. **`oneof` without the guard.** The generator encodes a proto `oneof` with a "none set" guard that makes generated clients reject valid input.
   - A message that is only a `oneof` (`ImageSource`) becomes a `oneOf` of single-member variants (`ImageSourceHttp`, `ImageSourceRegistry`, ...) with `additionalProperties: false`, so clients accept exactly one member.
   - A `oneof` among other fields (`HttpSource` credentials next to `url`) stays as optional properties; the description says only one may be set, and the server enforces it.
3. **Client bundle.** One document per API is kept as published. All of them are also merged into `kubevim.openapi.json`, with package prefixes stripped from schema names (`kubenfv.kubevim.compute.v1.Compute` becomes `Compute`) and service prefixes from operation ids (`ComputeService_CreateCompute` becomes `CreateCompute`). The Python client is generated from the bundle.

Security schemes and scopes are not added here; they come from the `openapiv3_document` and `openapiv3_operation` annotations in the protos.

The schema only describes the REST behaviour. The kube-vim gateway is what returns these status codes, errors and headers.

## Tests

`go test ./hack/...` runs the tool on `testdata/raw`, real `protoc-gen-openapiv3` v2.32.0 output for a draft compute API and the image API, and checks each rule above, the bundle and deterministic output.

## When to change it

Add a rule here only when the generator cannot express something and the rule applies by convention to every API. A rule becomes unnecessary when grpc-gateway can express it (response annotations, a client-safe `oneof` encoding); remove it then.
