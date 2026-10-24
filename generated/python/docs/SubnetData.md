# SubnetData

Layer 3 attributes of a subnet: one IP version of its network.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**address_pools** | [**List[IpAddressPool]**](IpAddressPool.md) | Address pools the VIM allocates from; the whole CIDR except the gateway when empty. | [optional] 
**cidr** | **str** | Network address and prefix length, e.g. \&quot;10.0.0.0/24\&quot;. | 
**dns_servers** | **List[str]** | DNS servers handed out by DHCP. | [optional] 
**gateway_ip** | **str** | Gateway address; the VIM picks one when unset. | [optional] 
**ip_version** | [**IpVersion**](IpVersion.md) | IP version. | 
**is_dhcp_enabled** | **bool** | DHCP is enabled; unset means true. | [optional] 

## Example

```python
from kubevim_client.models.subnet_data import SubnetData

# TODO update the JSON string below
json = "{}"
# create an instance of SubnetData from a JSON string
subnet_data_instance = SubnetData.from_json(json)
# print the JSON string representation of the object
print(SubnetData.to_json())

# convert the object into a dict
subnet_data_dict = subnet_data_instance.to_dict()
# create an instance of SubnetData from a dict
subnet_data_from_dict = SubnetData.from_dict(subnet_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


