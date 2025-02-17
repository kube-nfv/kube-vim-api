# VirtualCpuDataVirtualCpuPinningData

Information describing CPU pinning policy and rules for virtual CPU to physical CPU mapping of the virtualised compute resource.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**virtual_cpu_pinning_policy** | [**VirtualCpuPinningDataVirtualCpuPinningPolicy**](VirtualCpuPinningDataVirtualCpuPinningPolicy.md) |  | [optional] [default to VirtualCpuPinningDataVirtualCpuPinningPolicy.STATIC]
**virtual_cpu_pinning_rules** | [**List[VirtualCpuPinningDataVirtualCpuPinningRule]**](VirtualCpuPinningDataVirtualCpuPinningRule.md) | A list of rules that should be considered during the allocation of the virtual CPU-s to logical CPU-s in case of \&quot;static\&quot; virtualCpuPinningPolicy. | [optional] 

## Example

```python
from vivnfm_client.models.virtual_cpu_data_virtual_cpu_pinning_data import VirtualCpuDataVirtualCpuPinningData

# TODO update the JSON string below
json = "{}"
# create an instance of VirtualCpuDataVirtualCpuPinningData from a JSON string
virtual_cpu_data_virtual_cpu_pinning_data_instance = VirtualCpuDataVirtualCpuPinningData.from_json(json)
# print the JSON string representation of the object
print(VirtualCpuDataVirtualCpuPinningData.to_json())

# convert the object into a dict
virtual_cpu_data_virtual_cpu_pinning_data_dict = virtual_cpu_data_virtual_cpu_pinning_data_instance.to_dict()
# create an instance of VirtualCpuDataVirtualCpuPinningData from a dict
virtual_cpu_data_virtual_cpu_pinning_data_from_dict = VirtualCpuDataVirtualCpuPinningData.from_dict(virtual_cpu_data_virtual_cpu_pinning_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


