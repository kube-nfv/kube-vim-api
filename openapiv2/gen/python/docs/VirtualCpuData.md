# VirtualCpuData

Information describing a virtual CPU.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cpu_architecture** | **str** | CPU architecture type. Examples are x86, ARM. The cardinality can be 0 during the allocation request, if no particular CPU architecture type is requested. | [optional] 
**num_virtual_cpu** | **int** | Number of virtual CPUs. | [optional] 
**cpu_clock** | **float** | Minimum CPU clock rate (e.g. in MHz) available for the virtualised CPU resources. The cardinality can be 0 during the allocation request, if no particular value is requested. | [optional] 
**virtual_cpu_oversubscription_policy** | **str** |  | [optional] 
**virtual_cpu_pinning** | [**VirtualCpuDataVirtualCpuPinningData**](VirtualCpuDataVirtualCpuPinningData.md) |  | [optional] 
**power_state_reqs** | **str** | Virtual CPU power (state) requirements for the virtualised compute resource. | [optional] 

## Example

```python
from kubevim_vivnfm_client.models.virtual_cpu_data import VirtualCpuData

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


