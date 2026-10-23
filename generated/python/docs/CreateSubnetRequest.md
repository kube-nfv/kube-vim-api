# CreateSubnetRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**request_id** | **str** | Client-chosen UUID that makes the request idempotent (AIP-155). | [optional] 
**subnet** | [**Subnet**](Subnet.md) | Subnet to create; adds an IP version to a network that no interface uses yet. | 

## Example

```python
from kubevim_client.models.create_subnet_request import CreateSubnetRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CreateSubnetRequest from a JSON string
create_subnet_request_instance = CreateSubnetRequest.from_json(json)
# print the JSON string representation of the object
print(CreateSubnetRequest.to_json())

# convert the object into a dict
create_subnet_request_dict = create_subnet_request_instance.to_dict()
# create an instance of CreateSubnetRequest from a dict
create_subnet_request_from_dict = CreateSubnetRequest.from_dict(create_subnet_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


