# ImageSourceUpload


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**upload** | [**UploadSource**](UploadSource.md) | Upload by the client after creation. | 

## Example

```python
from kubevim_client.models.image_source_upload import ImageSourceUpload

# TODO update the JSON string below
json = "{}"
# create an instance of ImageSourceUpload from a JSON string
image_source_upload_instance = ImageSourceUpload.from_json(json)
# print the JSON string representation of the object
print(ImageSourceUpload.to_json())

# convert the object into a dict
image_source_upload_dict = image_source_upload_instance.to_dict()
# create an instance of ImageSourceUpload from a dict
image_source_upload_from_dict = ImageSourceUpload.from_dict(image_source_upload_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


