# VolumeUpdate

Changes to a volume: a patch. The name is replaced when set; metadata entries are merged.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**metadata** | **Dict[str, str]** | Metadata entries to add, or to overwrite when the key exists; keys not given are kept, and none is removed. | [optional] 
**volume_name** | **str** |  | [optional] 

## Example

```python
from kubevim_client.models.volume_update import VolumeUpdate

# TODO update the JSON string below
json = "{}"
# create an instance of VolumeUpdate from a JSON string
volume_update_instance = VolumeUpdate.from_json(json)
# print the JSON string representation of the object
print(VolumeUpdate.to_json())

# convert the object into a dict
volume_update_dict = volume_update_instance.to_dict()
# create an instance of VolumeUpdate from a dict
volume_update_from_dict = VolumeUpdate.from_dict(volume_update_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


