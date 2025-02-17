# PbCreateComputeFlavourResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**flavour_id** | [**Identifier**](Identifier.md) |  | [optional] 

## Example

```python
from kubevim_vivnfm_client.models.pb_create_compute_flavour_response import PbCreateComputeFlavourResponse

# TODO update the JSON string below
json = "{}"
# create an instance of PbCreateComputeFlavourResponse from a JSON string
pb_create_compute_flavour_response_instance = PbCreateComputeFlavourResponse.from_json(json)
# print the JSON string representation of the object
print(PbCreateComputeFlavourResponse.to_json())

# convert the object into a dict
pb_create_compute_flavour_response_dict = pb_create_compute_flavour_response_instance.to_dict()
# create an instance of PbCreateComputeFlavourResponse from a dict
pb_create_compute_flavour_response_from_dict = PbCreateComputeFlavourResponse.from_dict(pb_create_compute_flavour_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


