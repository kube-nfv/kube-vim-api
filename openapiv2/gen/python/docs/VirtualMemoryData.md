# VirtualMemoryData

Information describing virtual memory.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**virtual_mem_size** | **float** | Amount of virtual Memory (e.g. in MB). | 
**virtual_mem_oversubscription_policy** | **str** | Memory core oversubscription policy in terms of virtual memory to physical memory on the platform. The cardinality can be 0 during the allocation request, if no particular value is requested. | [optional] 
**numa_enabled** | **bool** |  | [optional] 

## Example

```python
from kubevim_vivnfm_client.models.virtual_memory_data import VirtualMemoryData

# TODO update the JSON string below
json = "{}"
# create an instance of VirtualMemoryData from a JSON string
virtual_memory_data_instance = VirtualMemoryData.from_json(json)
# print the JSON string representation of the object
print(VirtualMemoryData.to_json())

# convert the object into a dict
virtual_memory_data_dict = virtual_memory_data_instance.to_dict()
# create an instance of VirtualMemoryData from a dict
virtual_memory_data_from_dict = VirtualMemoryData.from_dict(virtual_memory_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


