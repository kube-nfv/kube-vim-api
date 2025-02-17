# VirtualComputeFlavour

The VirtualComputeFlavour information element encapsulates information for compute flavours. A compute flavour includes information about number of virtual CPUs, size of virtual memory, size of virtual storage, and virtual network interfaces.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**flavour_id** | [**Identifier**](Identifier.md) |  | [optional] 
**is_public** | **bool** | Scope of flavour accessibility. It indicates if the compute flavour is accessible and shared across clients. Default value is True (if not specified), which means public. False means private. | [optional] 
**virtual_memory** | [**VirtualMemoryData**](VirtualMemoryData.md) |  | [optional] 
**virtual_cpu** | [**VirtualCpuData**](VirtualCpuData.md) |  | [optional] 
**storage_attributes** | [**List[VirtualStorageData]**](VirtualStorageData.md) | Contains information about the size of virtualised storage resource (e.g. size of volume, in GB), the type of storage (e.g. volume, object), and support for RDMA. | [optional] 
**virtual_network_interface** | [**List[VirtualNetworkInterfaceData]**](VirtualNetworkInterfaceData.md) |  | [optional] 
**metadata** | [**Metadata**](Metadata.md) |  | [optional] 

## Example

```python
from kubevim_vivnfm_client.models.virtual_compute_flavour import VirtualComputeFlavour

# TODO update the JSON string below
json = "{}"
# create an instance of VirtualComputeFlavour from a JSON string
virtual_compute_flavour_instance = VirtualComputeFlavour.from_json(json)
# print the JSON string representation of the object
print(VirtualComputeFlavour.to_json())

# convert the object into a dict
virtual_compute_flavour_dict = virtual_compute_flavour_instance.to_dict()
# create an instance of VirtualComputeFlavour from a dict
virtual_compute_flavour_from_dict = VirtualComputeFlavour.from_dict(virtual_compute_flavour_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


