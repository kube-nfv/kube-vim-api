# VirtualCpu

Virtual CPUs of a compute resource.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cpu_architecture** | [**CpuArchitecture**](CpuArchitecture.md) |  | [optional] 
**num_virtual_cpu** | **int** |  | [optional] 
**virtual_cpu_pinning** | [**VirtualCpuPinning**](VirtualCpuPinning.md) |  | [optional] 

## Example

```python
from kubevim_client.models.virtual_cpu import VirtualCpu

# TODO update the JSON string below
json = "{}"
# create an instance of VirtualCpu from a JSON string
virtual_cpu_instance = VirtualCpu.from_json(json)
# print the JSON string representation of the object
print(VirtualCpu.to_json())

# convert the object into a dict
virtual_cpu_dict = virtual_cpu_instance.to_dict()
# create an instance of VirtualCpu from a dict
virtual_cpu_from_dict = VirtualCpu.from_dict(virtual_cpu_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


