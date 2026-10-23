# BlankSource

Empty disk.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**size** | **str** | Disk size in Kubernetes quantity format, e.g. \&quot;20Gi\&quot;. | 

## Example

```python
from kubevim_client.models.blank_source import BlankSource

# TODO update the JSON string below
json = "{}"
# create an instance of BlankSource from a JSON string
blank_source_instance = BlankSource.from_json(json)
# print the JSON string representation of the object
print(BlankSource.to_json())

# convert the object into a dict
blank_source_dict = blank_source_instance.to_dict()
# create an instance of BlankSource from a dict
blank_source_from_dict = BlankSource.from_dict(blank_source_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


