# IPAddress


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ip** | **str** |  | 

## Example

```python
from kubevim_vivnfm_client.models.ip_address import IPAddress

# TODO update the JSON string below
json = "{}"
# create an instance of IPAddress from a JSON string
ip_address_instance = IPAddress.from_json(json)
# print the JSON string representation of the object
print(IPAddress.to_json())

# convert the object into a dict
ip_address_dict = ip_address_instance.to_dict()
# create an instance of IPAddress from a dict
ip_address_from_dict = IPAddress.from_dict(ip_address_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


