# VirtualNetworkInterfaceData

A virtual network interface is a communication endpoint under a compute resource.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**network_id** | [**Identifier**](Identifier.md) |  | [optional] 
**network_port_id** | [**Identifier**](Identifier.md) |  | [optional] 
**bandwidth** | **float** | Note: There is only part of flavour as specified in ETSI GS NFV-IFA 005 and ETSI GS NFV-IFA 006 are included in this version of the present document, the following are attributes not included: typeVirtualNic, typeConfiguration typeVirtualNic &#x3D; 3; typeConfiguration &#x3D; 4; Bandwidth of the virtual network interface (in Mbps). | [optional] 
**acceleration_capability_for_virtual_network_interface** | **List[str]** | It specifies if the virtual network interface requires certain acceleration capabilities (e.g. RDMA, packet dispatch, TCP Chimney). The cardinality can be 0, if no particular acceleration capability is requested. | [optional] 
**metadata** | [**Metadata**](Metadata.md) |  | [optional] 

## Example

```python
from vivnfm_client.models.virtual_network_interface_data import VirtualNetworkInterfaceData

# TODO update the JSON string below
json = "{}"
# create an instance of VirtualNetworkInterfaceData from a JSON string
virtual_network_interface_data_instance = VirtualNetworkInterfaceData.from_json(json)
# print the JSON string representation of the object
print(VirtualNetworkInterfaceData.to_json())

# convert the object into a dict
virtual_network_interface_data_dict = virtual_network_interface_data_instance.to_dict()
# create an instance of VirtualNetworkInterfaceData from a dict
virtual_network_interface_data_from_dict = VirtualNetworkInterfaceData.from_dict(virtual_network_interface_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


