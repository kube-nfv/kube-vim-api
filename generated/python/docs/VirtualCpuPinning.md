# VirtualCpuPinning

CPU pinning of a compute resource.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cpu_map** | **Dict[str, int]** | Physical CPU each virtual CPU is pinned to; empty when the VIM cannot tell. | [optional] 
**cpu_pinning_policy** | [**VirtualCpuPinningPolicy**](VirtualCpuPinningPolicy.md) |  | [optional] 

## Example

```python
from kubevim_client.models.virtual_cpu_pinning import VirtualCpuPinning

# TODO update the JSON string below
json = "{}"
# create an instance of VirtualCpuPinning from a JSON string
virtual_cpu_pinning_instance = VirtualCpuPinning.from_json(json)
# print the JSON string representation of the object
print(VirtualCpuPinning.to_json())

# convert the object into a dict
virtual_cpu_pinning_dict = virtual_cpu_pinning_instance.to_dict()
# create an instance of VirtualCpuPinning from a dict
virtual_cpu_pinning_from_dict = VirtualCpuPinning.from_dict(virtual_cpu_pinning_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


