# VirtualInterfaceData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**network_id** | [**Identifier**](Identifier.md) |  | [optional] 
**ip_address** | [**IPAddress**](IPAddress.md) |  | [optional] 
**mac_address** | [**MacAddress**](MacAddress.md) |  | [optional] 

## Example

```python
from vivnfm_client.models.virtual_interface_data import VirtualInterfaceData

# TODO update the JSON string below
json = "{}"
# create an instance of VirtualInterfaceData from a JSON string
virtual_interface_data_instance = VirtualInterfaceData.from_json(json)
# print the JSON string representation of the object
print(VirtualInterfaceData.to_json())

# convert the object into a dict
virtual_interface_data_dict = virtual_interface_data_instance.to_dict()
# create an instance of VirtualInterfaceData from a dict
virtual_interface_data_from_dict = VirtualInterfaceData.from_dict(virtual_interface_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


