# PbTerminateComputeResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compute_id** | [**Identifier**](Identifier.md) |  | 

## Example

```python
from kubevim_vivnfm_client.models.pb_terminate_compute_response import PbTerminateComputeResponse

# TODO update the JSON string below
json = "{}"
# create an instance of PbTerminateComputeResponse from a JSON string
pb_terminate_compute_response_instance = PbTerminateComputeResponse.from_json(json)
# print the JSON string representation of the object
print(PbTerminateComputeResponse.to_json())

# convert the object into a dict
pb_terminate_compute_response_dict = pb_terminate_compute_response_instance.to_dict()
# create an instance of PbTerminateComputeResponse from a dict
pb_terminate_compute_response_from_dict = PbTerminateComputeResponse.from_dict(pb_terminate_compute_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


