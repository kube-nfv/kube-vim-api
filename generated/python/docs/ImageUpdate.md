# ImageUpdate

Changes to a software image: a patch. The name is replaced when set; metadata entries are merged.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**metadata** | **Dict[str, str]** | Metadata entries to add, or to overwrite when the key exists; keys not given are kept, and none is removed. | [optional] 
**name** | **str** |  | [optional] 

## Example

```python
from kubevim_client.models.image_update import ImageUpdate

# TODO update the JSON string below
json = "{}"
# create an instance of ImageUpdate from a JSON string
image_update_instance = ImageUpdate.from_json(json)
# print the JSON string representation of the object
print(ImageUpdate.to_json())

# convert the object into a dict
image_update_dict = image_update_instance.to_dict()
# create an instance of ImageUpdate from a dict
image_update_from_dict = ImageUpdate.from_dict(image_update_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


