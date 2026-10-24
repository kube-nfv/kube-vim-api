# RegistrySource

Container registry pull. Only one of basicAuth, credentialsRef may be set.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**basic_auth** | [**BasicAuth**](BasicAuth.md) | Inline username and password; never returned. | [optional] 
**ca_bundle** | **str** | PEM CA bundle to verify the registry certificate. | [optional] 
**credentials_ref** | **str** | Name of a credential set the VIM loads from its own environment or files; the request carries no secret. | [optional] 
**platform** | **str** | Platform of a multi-architecture image, e.g. \&quot;linux/arm64\&quot;. | [optional] 
**pull_method** | [**RegistrySourcePullMethod**](RegistrySourcePullMethod.md) | Pull method; the VIM default when unset. | [optional] 
**url** | **str** | Image reference, e.g. \&quot;docker://registry.example.com/images/ubuntu:24.04\&quot;. | 

## Example

```python
from kubevim_client.models.registry_source import RegistrySource

# TODO update the JSON string below
json = "{}"
# create an instance of RegistrySource from a JSON string
registry_source_instance = RegistrySource.from_json(json)
# print the JSON string representation of the object
print(RegistrySource.to_json())

# convert the object into a dict
registry_source_dict = registry_source_instance.to_dict()
# create an instance of RegistrySource from a dict
registry_source_from_dict = RegistrySource.from_dict(registry_source_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


