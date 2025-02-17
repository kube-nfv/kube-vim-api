# PbTerminateNetworkResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**network_resource_id** | [**Identifier**](Identifier.md) |  | [optional] 

## Example

```python
from vivnfm_client.models.pb_terminate_network_response import PbTerminateNetworkResponse

# TODO update the JSON string below
json = "{}"
# create an instance of PbTerminateNetworkResponse from a JSON string
pb_terminate_network_response_instance = PbTerminateNetworkResponse.from_json(json)
# print the JSON string representation of the object
print(PbTerminateNetworkResponse.to_json())

# convert the object into a dict
pb_terminate_network_response_dict = pb_terminate_network_response_instance.to_dict()
# create an instance of PbTerminateNetworkResponse from a dict
pb_terminate_network_response_from_dict = PbTerminateNetworkResponse.from_dict(pb_terminate_network_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


