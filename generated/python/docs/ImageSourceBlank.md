# ImageSourceBlank


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**blank** | [**BlankSource**](BlankSource.md) | Create an empty disk. | 

## Example

```python
from kubevim_client.models.image_source_blank import ImageSourceBlank

# TODO update the JSON string below
json = "{}"
# create an instance of ImageSourceBlank from a JSON string
image_source_blank_instance = ImageSourceBlank.from_json(json)
# print the JSON string representation of the object
print(ImageSourceBlank.to_json())

# convert the object into a dict
image_source_blank_dict = image_source_blank_instance.to_dict()
# create an instance of ImageSourceBlank from a dict
image_source_blank_from_dict = ImageSourceBlank.from_dict(image_source_blank_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


