# IpAddressPool

Range of addresses a subnet allocates from.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**end_ip** | **str** | Last address of the pool. | 
**start_ip** | **str** | First address of the pool. | 

## Example

```python
from kubevim_client.models.ip_address_pool import IpAddressPool

# TODO update the JSON string below
json = "{}"
# create an instance of IpAddressPool from a JSON string
ip_address_pool_instance = IpAddressPool.from_json(json)
# print the JSON string representation of the object
print(IpAddressPool.to_json())

# convert the object into a dict
ip_address_pool_dict = ip_address_pool_instance.to_dict()
# create an instance of IpAddressPool from a dict
ip_address_pool_from_dict = IpAddressPool.from_dict(ip_address_pool_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


