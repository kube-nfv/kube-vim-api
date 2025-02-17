# PbQueryImageResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**software_image_information** | [**SoftwareImageInformation**](SoftwareImageInformation.md) |  | [optional] 

## Example

```python
from kubevim_vivnfm_client.models.pb_query_image_response import PbQueryImageResponse

# TODO update the JSON string below
json = "{}"
# create an instance of PbQueryImageResponse from a JSON string
pb_query_image_response_instance = PbQueryImageResponse.from_json(json)
# print the JSON string representation of the object
print(PbQueryImageResponse.to_json())

# convert the object into a dict
pb_query_image_response_dict = pb_query_image_response_instance.to_dict()
# create an instance of PbQueryImageResponse from a dict
pb_query_image_response_from_dict = PbQueryImageResponse.from_dict(pb_query_image_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


