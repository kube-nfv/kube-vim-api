# ImageSourceLocal


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**local** | [**LocalSource**](LocalSource.md) | Import an image file of the local catalog configured in the VIM. | 

## Example

```python
from kubevim_client.models.image_source_local import ImageSourceLocal

# TODO update the JSON string below
json = "{}"
# create an instance of ImageSourceLocal from a JSON string
image_source_local_instance = ImageSourceLocal.from_json(json)
# print the JSON string representation of the object
print(ImageSourceLocal.to_json())

# convert the object into a dict
image_source_local_dict = image_source_local_instance.to_dict()
# create an instance of ImageSourceLocal from a dict
image_source_local_from_dict = ImageSourceLocal.from_dict(image_source_local_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


