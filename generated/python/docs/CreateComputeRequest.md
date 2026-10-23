# CreateComputeRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compute** | [**Compute**](Compute.md) | Compute resource to create: name, flavour, image, placement groups and metadata. | 
**network_interfaces** | [**List[NetworkInterfaceData]**](NetworkInterfaceData.md) | Network interfaces in addition to the management interface. | [optional] 
**start** | **bool** | Start the compute resource once created; unset means true. | [optional] 
**user_data** | [**UserData**](UserData.md) | User data passed at boot time. | [optional] 

## Example

```python
from kubevim_client.models.create_compute_request import CreateComputeRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CreateComputeRequest from a JSON string
create_compute_request_instance = CreateComputeRequest.from_json(json)
# print the JSON string representation of the object
print(CreateComputeRequest.to_json())

# convert the object into a dict
create_compute_request_dict = create_compute_request_instance.to_dict()
# create an instance of CreateComputeRequest from a dict
create_compute_request_from_dict = CreateComputeRequest.from_dict(create_compute_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


