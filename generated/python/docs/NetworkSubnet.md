# NetworkSubnet

The NetworkSubnet information element encapsulates information of an instantiated virtualised sub-network.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**resource_id** | [**Identifier**](Identifier.md) |  | 
**network_id** | [**Identifier**](Identifier.md) |  | [optional] 
**ip_version** | [**IPVersion**](IPVersion.md) |  | [optional] [default to IPVersion.IPV4]
**gateway_ip** | [**IPAddress**](IPAddress.md) |  | 
**cidr** | [**IPSubnetCIDR**](IPSubnetCIDR.md) |  | 
**is_dhcp_enabled** | **bool** | True when DHCP is enabled for this network/subnetwork, or false otherwise. | 
**address_pool** | [**List[IPAddressPool]**](IPAddressPool.md) | Address pools for the network/subnetwork. The cardinality can be 0 when VIM is allowed to allocate all addresses in the CIDR except for the address of the network/subnetwork gateway. | [optional] 
**metadata** | [**Metadata**](Metadata.md) |  | [optional] 

## Example

```python
from kubevim_vivnfm_client.models.network_subnet import NetworkSubnet

# TODO update the JSON string below
json = "{}"
# create an instance of NetworkSubnet from a JSON string
network_subnet_instance = NetworkSubnet.from_json(json)
# print the JSON string representation of the object
print(NetworkSubnet.to_json())

# convert the object into a dict
network_subnet_dict = network_subnet_instance.to_dict()
# create an instance of NetworkSubnet from a dict
network_subnet_from_dict = NetworkSubnet.from_dict(network_subnet_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


