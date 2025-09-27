# PbQueryComputeFlavourResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**flavours** | [**List[VirtualComputeFlavour]**](VirtualComputeFlavour.md) | List of Compute Flavours matching the query. | 

## Example

```python
from kubevim_vivnfm_client.models.pb_query_compute_flavour_response import PbQueryComputeFlavourResponse

# TODO update the JSON string below
json = "{}"
# create an instance of PbQueryComputeFlavourResponse from a JSON string
pb_query_compute_flavour_response_instance = PbQueryComputeFlavourResponse.from_json(json)
# print the JSON string representation of the object
print(PbQueryComputeFlavourResponse.to_json())

# convert the object into a dict
pb_query_compute_flavour_response_dict = pb_query_compute_flavour_response_instance.to_dict()
# create an instance of PbQueryComputeFlavourResponse from a dict
pb_query_compute_flavour_response_from_dict = PbQueryComputeFlavourResponse.from_dict(pb_query_compute_flavour_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


