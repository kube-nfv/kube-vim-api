# IPAddressPool


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**start_ip** | [**IPAddress**](IPAddress.md) |  | [optional] 
**end_ip** | [**IPAddress**](IPAddress.md) |  | [optional] 

## Example

```python
from kubevim_vivnfm_client.models.ip_address_pool import IPAddressPool

# TODO update the JSON string below
json = "{}"
# create an instance of IPAddressPool from a JSON string
ip_address_pool_instance = IPAddressPool.from_json(json)
# print the JSON string representation of the object
print(IPAddressPool.to_json())

# convert the object into a dict
ip_address_pool_dict = ip_address_pool_instance.to_dict()
# create an instance of IPAddressPool from a dict
ip_address_pool_from_dict = IPAddressPool.from_dict(ip_address_pool_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


