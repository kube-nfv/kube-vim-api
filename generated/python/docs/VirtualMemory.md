# VirtualMemory

Virtual memory of a compute resource.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**hugepage_size** | **str** | Hugepage size backing the memory; empty for regular pages. | [optional] 
**numa_enabled** | **bool** |  | [optional] 
**virtual_mem_oversubscription_policy** | **str** |  | [optional] 
**virtual_mem_size** | **str** | Amount of memory, in Kubernetes quantity format. | [optional] 

## Example

```python
from kubevim_client.models.virtual_memory import VirtualMemory

# TODO update the JSON string below
json = "{}"
# create an instance of VirtualMemory from a JSON string
virtual_memory_instance = VirtualMemory.from_json(json)
# print the JSON string representation of the object
print(VirtualMemory.to_json())

# convert the object into a dict
virtual_memory_dict = virtual_memory_instance.to_dict()
# create an instance of VirtualMemory from a dict
virtual_memory_from_dict = VirtualMemory.from_dict(virtual_memory_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


