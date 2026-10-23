# ApiVersionInformation

API versions served under a URI prefix (SOL 013 9.3.3).

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**api_versions** | [**List[ApiVersionInformationApiVersionsInner]**](ApiVersionInformationApiVersionsInner.md) |  | 
**uri_prefix** | **str** | URI prefix of the API, e.g. /compute/v1. | 

## Example

```python
from kubevim_client.models.api_version_information import ApiVersionInformation

# TODO update the JSON string below
json = "{}"
# create an instance of ApiVersionInformation from a JSON string
api_version_information_instance = ApiVersionInformation.from_json(json)
# print the JSON string representation of the object
print(ApiVersionInformation.to_json())

# convert the object into a dict
api_version_information_dict = api_version_information_instance.to_dict()
# create an instance of ApiVersionInformation from a dict
api_version_information_from_dict = ApiVersionInformation.from_dict(api_version_information_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


