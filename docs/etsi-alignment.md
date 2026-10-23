# Kube-VIM API and ETSI NFV

The Kube-VIM API is not a home-grown design. It is built on the ETSI NFV-MANO specifications, the industry standard for managing virtualized network functions, which orchestrators such as [OSM](https://osm.etsi.org) are built around. This page explains what that means, what Kube-VIM adds for Kubernetes, and what it leaves out.

## ETSI NFV in short

ETSI NFV-MANO (Management and Orchestration) splits the job of running virtual network functions into three roles:

- the **NFVO** (orchestrator) manages network services,
- the **VNFM** (VNF manager) manages the lifecycle of each network function,
- the **VIM** (Virtualized Infrastructure Manager) manages compute, network and storage resources. Kube-VIM is a VIM.

The NFVO talks to the VIM over the **Or-Vi** reference point and the VNFM over **Vi-Vnfm**. ETSI describes these in two families of specifications:

- **IFA** specifications define the information model: which resources exist, their attributes and their operations. [IFA 005](https://www.etsi.org/deliver/etsi_gs/nfv-ifa/001_099/005/05.04.01_60/gs_nfv-ifa005v050401p.pdf) covers Or-Vi and [IFA 006](https://www.etsi.org/deliver/etsi_gs/nfv-ifa/001_099/006/05.04.01_60/gs_nfv-ifa006v050401p.pdf) Vi-Vnfm.
- **SOL** specifications define protocols. [SOL 013](https://www.etsi.org/deliver/etsi_gs/NFV-SOL/001_099/013/05.04.01_60/gs_NFV-SOL013v050401p.pdf) sets the common rules for ETSI REST APIs: URIs, errors, filtering, paging and authorization.

ETSI publishes no REST protocol for Or-Vi or Vi-Vnfm themselves, so every VIM API is a mapping of the IFA model onto some protocol. Kube-VIM maps it onto gRPC and REST, following SOL 013. The API is built on and aligned with these specifications; it is not certified against them, as no ETSI programme certifies Or-Vi or Vi-Vnfm implementations.

## What comes from ETSI

- **The interfaces.** One API per IFA management interface: software images, compute flavours, compute, networks and storage.
- **The resources and attributes.** Images, flavours, virtual machines (IFA's virtualised compute), networks, subnets and volumes carry IFA's attributes under IFA's names, e.g. `virtualMemSize`, `numVirtualCpu`, `segmentationId`, `operationalState`.
- **The operations.** Allocate, query, update, terminate, operate (start, stop, pause, reboot), scale, migrate, attach and detach storage, and snapshots. Each RPC documents the IFA operation it implements.
- **Or-Vi and Vi-Vnfm.** IFA gives both reference points the same interfaces; Kube-VIM serves both with the same APIs and tells them apart by OAuth 2.0 scopes.
- **The REST conventions** of SOL 013: `/{api}/v1/` paths, ProblemDetails errors, the attribute filter syntax, paging with `nextpage_opaque_marker` and the `Link` header, 201 with `Location` on creation, and OAuth 2.0 scopes.

## How the API presents the model

IFA describes information, not an API. Kube-VIM keeps its meaning and presents it the way modern APIs are designed ([Google AIP](https://google.aip.dev) and the Protocol Buffers style guide):

- IFA's Allocate, Query and Terminate are Create, Get, List and Delete; Query is split into Get for one resource and List with a filter.
- Identifiers are named after their resource (`computeId`, `networkId`), and every reference uses the same name.
- Values that IFA leaves as free text are typed enums, such as image status or compute operation.
- Sizes are Kubernetes quantities such as `"10Gi"` instead of numbers with an implied unit.
- Update changes only the fields it is given.
- Snapshots and placement groups are resources that can be listed and read back; in IFA they can only be created.
- SOL 013's version negotiation (`Version` header, `api_versions`) is not used: the major version is in the path and the release version in the schema.

## Extensions

What Kube-VIM adds to the IFA model.

**All APIs**
- A `metadata` map and a name on every resource, flavours and subnets included.

**Images**
- Import from HTTP, S3, a container registry, a client upload, or the Glance or local catalog configured in Kube-VIM, instead of sending the image in the request; Glance and local catalog images are listed without importing them first.
- Import policy: import when the image is created, or when the first VM uses it.
- Import progress and failure reason.
- Firmware (BIOS, EFI, EFI with secure boot) and boot disk bus, for images that need them.

**Flavours**
- Hugepages and guest CPU topology (sockets, cores, threads).

**Networks**
- Provider networks by name; Kube-VIM chooses the backend (overlay, VLAN or SR-IOV).
- Dual-stack: a network holds one IPv4 and one IPv6 subnet, and an interface gets an address from each.
- MTU and DNS servers.

**Compute**
- Placement groups: affinity or anti-affinity per node or zone, required or preferred.
- Create a VM stopped, resume after pause.
- Interface name, NIC model, address assignment mode, and the guest and host PCI addresses and guest interface name, so consumers can find each interface inside the VM.
- A management interface on every VM, a failure reason, and a flag when a change needs a restart.
- Deleting a VM can also delete its attached volumes.

**Storage**
- Kubernetes storage classes for volumes and flavour disks, and migration to another storage class.

## Left out

IFA attributes and operations Kube-VIM has no backend for. Each can be added later without breaking clients.

**All APIs**
- Reservations, quotas, location constraints and resource groups.
- Tenancy attributes: image visibility, private flavours, non-shared networks. Every resource is visible to every client.
- Per-request affinity constraints and network or storage affinity groups; placement groups replace them.

**Images**
- Checksum, which the import backend does not report.

**Flavours**
- CPU clock, CPU oversubscription, CPU power states and static pinning rules; pinning is dynamic.
- RDMA storage, acceleration capabilities, and network interfaces in the flavour (they are given when the VM is created).

**Networks**
- Ports and trunks: interfaces are created with the VM.
- Network bandwidth and QoS, connected networks, more than one subnet per IP version.

**Compute**
- Updating an interface in place (remove and add it instead) and power state through Update (use operate).
- Offline migration and choosing the target host: migration is live and the scheduler picks the host.
- Suspend to disk, certificate data, disk mount points, parameters for operate.

**Storage**
- Storage flavours and object storage.

**Interfaces not in v1**
- Change notifications, resource information and capacity, fault and performance management, reservations and quotas. They are separate ETSI interfaces and would come as separate APIs.
