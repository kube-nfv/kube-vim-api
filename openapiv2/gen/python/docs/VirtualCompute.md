# VirtualCompute


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compute_id** | [**Identifier**](Identifier.md) |  | [optional] 
**compute_name** | **str** | Name of the virtualised compute resource. | [optional] 
**flavour_id** | [**Identifier**](Identifier.md) |  | [optional] 
**virtual_cpu** | **object** |  | [optional] 
**virtual_memory** | **object** |  | [optional] 
**virtual_network_interface** | **List[object]** | Provides information of the instantiated virtual network interfaces of the compute resource. | [optional] 
**virtual_disks** | **List[object]** | Provides information of the virtualised storage resources (volumes, ephemeral) that are attached to the compute resource. | [optional] 
**vc_image_id** | [**Identifier**](Identifier.md) |  | [optional] 
**zone_id** | [**Identifier**](Identifier.md) |  | [optional] 
**host_id** | [**Identifier**](Identifier.md) |  | [optional] 
**operational_state** | [**OperationalState**](OperationalState.md) |  | [optional] [default to OperationalState.ENABLED]
**running_state** | [**ComputeRunningState**](ComputeRunningState.md) |  | [optional] [default to ComputeRunningState.STARTED]
**metadata** | [**Metadata**](Metadata.md) |  | [optional] 

## Example

```python
from kubevim_vivnfm_client.models.virtual_compute import VirtualCompute

# TODO update the JSON string below
json = "{}"
# create an instance of VirtualCompute from a JSON string
virtual_compute_instance = VirtualCompute.from_json(json)
# print the JSON string representation of the object
print(VirtualCompute.to_json())

# convert the object into a dict
virtual_compute_dict = virtual_compute_instance.to_dict()
# create an instance of VirtualCompute from a dict
virtual_compute_from_dict = VirtualCompute.from_dict(virtual_compute_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


