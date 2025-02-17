# Metadata

List of metadata key-value pairs used by the consumer to associate meaningful metadata to the related virtualised resource. Note: metadata represent as a string, string map. Make value of pb.Any is not mianingfull since consumer won't be able to deserialize it.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**fields** | **Dict[str, str]** |  | [optional] 

## Example

```python
from vivnfm_client.models.metadata import Metadata

# TODO update the JSON string below
json = "{}"
# create an instance of Metadata from a JSON string
metadata_instance = Metadata.from_json(json)
# print the JSON string representation of the object
print(Metadata.to_json())

# convert the object into a dict
metadata_dict = metadata_instance.to_dict()
# create an instance of Metadata from a dict
metadata_from_dict = Metadata.from_dict(metadata_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


