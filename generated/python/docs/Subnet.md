# Subnet

A virtualised subnet: one IP version of a network.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**data** | [**SubnetData**](SubnetData.md) | Layer 3 attributes. | 
**metadata** | **Dict[str, str]** | Metadata set by the consumer and by the VIM; for keys under \&quot;kubevim.kubenfv.io/\&quot; that the VIM maintains, its value wins. | [optional] 
**network_id** | **str** | Network the subnet belongs to. | 
**subnet_id** | **str** | Identifier of the subnet: a UUID assigned by the VIM. | [optional] [readonly] 
**subnet_name** | **str** | Name of the subnet, given by the consumer; not unique. | [optional] 

## Example

```python
from kubevim_client.models.subnet import Subnet

# TODO update the JSON string below
json = "{}"
# create an instance of Subnet from a JSON string
subnet_instance = Subnet.from_json(json)
# print the JSON string representation of the object
print(Subnet.to_json())

# convert the object into a dict
subnet_dict = subnet_instance.to_dict()
# create an instance of Subnet from a dict
subnet_from_dict = Subnet.from_dict(subnet_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


