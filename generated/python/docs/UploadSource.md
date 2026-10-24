# UploadSource

Upload by the client. Create the image with this source, then send the image file to upload_url with upload_token before upload_expire_time.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**upload_expire_time** | **datetime** | Time after which the token is no longer accepted. | [optional] [readonly] 
**upload_token** | **str** | Bearer token for the upload request. | [optional] [readonly] 
**upload_url** | **str** | URL to upload the image file to. | [optional] [readonly] 

## Example

```python
from kubevim_client.models.upload_source import UploadSource

# TODO update the JSON string below
json = "{}"
# create an instance of UploadSource from a JSON string
upload_source_instance = UploadSource.from_json(json)
# print the JSON string representation of the object
print(UploadSource.to_json())

# convert the object into a dict
upload_source_dict = upload_source_instance.to_dict()
# create an instance of UploadSource from a dict
upload_source_from_dict = UploadSource.from_dict(upload_source_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


