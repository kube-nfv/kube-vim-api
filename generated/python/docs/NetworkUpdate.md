# NetworkUpdate

Changes to a network: a patch. The name is replaced when set; metadata entries are merged.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**metadata** | **Dict[str, str]** | Metadata entries to add, or to overwrite when the key exists; keys not given are kept, and none is removed. | [optional] 
**network_name** | **str** |  | [optional] 

## Example

```python
from kubevim_client.models.network_update import NetworkUpdate

# TODO update the JSON string below
json = "{}"
# create an instance of NetworkUpdate from a JSON string
network_update_instance = NetworkUpdate.from_json(json)
# print the JSON string representation of the object
print(NetworkUpdate.to_json())

# convert the object into a dict
network_update_dict = network_update_instance.to_dict()
# create an instance of NetworkUpdate from a dict
network_update_from_dict = NetworkUpdate.from_dict(network_update_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


