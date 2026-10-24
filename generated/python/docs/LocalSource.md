# LocalSource

Image file in the local catalog directory configured in the VIM.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**file_name** | **str** | Name of the file in the catalog directory, without a path. | 

## Example

```python
from kubevim_client.models.local_source import LocalSource

# TODO update the JSON string below
json = "{}"
# create an instance of LocalSource from a JSON string
local_source_instance = LocalSource.from_json(json)
# print the JSON string representation of the object
print(LocalSource.to_json())

# convert the object into a dict
local_source_dict = local_source_instance.to_dict()
# create an instance of LocalSource from a dict
local_source_from_dict = LocalSource.from_dict(local_source_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


