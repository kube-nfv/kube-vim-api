# ImageSourceGlance


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**glance** | [**GlanceSource**](GlanceSource.md) | Import an image of the Glance catalog configured in the VIM. | 

## Example

```python
from kubevim_client.models.image_source_glance import ImageSourceGlance

# TODO update the JSON string below
json = "{}"
# create an instance of ImageSourceGlance from a JSON string
image_source_glance_instance = ImageSourceGlance.from_json(json)
# print the JSON string representation of the object
print(ImageSourceGlance.to_json())

# convert the object into a dict
image_source_glance_dict = image_source_glance_instance.to_dict()
# create an instance of ImageSourceGlance from a dict
image_source_glance_from_dict = ImageSourceGlance.from_dict(image_source_glance_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


