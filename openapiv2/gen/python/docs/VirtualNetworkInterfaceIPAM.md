# VirtualNetworkInterfaceIPAM


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**network_id** | [**Identifier**](Identifier.md) |  | [optional] 
**subnet_id** | [**Identifier**](Identifier.md) |  | [optional] 
**ip_address** | [**IPAddress**](IPAddress.md) |  | [optional] 
**mac_address** | [**MacAddress**](MacAddress.md) |  | [optional] 

## Example

```python
from kubevim_vivnfm_client.models.virtual_network_interface_ipam import VirtualNetworkInterfaceIPAM

# TODO update the JSON string below
json = "{}"
# create an instance of VirtualNetworkInterfaceIPAM from a JSON string
virtual_network_interface_ipam_instance = VirtualNetworkInterfaceIPAM.from_json(json)
# print the JSON string representation of the object
print(VirtualNetworkInterfaceIPAM.to_json())

# convert the object into a dict
virtual_network_interface_ipam_dict = virtual_network_interface_ipam_instance.to_dict()
# create an instance of VirtualNetworkInterfaceIPAM from a dict
virtual_network_interface_ipam_from_dict = VirtualNetworkInterfaceIPAM.from_dict(virtual_network_interface_ipam_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


