# PbCreateComputeFlavourRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**flavour** | [**VirtualComputeFlavour**](VirtualComputeFlavour.md) |  | 
**meta_data** | [**Metadata**](Metadata.md) |  | [optional] 

## Example

```python
from kubevim_vivnfm_client.models.pb_create_compute_flavour_request import PbCreateComputeFlavourRequest

# TODO update the JSON string below
json = "{}"
# create an instance of PbCreateComputeFlavourRequest from a JSON string
pb_create_compute_flavour_request_instance = PbCreateComputeFlavourRequest.from_json(json)
# print the JSON string representation of the object
print(PbCreateComputeFlavourRequest.to_json())

# convert the object into a dict
pb_create_compute_flavour_request_dict = pb_create_compute_flavour_request_instance.to_dict()
# create an instance of PbCreateComputeFlavourRequest from a dict
pb_create_compute_flavour_request_from_dict = PbCreateComputeFlavourRequest.from_dict(pb_create_compute_flavour_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


