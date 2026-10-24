# PlacementGroup

A group of compute resources placed together or apart.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compute_ids** | **List[str]** | Compute resources in the group. | [optional] [readonly] 
**enforcement** | [**PlacementEnforcement**](PlacementEnforcement.md) | Whether the policy must hold or is applied when possible; PLACEMENT_ENFORCEMENT_REQUIRED when unset. | [optional] 
**metadata** | **Dict[str, str]** | Metadata set by the consumer and by the VIM; for keys under \&quot;kubevim.kubenfv.io/\&quot; that the VIM maintains, its value wins. | [optional] 
**placement_group_id** | **str** | Identifier of the group: a UUID assigned by the VIM. | [optional] [readonly] 
**placement_group_name** | **str** | Name of the group, given by the consumer; not unique. | [optional] 
**scope** | [**PlacementScope**](PlacementScope.md) | Topology domain the policy applies to; PLACEMENT_SCOPE_NODE when unset. | [optional] 
**type** | [**PlacementType**](PlacementType.md) | Whether members are placed together or apart. | 

## Example

```python
from kubevim_client.models.placement_group import PlacementGroup

# TODO update the JSON string below
json = "{}"
# create an instance of PlacementGroup from a JSON string
placement_group_instance = PlacementGroup.from_json(json)
# print the JSON string representation of the object
print(PlacementGroup.to_json())

# convert the object into a dict
placement_group_dict = placement_group_instance.to_dict()
# create an instance of PlacementGroup from a dict
placement_group_from_dict = PlacementGroup.from_dict(placement_group_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


