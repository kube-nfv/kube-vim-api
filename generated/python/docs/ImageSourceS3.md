# ImageSourceS3


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**s3** | [**S3Source**](S3Source.md) | Download from an S3-compatible object store. | 

## Example

```python
from kubevim_client.models.image_source_s3 import ImageSourceS3

# TODO update the JSON string below
json = "{}"
# create an instance of ImageSourceS3 from a JSON string
image_source_s3_instance = ImageSourceS3.from_json(json)
# print the JSON string representation of the object
print(ImageSourceS3.to_json())

# convert the object into a dict
image_source_s3_dict = image_source_s3_instance.to_dict()
# create an instance of ImageSourceS3 from a dict
image_source_s3_from_dict = ImageSourceS3.from_dict(image_source_s3_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


