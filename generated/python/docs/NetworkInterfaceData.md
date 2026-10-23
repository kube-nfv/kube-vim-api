# NetworkInterfaceData

A network interface requested for a compute resource.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**bandwidth** | **int** | Rate limit in Mbps; 0 means no limit. | [optional] 
**ip_addresses** | **List[str]** | Addresses for the static assignment: at most one per IP version; a family left out is allocated automatically. | [optional] 
**ip_assignment** | [**IpAssignment**](IpAssignment.md) | How the interface gets its addresses; automatic when unset. | [optional] 
**mac_address** | **str** | MAC address; the VIM assigns one when unset. | [optional] 
**model** | [**NicModel**](NicModel.md) | NIC model the guest sees; virtio when unset. | [optional] 
**name** | **str** | Name of the interface, unique within the compute resource; the VIM chooses one when unset. | [optional] 
**network_id** | **str** | Network to attach to; the interface gets one address per subnet of the network. | 
**type_virtual_nic** | [**TypeVirtualNic**](TypeVirtualNic.md) | How the interface is realised; the VIM chooses from the network when unset. | [optional] 

## Example

```python
from kubevim_client.models.network_interface_data import NetworkInterfaceData

# TODO update the JSON string below
json = "{}"
# create an instance of NetworkInterfaceData from a JSON string
network_interface_data_instance = NetworkInterfaceData.from_json(json)
# print the JSON string representation of the object
print(NetworkInterfaceData.to_json())

# convert the object into a dict
network_interface_data_dict = network_interface_data_instance.to_dict()
# create an instance of NetworkInterfaceData from a dict
network_interface_data_from_dict = NetworkInterfaceData.from_dict(network_interface_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


