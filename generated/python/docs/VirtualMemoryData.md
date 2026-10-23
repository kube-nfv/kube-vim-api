# VirtualMemoryData

Virtual memory of a flavour.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**hugepage_size** | **str** | Hugepage size backing the memory, \&quot;2Mi\&quot; or \&quot;1Gi\&quot;; regular pages when unset. | [optional] 
**numa_enabled** | **bool** | Memory is allocated on the NUMA nodes of the pinned CPUs; needs dynamic pinning and hugepages. | [optional] 
**virtual_mem_oversubscription_policy** | **str** | Memory oversubscription as \&quot;&lt;virtual&gt;:&lt;physical&gt;\&quot;, at least 1:1, e.g. \&quot;2:1\&quot; (KubeVirt overcommitPercent). | [optional] 
**virtual_mem_size** | **str** | Amount of memory, in Kubernetes quantity format (e.g. \&quot;4Gi\&quot;). | 

## Example

```python
from kubevim_client.models.virtual_memory_data import VirtualMemoryData

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


