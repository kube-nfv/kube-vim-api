# HttpSource

HTTP or HTTPS download. Only one of basicAuth, credentialsRef may be set.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**basic_auth** | [**BasicAuth**](BasicAuth.md) | Inline username and password; never returned. | [optional] 
**ca_bundle** | **str** | PEM CA bundle to verify the server certificate. | [optional] 
**credentials_ref** | **str** | Name of a credential set the VIM loads from its own environment or files; the request carries no secret. | [optional] 
**extra_headers** | **List[str]** | Extra request headers, as \&quot;Name: value\&quot;; never returned, since they may carry secrets. | [optional] 
**url** | **str** | URL of the image file. | 

## Example

```python
from kubevim_client.models.http_source import HttpSource

# TODO update the JSON string below
json = "{}"
# create an instance of HttpSource from a JSON string
http_source_instance = HttpSource.from_json(json)
# print the JSON string representation of the object
print(HttpSource.to_json())

# convert the object into a dict
http_source_dict = http_source_instance.to_dict()
# create an instance of HttpSource from a dict
http_source_from_dict = HttpSource.from_dict(http_source_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


