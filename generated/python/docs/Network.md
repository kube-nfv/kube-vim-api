# Network

A virtualised network: one layer 2 segment, with at most one IPv4 and one IPv6 subnet.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**layer3_attributes** | [**List[SubnetData]**](SubnetData.md) | Subnets created with the network, in one call that rolls back as a whole: at most one per IP version. A Kube-OVN network needs at least one; GENEVE networks always use Kube-OVN. | [optional] 
**metadata** | **Dict[str, str]** | Metadata set by the consumer and by the VIM; for keys under \&quot;kubevim.kubenfv.io/\&quot; that the VIM maintains, its value wins. | [optional] 
**mtu** | **int** | MTU of the network in bytes; the VIM default when unset. | [optional] 
**network_id** | **str** | Identifier of the network: a UUID assigned by the VIM. | [optional] [readonly] 
**network_name** | **str** | Name of the network, given by the consumer; not unique. | [optional] 
**network_type** | [**NetworkType**](NetworkType.md) | How traffic of the network is separated. | 
**operational_state** | [**OperationalState**](OperationalState.md) | Operational state of the network. | [optional] [readonly] 
**provider_network** | **str** | Name of a provider network declared in the VIM configuration, which also sets its backend (Kube-OVN or SR-IOV); required for VLAN and FLAT networks, absent for GENEVE. | [optional] 
**segmentation_id** | **int** | VLAN ID of a VLAN network; absent for GENEVE and FLAT networks. | [optional] 
**subnet_ids** | **List[str]** | Subnets of the network: at most one per IP version. | [optional] [readonly] 

## Example

```python
from kubevim_client.models.network import Network

# TODO update the JSON string below
json = "{}"
# create an instance of Network from a JSON string
network_instance = Network.from_json(json)
# print the JSON string representation of the object
print(Network.to_json())

# convert the object into a dict
network_dict = network_instance.to_dict()
# create an instance of Network from a dict
network_from_dict = Network.from_dict(network_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


