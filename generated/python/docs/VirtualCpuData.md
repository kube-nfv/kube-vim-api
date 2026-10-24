# VirtualCpuData

Virtual CPUs of a flavour.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cpu_architecture** | [**CpuArchitecture**](CpuArchitecture.md) | CPU architecture; the VIM default when unset. | [optional] 
**num_virtual_cpu** | **int** | Number of virtual CPUs. | 
**topology** | [**CpuTopology**](CpuTopology.md) | Guest CPU topology; the product of its fields equals num_virtual_cpu. The VIM rejects a topology its backend cannot express (KubeVirt takes the vCPU count and a preferred layout). | [optional] 
**virtual_cpu_pinning** | [**VirtualCpuPinningData**](VirtualCpuPinningData.md) | CPU pinning. | [optional] 

## Example

```python
from kubevim_client.models.virtual_cpu_data import VirtualCpuData

# TODO update the JSON string below
json = "{}"
# create an instance of VirtualCpuData from a JSON string
virtual_cpu_data_instance = VirtualCpuData.from_json(json)
# print the JSON string representation of the object
print(VirtualCpuData.to_json())

# convert the object into a dict
virtual_cpu_data_dict = virtual_cpu_data_instance.to_dict()
# create an instance of VirtualCpuData from a dict
virtual_cpu_data_from_dict = VirtualCpuData.from_dict(virtual_cpu_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


