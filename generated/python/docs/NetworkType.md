# NetworkType

- NETWORK_TYPE_SRIOV: SR-IOV VF pool realized via sriov-cni + Multus NetworkAttachmentDefinition. providerNetwork = device plugin resource name (k8s.v1.cni.cncf.io/resourceName) segmentationId  = VLAN ID (0 = untagged) bandwidth       = min_tx_rate in Mbps (0 = no guarantee)

## Enum

* `NETWORK_TYPE_OVERLAY` (value: `'NETWORK_TYPE_OVERLAY'`)

* `NETWORK_TYPE_UNDERLAY` (value: `'NETWORK_TYPE_UNDERLAY'`)

* `NETWORK_TYPE_SRIOV` (value: `'NETWORK_TYPE_SRIOV'`)

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


