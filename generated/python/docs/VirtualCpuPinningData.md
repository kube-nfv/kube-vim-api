# VirtualCpuPinningData

CPU pinning of a flavour.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**virtual_cpu_pinning_policy** | [**VirtualCpuPinningPolicy**](VirtualCpuPinningPolicy.md) | Pinning policy. | 

## Example

```python
from kubevim_client.models.virtual_cpu_pinning_data import VirtualCpuPinningData

# TODO update the JSON string below
json = "{}"
# create an instance of VirtualCpuPinningData from a JSON string
virtual_cpu_pinning_data_instance = VirtualCpuPinningData.from_json(json)
# print the JSON string representation of the object
print(VirtualCpuPinningData.to_json())

# convert the object into a dict
virtual_cpu_pinning_data_dict = virtual_cpu_pinning_data_instance.to_dict()
# create an instance of VirtualCpuPinningData from a dict
virtual_cpu_pinning_data_from_dict = VirtualCpuPinningData.from_dict(virtual_cpu_pinning_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


