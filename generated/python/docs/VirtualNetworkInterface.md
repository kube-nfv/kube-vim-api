# VirtualNetworkInterface

A network interface of a compute resource.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**bandwidth** | **int** | Rate limit in Mbps; 0 when not limited. | [optional] 
**compute_id** | **str** | Compute resource the interface belongs to. | [optional] 
**guest_interface_name** | **str** | Interface name in the guest, e.g. \&quot;eth1\&quot;; empty when the guest agent does not report it. | [optional] 
**guest_pci_address** | **str** | PCI address of the device in the guest; empty until known. | [optional] 
**host_pci_address** | **str** | PCI address of the device on the host, for SR-IOV interfaces. | [optional] 
**ip_addresses** | **List[str]** | Addresses of the interface, one per subnet of its network; may be empty until DHCP or the guest configures them. | [optional] 
**is_management** | **bool** | The management interface every compute resource gets on the pod network. | [optional] 
**mac_address** | **str** | MAC address; empty until known. | [optional] 
**model** | [**NicModel**](NicModel.md) | NIC model the guest sees; unspecified for SRIOV and MASQUERADE interfaces. | [optional] 
**name** | **str** | Name of the interface, unique within the compute resource: the one given at creation, or one the VIM chose. | [optional] 
**network_id** | **str** | Network the interface is attached to; empty for the management interface. | [optional] 
**operational_state** | [**OperationalState**](OperationalState.md) | ENABLED once the interface is attached and ready. | [optional] 
**type_virtual_nic** | [**TypeVirtualNic**](TypeVirtualNic.md) |  | [optional] 

## Example

```python
from kubevim_client.models.virtual_network_interface import VirtualNetworkInterface

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


