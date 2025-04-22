# VirtualNetworkInterface

A virtual network interface resource is a communication endpoint under an instantiated compute resource.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**resource_id** | [**Identifier**](Identifier.md) |  | 
**network_id** | [**Identifier**](Identifier.md) |  | [optional] 
**subnet_id** | [**Identifier**](Identifier.md) |  | [optional] 
**network_port_id** | [**Identifier**](Identifier.md) |  | [optional] 
**ip_address** | [**List[IPAddress]**](IPAddress.md) | The virtual network interface can be configured with specific IP address(es) associated to the network to be attached to. The cardinality can be 0 in the case that a network interface is created without being attached to any specific network, or when an IP address can be automatically configured, e.g. by DHCP. Note(dmalovan): In general IPaddresses should be passed even if them are allocated dynamically e.g. via DHCP. If DHCP client not yet aquired the address cardinality can be 0. | [optional] 
**type_virtual_nic** | [**TypeVirtualNic**](TypeVirtualNic.md) |  | [default to TypeVirtualNic.BRIDGE]
**type_configuration** | **List[str]** | Extra configuration that the virtual network interface supports based on the type of virtual network interface. TODO: That interface might change. | [optional] 
**mac_address** | [**MacAddress**](MacAddress.md) |  | 
**bandwidth** | **float** | Bandwidth of the virtual network interface (in Mbps). | 
**acceleration_capability** | **List[str]** | The cardinality can be 0, if no particular acceleration capability is requested. TODO: That interface might change. | [optional] 
**operational_state** | [**OperationalState**](OperationalState.md) |  | [default to OperationalState.ENABLED]
**metadata** | [**Metadata**](Metadata.md) |  | [optional] 

## Example

```python
from kubevim_vivnfm_client.models.virtual_network_interface import VirtualNetworkInterface

# TODO update the JSON string below
json = "{}"
# create an instance of VirtualNetworkInterface from a JSON string
virtual_network_interface_instance = VirtualNetworkInterface.from_json(json)
# print the JSON string representation of the object
print(VirtualNetworkInterface.to_json())

# convert the object into a dict
virtual_network_interface_dict = virtual_network_interface_instance.to_dict()
# create an instance of VirtualNetworkInterface from a dict
virtual_network_interface_from_dict = VirtualNetworkInterface.from_dict(virtual_network_interface_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


