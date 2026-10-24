# S3AccessKey

Access key for an S3-compatible object store.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**access_key_id** | **str** |  | 
**secret_access_key** | **str** |  | 

## Example

```python
from kubevim_client.models.s3_access_key import S3AccessKey

# TODO update the JSON string below
json = "{}"
# create an instance of S3AccessKey from a JSON string
s3_access_key_instance = S3AccessKey.from_json(json)
# print the JSON string representation of the object
print(S3AccessKey.to_json())

# convert the object into a dict
s3_access_key_dict = s3_access_key_instance.to_dict()
# create an instance of S3AccessKey from a dict
s3_access_key_from_dict = S3AccessKey.from_dict(s3_access_key_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


