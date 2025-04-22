# TypeVirtualNic

Type of network interface. The type allows for defining how such interface is to be realized, e.g. normal virtual NIC, with direct PCI pass-through, SR-IOV, etc. Note(dmalovan): typeVirtualNic should define how the VM nic should be connected to the host network.   - BRIDGE: In the simple mode VM's are connected to the network backend through a linux \"bridge\".  - PATHTHROUGH: In masquerade mode, kubevirt allocates internal IP address to the virtual machines and hide them behind NAT. All egress treaffic is \"SNAT'ed\". Note(dmalovan): actually don't think that is needed for NFV usecases. MASQUERADE = 1; PCI Path-through allows a compute instance to have a direct and exclusive access to the physical PCI devices.  - SRIOV: Extension of PCIe that allows a single PCI device (NIC) to be shared among multiple VMs.

## Enum

* `BRIDGE` (value: `'BRIDGE'`)

* `PATHTHROUGH` (value: `'PATHTHROUGH'`)

* `SRIOV` (value: `'SRIOV'`)

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


