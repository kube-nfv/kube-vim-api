# VirtualNetwork


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**network_resource_id** | [**Identifier**](Identifier.md) |  | [optional] 
**network_resource_name** | **str** | Name of the virtualised network resource. | [optional] 
**subnet_id** | [**List[Identifier]**](Identifier.md) | References the network subnet. Only present if the network provides layer 3 connectivity. | [optional] 
**network_port** | **List[object]** | Provides information on an instantiated virtual network port. | [optional] 
**bandwidth** | **float** |  | [optional] 
**network_type** | [**NetworkType**](NetworkType.md) |  | [optional] [default to NetworkType.OVERLAY]
**provider_network** | **str** | Name of the infrastructure provider network used to realize the virtual network. Cardinality can be \&quot;0\&quot; to cover the case where virtual network is not based on infrastructure provider network. | [optional] 
**segmentation_id** | **str** | The segmentation identifier of the network that maps to the virtualised network, for which, the segmentation model is defined by the networkType attribute. For instance, for a \&quot;vlan\&quot; networkType, it corresponds to the vlan identifier; and for a \&quot;gre\&quot; networkType, it corresponds to a gre key. Cardinality can be \&quot;0\&quot; to cover the case where networkType is flat network without any specific segmentation. | [optional] 
**network_qo_s** | [**List[NetworkQoS]**](NetworkQoS.md) | Provides information about Quality of Service attributes that the network supports. Cardinality can be \&quot;0\&quot; for virtual network without any QoS requirements. | [optional] 
**is_shared** | **bool** | Defines whether the virtualised network is shared among consumers. | [optional] 
**zone_id** | [**Identifier**](Identifier.md) |  | [optional] 
**operational_state** | [**OperationalState**](OperationalState.md) |  | [optional] [default to OperationalState.ENABLED]
**metadata** | [**Metadata**](Metadata.md) |  | [optional] 
**connected_networks** | [**List[Identifier]**](Identifier.md) | Specifies the virtual network resources to which the newly created virtual network is intended to be explicitly interconnected. | [optional] 

## Example

```python
from kubevim_vivnfm_client.models.virtual_network import VirtualNetwork

# TODO update the JSON string below
json = "{}"
# create an instance of VirtualNetwork from a JSON string
virtual_network_instance = VirtualNetwork.from_json(json)
# print the JSON string representation of the object
print(VirtualNetwork.to_json())

# convert the object into a dict
virtual_network_dict = virtual_network_instance.to_dict()
# create an instance of VirtualNetwork from a dict
virtual_network_from_dict = VirtualNetwork.from_dict(virtual_network_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


