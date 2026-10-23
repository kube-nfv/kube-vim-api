# GlanceSource

Image in the Glance catalog configured in the VIM; the VIM holds the Glance endpoint and credentials.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**glance_image_id** | **str** | Identifier of the image in Glance. | 

## Example

```python
from kubevim_client.models.glance_source import GlanceSource

# TODO update the JSON string below
json = "{}"
# create an instance of GlanceSource from a JSON string
glance_source_instance = GlanceSource.from_json(json)
# print the JSON string representation of the object
print(GlanceSource.to_json())

# convert the object into a dict
glance_source_dict = glance_source_instance.to_dict()
# create an instance of GlanceSource from a dict
glance_source_from_dict = GlanceSource.from_dict(glance_source_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


