# PbQueryImageRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**software_image_id** | [**Identifier**](Identifier.md) |  | [optional] 

## Example

```python
from kubevim_vivnfm_client.models.pb_query_image_request import PbQueryImageRequest

# TODO update the JSON string below
json = "{}"
# create an instance of PbQueryImageRequest from a JSON string
pb_query_image_request_instance = PbQueryImageRequest.from_json(json)
# print the JSON string representation of the object
print(PbQueryImageRequest.to_json())

# convert the object into a dict
pb_query_image_request_dict = pb_query_image_request_instance.to_dict()
# create an instance of PbQueryImageRequest from a dict
pb_query_image_request_from_dict = PbQueryImageRequest.from_dict(pb_query_image_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


