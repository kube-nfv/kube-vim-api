# CreateFlavourRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**flavour** | [**Flavour**](Flavour.md) | Flavour to create. | 
**request_id** | **str** | Client-chosen UUID that makes the request idempotent (AIP-155). | [optional] 

## Example

```python
from kubevim_client.models.create_flavour_request import CreateFlavourRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CreateFlavourRequest from a JSON string
create_flavour_request_instance = CreateFlavourRequest.from_json(json)
# print the JSON string representation of the object
print(CreateFlavourRequest.to_json())

# convert the object into a dict
create_flavour_request_dict = create_flavour_request_instance.to_dict()
# create an instance of CreateFlavourRequest from a dict
create_flavour_request_from_dict = CreateFlavourRequest.from_dict(create_flavour_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


