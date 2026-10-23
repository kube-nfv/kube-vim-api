# Kube-VIM API

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/kube-nfv/kube-vim-api.svg)](https://pkg.go.dev/github.com/kube-nfv/kube-vim-api)

The API of [Kube-VIM](https://github.com/kube-nfv/kube-vim), a Kubernetes-native Virtualized Infrastructure Manager (VIM) for Network Function Virtualization (NFV).

The API is built on the ETSI NFV-MANO specifications and aligned with them: resources, attributes and operations come from the ETSI GS NFV-IFA 005 and IFA 006 information model (the Or-Vi and Vi-Vnfm reference points), and the REST interface follows the ETSI GS NFV-SOL 013 conventions. Kube-VIM extends it where cloud-native Kubernetes infrastructure needs more (placement groups, storage classes, provider networks, image import) and where it makes the API easier to use and maintain. [Kube-VIM API and ETSI NFV](docs/etsi-alignment.md) lists the extensions and what is left out.

The API is defined in Protocol Buffers and served over gRPC and, through a gateway, as a REST API with an OpenAPI 3.1 schema. This repository publishes the protos, the generated Go code, the OpenAPI schemas and a Python client.

> **Status:** v1 is in alpha (`v1.0.0-alpha.N` tags). It may still change in breaking ways until `v1.0.0`. The v0 API and its `kubevim-vivnfm-client` package are superseded.

## APIs

One API per ETSI management interface:

| API | REST prefix | gRPC service | Manages | ETSI interface |
| --- | --- | --- | --- | --- |
| Software image | `/image/v1` | `kubenfv.kubevim.image.v1.ImageService` | Images and their import | IFA 005 Software Image Management |
| Compute flavour | `/flavour/v1` | `kubenfv.kubevim.flavour.v1.FlavourService` | Flavours: CPU, memory, disks | IFA 006 Compute Flavour Management |
| Network | `/network/v1` | `kubenfv.kubevim.network.v1.NetworkService` | Networks and subnets | IFA 006 Network Resources Management |
| Compute | `/compute/v1` | `kubenfv.kubevim.compute.v1.ComputeService` | Virtual machines, their snapshots, placement groups | IFA 006 Compute Resources Management |
| Storage | `/storage/v1` | `kubenfv.kubevim.storage.v1.StorageService` | Volumes and their snapshots | IFA 006 Storage Resources Management |

Or-Vi (NFVO) and Vi-Vnfm (VNFM) clients use the same APIs; they differ only in the OAuth 2.0 scopes they are granted.

## How the API works

- **Methods.** Each resource has Create, Get, List, Update and Delete (IFA's Allocate, Query, Update and Terminate). Actions are task endpoints, e.g. `POST /compute/v1/computes/{compute_id}/operate` to start, stop or reboot. Each RPC names the IFA operation it implements.
- **Identifiers** are UUIDs assigned by Kube-VIM, stable for the life of the resource.
- **Create** returns `201 Created` with a `Location` header and the resource. **Delete** returns `204 No Content`.
- **Update** is a partial update (`PATCH`): fields left out keep their value. Every resource with an Update can change its name and `metadata`.
- **Metadata.** Every resource has one `metadata` map. Updates merge keys into it; keys under `kubevim.kubenfv.io/` are written by Kube-VIM.
- **Asynchronous work** (image import, VM start) is reported in the resource's state; poll Get until it is done.
- **List** returns a JSON array and accepts a SOL 013 `filter`, e.g. `filter=(eq,computeName,router-1)`. When results are paged, the `Link` header carries the URL of the next page.
- **Errors** are `application/problem+json` ProblemDetails (RFC 9457). The `type` names the reason, e.g. `https://kubevim.kubenfv.io/problems/resource-in-use`; the reasons are listed in [errors.proto](api/kubenfv/kubevim/common/v1/errors.proto). Over gRPC the same reason is in a `google.rpc.ErrorInfo` detail.
- **Authentication** is an OAuth 2.0 bearer token. Each operation needs one scope, named `<api>:v1:<permission>`, e.g. `compute:v1:allocate` or `image:v1:query`. An NFVO is typically granted every scope; a VNFM every scope except adding, updating and deleting images.
- **Validation.** Requests are validated against the rules in the protos ([protovalidate](https://protovalidate.com/)). A request field that Kube-VIM does not support yet is rejected, never ignored.

## Clients

| Client | Install | Source |
| --- | --- | --- |
| Go (gRPC and types) | `go get github.com/kube-nfv/kube-vim-api` | [pkg/apis](pkg/apis) |
| Python (REST) | `pip install kubevim-client` | [generated/python](generated/python) |
| OpenAPI 3.1 schemas | GitHub release assets | [generated/openapi](generated/openapi): one per API, plus `kubevim.openapi.json` with all five |

### Python

```python
import kubevim_client
from kubevim_client import ComputeServiceApi, CreateComputeRequest, CreateImageRequest, ImageServiceApi

config = kubevim_client.Configuration(host="https://kubevim.example.com", access_token=token)
with kubevim_client.ApiClient(config) as client:
    image = ImageServiceApi(client).create_image(CreateImageRequest.from_dict({
        "image": {
            "name": "ubuntu-24.04",
            "source": {"http": {"url": "https://cloud-images.ubuntu.com/noble/current/noble-server-cloudimg-amd64.img"}},
        },
    }))

    compute = ComputeServiceApi(client).create_compute(CreateComputeRequest.from_dict({
        "compute": {"computeName": "router-1", "flavourId": flavour_id, "imageId": image.image_id},
        "networkInterfaces": [{"name": "data0", "networkId": network_id}],
    }))
    print(compute.compute_id, compute.running_state)
```

Errors raise `kubevim_client.exceptions.ApiException`; its `status` is the HTTP status and its `body` the ProblemDetails. Enum values added by a newer server read as `UNKNOWN_DEFAULT_OPEN_API` instead of failing.

### Go

```go
import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	computev1 "github.com/kube-nfv/kube-vim-api/pkg/apis/compute/v1"
)

conn, err := grpc.NewClient("kubevim.example.com:443", grpc.WithTransportCredentials(credentials.NewTLS(nil)))
if err != nil {
	return err
}
defer conn.Close()

ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
resp, err := computev1.NewComputeServiceClient(conn).ListComputes(ctx, &computev1.ListComputesRequest{
	Filter: "(eq,computeName,router-1)",
})
```

### REST

```bash
curl -H "Authorization: Bearer $TOKEN" \
  "https://kubevim.example.com/compute/v1/computes?filter=(eq,computeName,router-1)"
```

## Repository layout

```
api/kubenfv/kubevim/<api>/v1/   protos, one package per API
api/kubenfv/kubevim/common/v1/  shared types and the error reasons
api/kubenfv/kubevim/options/v1/ options that map each element to IFA
pkg/apis/                       generated Go
generated/openapi/              OpenAPI 3.1 schemas
generated/python/               Python client
hack/                           OpenAPI post-processor and Python generator inputs
test/                           wire format, validation and Python client tests
kube-ovn-api/                   separate Go module with Kube-OVN types and clients
```

Generated code is committed; CI fails when it is out of date.

## Development

```bash
make generate       # Go, OpenAPI and Python client from the protos (needs Go and Docker)
make lint           # buf lint and format check
make test           # Go tests
make test-python    # Python client against a stub server (needs Python 3.10+)
```

Tools are installed into `bin/` at pinned versions. The module needs Go 1.26.

## Releases

The repository tag versions everything: the Go module, the `kubevim-client` package and the schemas' `info.version`. The major version is also in every REST path and proto package (`v1`). Tags are `vX.Y.Z[-alpha|beta|rc.N]`; from the first release candidate on, CI rejects breaking changes. The `kube-ovn-api` module is tagged separately as `kube-ovn-api/vX.Y.Z`.

## References

- [ETSI GS NFV-IFA 005](https://www.etsi.org/deliver/etsi_gs/nfv-ifa/001_099/005/05.04.01_60/gs_nfv-ifa005v050401p.pdf) v5.4.1: Or-Vi reference point, interface and information model
- [ETSI GS NFV-IFA 006](https://www.etsi.org/deliver/etsi_gs/nfv-ifa/001_099/006/05.04.01_60/gs_nfv-ifa006v050401p.pdf) v5.4.1: Vi-Vnfm reference point, interface and information model
- [ETSI GS NFV-SOL 013](https://www.etsi.org/deliver/etsi_gs/NFV-SOL/001_099/013/05.04.01_60/gs_NFV-SOL013v050401p.pdf) v5.4.1: common aspects of RESTful NFV-MANO APIs

## License

[Apache License 2.0](LICENSE)
