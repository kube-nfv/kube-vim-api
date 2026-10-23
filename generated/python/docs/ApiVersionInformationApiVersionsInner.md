# ApiVersionInformationApiVersionsInner


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**is_deprecated** | **bool** |  | [optional] 
**retirement_date** | **datetime** |  | [optional] 
**version** | **str** | Version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0. | 

## Example

```python
from kubevim_client.models.api_version_information_api_versions_inner import ApiVersionInformationApiVersionsInner

# TODO update the JSON string below
json = "{}"
# create an instance of ApiVersionInformationApiVersionsInner from a JSON string
api_version_information_api_versions_inner_instance = ApiVersionInformationApiVersionsInner.from_json(json)
# print the JSON string representation of the object
print(ApiVersionInformationApiVersionsInner.to_json())

# convert the object into a dict
api_version_information_api_versions_inner_dict = api_version_information_api_versions_inner_instance.to_dict()
# create an instance of ApiVersionInformationApiVersionsInner from a dict
api_version_information_api_versions_inner_from_dict = ApiVersionInformationApiVersionsInner.from_dict(api_version_information_api_versions_inner_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


