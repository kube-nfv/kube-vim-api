# Checksum

Checksum of a file (SOL 013 7.1.7).

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**algorithm** | **str** | Hash algorithm, e.g. \&quot;SHA-256\&quot;. | [optional] [readonly] 
**hash** | **str** | Hexadecimal hash value. | [optional] [readonly] 

## Example

```python
from kubevim_client.models.checksum import Checksum

# TODO update the JSON string below
json = "{}"
# create an instance of Checksum from a JSON string
checksum_instance = Checksum.from_json(json)
# print the JSON string representation of the object
print(Checksum.to_json())

# convert the object into a dict
checksum_dict = checksum_instance.to_dict()
# create an instance of Checksum from a dict
checksum_from_dict = Checksum.from_dict(checksum_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


