# S3Source

S3-compatible object store download. Only one of accessKey, credentialsRef may be set.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**access_key** | [**S3AccessKey**](S3AccessKey.md) | Inline access key; never returned. | [optional] 
**ca_bundle** | **str** | PEM CA bundle to verify the endpoint certificate. | [optional] 
**credentials_ref** | **str** | Name of a credential set the VIM loads from its own environment or files; the request carries no secret. | [optional] 
**url** | **str** | URL of the object, e.g. \&quot;https://s3.example.com/bucket/image.qcow2\&quot;. | 

## Example

```python
from kubevim_client.models.s3_source import S3Source

# TODO update the JSON string below
json = "{}"
# create an instance of S3Source from a JSON string
s3_source_instance = S3Source.from_json(json)
# print the JSON string representation of the object
print(S3Source.to_json())

# convert the object into a dict
s3_source_dict = s3_source_instance.to_dict()
# create an instance of S3Source from a dict
s3_source_from_dict = S3Source.from_dict(s3_source_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


