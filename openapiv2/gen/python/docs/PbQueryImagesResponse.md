# PbQueryImagesResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**software_images_information** | [**List[SoftwareImageInformation]**](SoftwareImageInformation.md) |  | 

## Example

```python
from kubevim_vivnfm_client.models.pb_query_images_response import PbQueryImagesResponse

# TODO update the JSON string below
json = "{}"
# create an instance of PbQueryImagesResponse from a JSON string
pb_query_images_response_instance = PbQueryImagesResponse.from_json(json)
# print the JSON string representation of the object
print(PbQueryImagesResponse.to_json())

# convert the object into a dict
pb_query_images_response_dict = pb_query_images_response_instance.to_dict()
# create an instance of PbQueryImagesResponse from a dict
pb_query_images_response_from_dict = PbQueryImagesResponse.from_dict(pb_query_images_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


