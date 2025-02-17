# VirtualCpuPinningDataVirtualCpuPinningRule


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cores** | **float** | The number of core in the virtual CPU. | [optional] 
**sockets** | **float** | The number of socket in the virtual CPU. | [optional] 
**threads** | **float** | The number of thread in the virtual CPU. | [optional] 

## Example

```python
from vivnfm_client.models.virtual_cpu_pinning_data_virtual_cpu_pinning_rule import VirtualCpuPinningDataVirtualCpuPinningRule

# TODO update the JSON string below
json = "{}"
# create an instance of VirtualCpuPinningDataVirtualCpuPinningRule from a JSON string
virtual_cpu_pinning_data_virtual_cpu_pinning_rule_instance = VirtualCpuPinningDataVirtualCpuPinningRule.from_json(json)
# print the JSON string representation of the object
print(VirtualCpuPinningDataVirtualCpuPinningRule.to_json())

# convert the object into a dict
virtual_cpu_pinning_data_virtual_cpu_pinning_rule_dict = virtual_cpu_pinning_data_virtual_cpu_pinning_rule_instance.to_dict()
# create an instance of VirtualCpuPinningDataVirtualCpuPinningRule from a dict
virtual_cpu_pinning_data_virtual_cpu_pinning_rule_from_dict = VirtualCpuPinningDataVirtualCpuPinningRule.from_dict(virtual_cpu_pinning_data_virtual_cpu_pinning_rule_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


