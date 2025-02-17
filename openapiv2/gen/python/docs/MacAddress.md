# MacAddress


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**mac** | **str** |  | [optional] 

## Example

```python
from vivnfm_client.models.mac_address import MacAddress

# TODO update the JSON string below
json = "{}"
# create an instance of MacAddress from a JSON string
mac_address_instance = MacAddress.from_json(json)
# print the JSON string representation of the object
print(MacAddress.to_json())

# convert the object into a dict
mac_address_dict = mac_address_instance.to_dict()
# create an instance of MacAddress from a dict
mac_address_from_dict = MacAddress.from_dict(mac_address_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


