# KubeVim API

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/kube-nfv/kube-vim-api.svg)](https://pkg.go.dev/github.com/kube-nfv/kube-vim-api)

Protocol Buffer definitions and generated clients for KubeVim ETSI MANO (Management and Orchestration) interfaces.

## Overview

This repository contains Protocol Buffer definitions of the **VI-VNFM** and **OR-VI** ETSI MANO reference points described in ETSI GS NFV-IFA 005 (OR-VI) [^fn1] and ETSI GS NFV-IFA 006 (VI-VNFM) [^fn2]. Some model definitions are also used from ETSI GS NFV-SOL 014.

The RESTful reverse proxy is generated using grpc-gateway and conforms to ETSI GS NFV-SOL 013 [^fn3] specifications.

### ETSI Standards Compliance

This implementation provides:
- **VI-VNFM Interface**: Standardized interface between Virtualized Infrastructure Manager (VIM) and VNF Manager (VNFM)
- **OR-VI Interface**: Orchestrator to VIM interface for resource management
- **RESTful API Binding**: HTTP/REST mapping following ETSI SOL 013 guidelines
- **NFV MANO Compliance**: Full compliance with ETSI NFV Management and Orchestration standards

## Repository Structure

```
kube-vim-api/
├── api/                          # Protocol Buffer definitions
│   ├── common/                   # Shared types and ETSI common models
│   │   ├── common.proto          # Basic types (Identifier, Filter, etc.)
│   │   └── types.proto           # ETSI NFV types (SoftwareImageInformation, etc.)
│   ├── vivnfm/                   # ETSI VI-VNFM interface (IFA-006)
│   │   └── vi-vnfm.proto         # Complete VI-VNFM API definition
│   └── admin/                    # KubeVim administrative extensions
│       └── kubevim-admin.proto   # Non-standard admin operations
├── pkg/apis/                     # Generated Go code
│   ├── common.pb.go              # Generated common types
│   ├── types.pb.go               # Generated ETSI types
│   ├── vivnfm/                   # VI-VNFM Go client
│   │   ├── vi-vnfm.pb.go         # Protocol Buffer types
│   │   ├── vi-vnfm_grpc.pb.go    # gRPC client/server
│   │   └── vi-vnfm.pb.gw.go      # gRPC-Gateway reverse proxy
│   └── admin/                    # Admin API Go client
│       ├── kubevim-admin.pb.go
│       ├── kubevim-admin_grpc.pb.go
│       └── kubevim-admin.pb.gw.go
├── generated/                    # Generated clients (non-Go)
│   ├── openapi/                  # OpenAPI/Swagger specifications
│   │   └── vivnfm/
│   │       └── vi-vnfm.swagger.json
│   └── python/                   # Python client library
│       ├── kubevim_vivnfm_client/
│       ├── docs/                 # API documentation
│       └── setup.py              # Python package setup
└── kube-ovn-api/                # Kube-OVN integration types
    └── pkg/apis/kubeovn/         # Golang types for kube-ovn networking
```

## Quick Start

### Prerequisites
- Go 1.23+
- Docker
- Make
- Python 3.8+ (for Python client)

### Build All Clients
```bash
# Generate all code (Go, OpenAPI specs, Python client)
make all
```

### Generate Specific Clients
```bash
make proto-compile    # Go code and OpenAPI specs
make python-gen       # Python client only
make proto-clean      # Clean all generated code
```

## Usage Examples

### Go Client (VI-VNFM Interface)
```go
package main

import (
    "context"
    "log"

    "github.com/kube-nfv/kube-vim-api/pkg/apis"
    "github.com/kube-nfv/kube-vim-api/pkg/apis/vivnfm"
    "google.golang.org/grpc"
)

func main() {
    // Connect to KubeVim gRPC server
    conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    // Create VI-VNFM client
    client := vivnfm.NewViVnfmClient(conn)

    // Query available software images
    req := &vivnfm.QueryImagesRequest{
        ImageQueryFilter: &apis.Filter{Value: "ubuntu"},
    }

    resp, err := client.QueryImages(context.Background(), req)
    if err != nil {
        log.Fatal(err)
    }

    for _, image := range resp.SoftwareImagesInformation {
        log.Printf("Image: %s, Provider: %s", image.Name, image.Provider)
    }
}
```

### Python Client
```python
import kubevim_vivnfm_client
from kubevim_vivnfm_client.api import vi_vnfm_api
from kubevim_vivnfm_client.model.pb_query_images_request import PbQueryImagesRequest

# Configure client
configuration = kubevim_vivnfm_client.Configuration(
    host="http://localhost:8080"  # KubeVim REST API endpoint
)

# Create API instance
with kubevim_vivnfm_client.ApiClient(configuration) as api_client:
    api_instance = vi_vnfm_api.ViVnfmApi(api_client)

    try:
        # Query software images
        response = api_instance.query_images()

        for image in response.software_images_information:
            print(f"Image: {image.name}, Provider: {image.provider}")

    except kubevim_vivnfm_client.ApiException as e:
        print(f"Exception when calling ViVnfmApi->query_images: {e}")
```

### REST API Examples
```bash
# Query software images
curl -X GET "http://localhost:8080/vivnfm/v5/images" \
     -H "accept: application/json"

# Allocate compute resources
curl -X POST "http://localhost:8080/vivnfm/v5/compute" \
     -H "accept: application/json" \
     -H "Content-Type: application/json" \
     -d '{
       "computeFlavourId": {"value": "m1.small"},
       "vcImageId": {"value": "ubuntu-20.04"}
     }'

# Admin: Download software image (KubeVim extension)
curl -X POST "http://localhost:8080/admin/v1/images" \
     -H "accept: application/json" \
     -H "Content-Type: application/json" \
     -d '{}'
```

## API Documentation

### Available APIs

| API | Description | Standard | Package |
|-----|-------------|----------|---------|
| **VI-VNFM** | VIM-VNFM interface for resource management | ETSI IFA-006 | `pkg/apis/vivnfm` |
| **Admin** | KubeVim administrative extensions | Custom | `pkg/apis/admin` |
| **Common** | Shared ETSI types and utilities | ETSI IFA-005/006 | `pkg/apis` |

### Key Operations (VI-VNFM)

#### Software Image Management
- `QueryImages` - Query available software images
- `QueryImage` - Get specific software image details

#### Virtualized Compute Resources
- `AllocateVirtualisedComputeResource` - Request compute allocation
- `QueryVirtualisedComputeResource` - Query compute resources
- `TerminateVirtualisedComputeResource` - Terminate compute resources
- `OperateVirtualisedComputeResource` - Perform operations on compute

#### Compute Flavour Management
- `CreateComputeFlavour` - Create compute flavours
- `QueryComputeFlavour` - Query available flavours
- `DeleteComputeFlavour` - Remove compute flavours

#### Virtualized Network Resources
- `AllocateVirtualisedNetworkResource` - Allocate network resources
- `QueryVirtualisedNetworkResource` - Query network resources
- `TerminateVirtualisedNetworkResource` - Terminate network resources

For complete API reference, see the [OpenAPI documentation](generated/openapi/).

## Development

### Adding New APIs
1. Create `.proto` files in appropriate `api/` subdirectory
2. Update imports in existing protos if needed
3. Run `make proto-compile` to generate code
4. Update documentation and examples

### Testing
```bash
# Run Go tests
go test ./pkg/apis/...

# Test Python client (if generated)
cd generated/python
python -m pytest test/
```

### Code Generation
The build system uses:
- **protoc** with grpc-gateway for Go code generation
- **OpenAPI Generator** for Python client generation
- **Docker** for reproducible builds across platforms

## Installation

### Go Module
```bash
go get github.com/kube-nfv/kube-vim-api
```

### Python Package
```bash
# Install from generated code
pip install ./generated/python

# Or install in development mode
cd generated/python && pip install -e .
```

## Integration

### Kube-OVN Support
The `/kube-ovn-api` directory contains Golang types to work with [kube-ovn](https://www.kube-ovn.io) networking, including code autogeneration logic for Golang ClientSet and helpers.

## References

[^fn1]: [ETSI GS NFV-IFA 005 Release v5.2.1 (2024-11)](https://www.etsi.org/deliver/etsi_gs/nfv-ifa/001_099/005/05.02.01_60/gs_nfv-ifa005v050201p.pdf) - Network Functions Virtualisation (NFV) Release 4; Management and Orchestration; Or-Vi reference point - Interface and Information Model Specification

[^fn2]: [ETSI GS NFV-IFA 006 Release v5.2.1 (2024-12)](https://www.etsi.org/deliver/etsi_gs/nfv-ifa/001_099/006/05.02.01_60/gs_nfv-ifa006v050201p.pdf) - Network Functions Virtualisation (NFV) Release 4; Management and Orchestration; Vi-Vnfm reference point - Interface and Information Model Specification

[^fn3]: [ETSI GS NFV-SOL 013 Release v5.2.1 (2024-12)](https://www.etsi.org/deliver/etsi_gs/NFV-SOL/001_099/013/05.02.01_60/gs_nfv-sol013v050201p.pdf) - Network Functions Virtualisation (NFV) Release 4; Protocols and Data Models; RESTful protocols specification for the Or-Vi reference point
