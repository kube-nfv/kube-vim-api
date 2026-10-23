# CreatePlacementGroupRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**placement_group** | [**PlacementGroup**](PlacementGroup.md) | Group to create. | 
**request_id** | **str** | Client-chosen UUID that makes the request idempotent (AIP-155). | [optional] 

## Example

```python
from kubevim_client.models.create_placement_group_request import CreatePlacementGroupRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CreatePlacementGroupRequest from a JSON string
create_placement_group_request_instance = CreatePlacementGroupRequest.from_json(json)
# print the JSON string representation of the object
print(CreatePlacementGroupRequest.to_json())

# convert the object into a dict
create_placement_group_request_dict = create_placement_group_request_instance.to_dict()
# create an instance of CreatePlacementGroupRequest from a dict
create_placement_group_request_from_dict = CreatePlacementGroupRequest.from_dict(create_placement_group_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


