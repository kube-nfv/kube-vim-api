# ImageSource

Source of a software image; exactly one member is set.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**http** | [**HttpSource**](HttpSource.md) | Download over HTTP or HTTPS. | 
**s3** | [**S3Source**](S3Source.md) | Download from an S3-compatible object store. | 
**registry** | [**RegistrySource**](RegistrySource.md) | Pull from a container registry. | 
**upload** | [**UploadSource**](UploadSource.md) | Upload by the client after creation. | 
**glance** | [**GlanceSource**](GlanceSource.md) | Import an image of the Glance catalog configured in the VIM. | 
**local** | [**LocalSource**](LocalSource.md) | Import an image file of the local catalog configured in the VIM. | 

## Example

```python
from kubevim_client.models.image_source import ImageSource

# TODO update the JSON string below
json = "{}"
# create an instance of ImageSource from a JSON string
image_source_instance = ImageSource.from_json(json)
# print the JSON string representation of the object
print(ImageSource.to_json())

# convert the object into a dict
image_source_dict = image_source_instance.to_dict()
# create an instance of ImageSource from a dict
image_source_from_dict = ImageSource.from_dict(image_source_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


