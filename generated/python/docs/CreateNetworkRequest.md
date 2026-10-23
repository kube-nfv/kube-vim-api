# CreateNetworkRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**network** | [**Network**](Network.md) | Network to create. | 
**request_id** | **str** | Client-chosen UUID that makes the request idempotent (AIP-155). | [optional] 

## Example

```python
from kubevim_client.models.create_network_request import CreateNetworkRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CreateNetworkRequest from a JSON string
create_network_request_instance = CreateNetworkRequest.from_json(json)
# print the JSON string representation of the object
print(CreateNetworkRequest.to_json())

# convert the object into a dict
create_network_request_dict = create_network_request_instance.to_dict()
# create an instance of CreateNetworkRequest from a dict
create_network_request_from_dict = CreateNetworkRequest.from_dict(create_network_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


