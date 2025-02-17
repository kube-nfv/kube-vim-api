# VirtualNetworkData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**bandwidth** | **float** | Minimum network bandwidth (in Mbps). | [optional] 
**network_type** | **str** | Type of network that maps to the virtualised network. This list is extensible. Examples are:  • \&quot;local\&quot;;  • \&quot;vlan\&quot;;  • \&quot;vxlan\&quot;;  • \&quot;gre\&quot;;  • \&quot;l3-vpn\&quot;;  • etc.  Cardinality can be \&quot;0\&quot; to cover the case where this attribute is not required to create the virtualised network. | [optional] 
**segmentation_id** | **str** | The segmentation identifier of the network that maps to the virtualised network, for which, the segmentation model is defined by the networkType attribute. For instance, for a \&quot;vlan\&quot; networkType, it corresponds to the vlan identifier; and for a \&quot;gre\&quot; networkType, it corresponds to a gre key. Cardinality can be \&quot;0\&quot; to cover the case where networkType is flat network without any specific segmentation. | [optional] 
**network_qo_s** | [**List[NetworkQoS]**](NetworkQoS.md) | Provides information about Quality of Service attributes that the network shall support. Cardinality can be \&quot;0\&quot; for networks without any specified QoS requirements. | [optional] 
**is_shared** | **bool** | Specifies whether the virtualised network is shared among consumers. | [optional] 
**layer3_attributes** | [**List[NetworkSubnetData]**](NetworkSubnetData.md) | Attribute list allows setting up a network providing defined layer 3 connectivity. | [optional] 
**metadata** | [**Metadata**](Metadata.md) |  | [optional] 

## Example

```python
from kubevim_vivnfm_client.models.virtual_network_data import VirtualNetworkData

# TODO update the JSON string below
json = "{}"
# create an instance of VirtualNetworkData from a JSON string
virtual_network_data_instance = VirtualNetworkData.from_json(json)
# print the JSON string representation of the object
print(VirtualNetworkData.to_json())

# convert the object into a dict
virtual_network_data_dict = virtual_network_data_instance.to_dict()
# create an instance of VirtualNetworkData from a dict
virtual_network_data_from_dict = VirtualNetworkData.from_dict(virtual_network_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


