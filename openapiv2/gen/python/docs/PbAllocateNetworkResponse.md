# PbAllocateNetworkResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**network_data** | [**VirtualNetwork**](VirtualNetwork.md) |  | [optional] 
**subnet_data** | [**NetworkSubnet**](NetworkSubnet.md) |  | [optional] 
**network_port_data** | **object** |  | [optional] 
**trunk_data** | **object** |  | [optional] 
**routing_resource_data** | **object** |  | [optional] 

## Example

```python
from vivnfm_client.models.pb_allocate_network_response import PbAllocateNetworkResponse

# TODO update the JSON string below
json = "{}"
# create an instance of PbAllocateNetworkResponse from a JSON string
pb_allocate_network_response_instance = PbAllocateNetworkResponse.from_json(json)
# print the JSON string representation of the object
print(PbAllocateNetworkResponse.to_json())

# convert the object into a dict
pb_allocate_network_response_dict = pb_allocate_network_response_instance.to_dict()
# create an instance of PbAllocateNetworkResponse from a dict
pb_allocate_network_response_from_dict = PbAllocateNetworkResponse.from_dict(pb_allocate_network_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


