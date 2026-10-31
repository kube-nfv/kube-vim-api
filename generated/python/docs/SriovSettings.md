# SriovSettings

Settings for SRIOV interfaces on a network.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**pool** | **str** | Name of the pool of SR-IOV virtual functions on the provider network&#39;s uplink, known to the VIM. | 

## Example

```python
from kubevim_client.models.sriov_settings import SriovSettings

# TODO update the JSON string below
json = "{}"
# create an instance of SriovSettings from a JSON string
sriov_settings_instance = SriovSettings.from_json(json)
# print the JSON string representation of the object
print(SriovSettings.to_json())

# convert the object into a dict
sriov_settings_dict = sriov_settings_instance.to_dict()
# create an instance of SriovSettings from a dict
sriov_settings_from_dict = SriovSettings.from_dict(sriov_settings_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


