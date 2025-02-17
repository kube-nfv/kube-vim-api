# NetworkSubnetData

The NetworkSubnetData information element encapsulates information to allocate or update virtualised sub-networks.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**network_id** | [**Identifier**](Identifier.md) |  | [optional] 
**ip_version** | [**IPVersion**](IPVersion.md) |  | [optional] [default to IPVersion.IPV4]
**gateway_ip** | [**IPAddress**](IPAddress.md) |  | [optional] 
**cidr** | [**IPSubnetCIDR**](IPSubnetCIDR.md) |  | [optional] 
**is_dhcp_enabled** | **bool** | True when DHCP is to be enabled for this network/subnetwork, or false otherwise. | [optional] 
**address_pool** | [**IPAddressPool**](IPAddressPool.md) |  | [optional] 
**metadata** | [**Metadata**](Metadata.md) |  | [optional] 

## Example

```python
from vivnfm_client.models.network_subnet_data import NetworkSubnetData

# TODO update the JSON string below
json = "{}"
# create an instance of NetworkSubnetData from a JSON string
network_subnet_data_instance = NetworkSubnetData.from_json(json)
# print the JSON string representation of the object
print(NetworkSubnetData.to_json())

# convert the object into a dict
network_subnet_data_dict = network_subnet_data_instance.to_dict()
# create an instance of NetworkSubnetData from a dict
network_subnet_data_from_dict = NetworkSubnetData.from_dict(network_subnet_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


