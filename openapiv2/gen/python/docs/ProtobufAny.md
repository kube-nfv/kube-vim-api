# ProtobufAny


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | **str** |  | [optional] 

## Example

```python
from kubevim_vivnfm_client.models.protobuf_any import ProtobufAny

# TODO update the JSON string below
json = "{}"
# create an instance of ProtobufAny from a JSON string
protobuf_any_instance = ProtobufAny.from_json(json)
# print the JSON string representation of the object
print(ProtobufAny.to_json())

# convert the object into a dict
protobuf_any_dict = protobuf_any_instance.to_dict()
# create an instance of ProtobufAny from a dict
protobuf_any_from_dict = ProtobufAny.from_dict(protobuf_any_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


