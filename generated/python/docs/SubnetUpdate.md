# SubnetUpdate

Changes to a subnet: a patch. The name is replaced when set; metadata entries are merged.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**metadata** | **Dict[str, str]** | Metadata entries to add, or to overwrite when the key exists; keys not given are kept, and none is removed. | [optional] 
**subnet_name** | **str** |  | [optional] 

## Example

```python
from kubevim_client.models.subnet_update import SubnetUpdate

# TODO update the JSON string below
json = "{}"
# create an instance of SubnetUpdate from a JSON string
subnet_update_instance = SubnetUpdate.from_json(json)
# print the JSON string representation of the object
print(SubnetUpdate.to_json())

# convert the object into a dict
subnet_update_dict = subnet_update_instance.to_dict()
# create an instance of SubnetUpdate from a dict
subnet_update_from_dict = SubnetUpdate.from_dict(subnet_update_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


