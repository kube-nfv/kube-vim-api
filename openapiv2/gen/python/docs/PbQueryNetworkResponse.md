# PbQueryNetworkResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**query_network_result** | [**List[VirtualNetwork]**](VirtualNetwork.md) |  | [optional] 
**query_subnet_result** | [**List[NetworkSubnet]**](NetworkSubnet.md) |  | [optional] 
**query_network_port_result** | **List[object]** |  | [optional] 
**query_trunk_result** | **List[object]** |  | [optional] 

## Example

```python
from kubevim_vivnfm_client.models.pb_query_network_response import PbQueryNetworkResponse

# TODO update the JSON string below
json = "{}"
# create an instance of PbQueryNetworkResponse from a JSON string
pb_query_network_response_instance = PbQueryNetworkResponse.from_json(json)
# print the JSON string representation of the object
print(PbQueryNetworkResponse.to_json())

# convert the object into a dict
pb_query_network_response_dict = pb_query_network_response_instance.to_dict()
# create an instance of PbQueryNetworkResponse from a dict
pb_query_network_response_from_dict = PbQueryNetworkResponse.from_dict(pb_query_network_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


