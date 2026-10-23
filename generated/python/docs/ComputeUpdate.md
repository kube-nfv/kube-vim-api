# ComputeUpdate

Changes to a compute resource: a patch. The name is replaced when set, metadata entries are merged, and interfaces are added (hotplug) or removed by name.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compute_name** | **str** |  | [optional] 
**metadata** | **Dict[str, str]** | Metadata entries to add, or to overwrite when the key exists; keys not given are kept, and none is removed. | [optional] 
**network_interface_names_to_remove** | **List[str]** | Names of network interfaces to remove. | [optional] 
**network_interfaces_to_add** | [**List[NetworkInterfaceData]**](NetworkInterfaceData.md) | Network interfaces to add. | [optional] 

## Example

```python
from kubevim_client.models.compute_update import ComputeUpdate

# TODO update the JSON string below
json = "{}"
# create an instance of ComputeUpdate from a JSON string
compute_update_instance = ComputeUpdate.from_json(json)
# print the JSON string representation of the object
print(ComputeUpdate.to_json())

# convert the object into a dict
compute_update_dict = compute_update_instance.to_dict()
# create an instance of ComputeUpdate from a dict
compute_update_from_dict = ComputeUpdate.from_dict(compute_update_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


