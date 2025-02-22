# PbAllocateComputeResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compute_data** | [**VirtualCompute**](VirtualCompute.md) |  | [optional] 

## Example

```python
from kubevim_vivnfm_client.models.pb_allocate_compute_response import PbAllocateComputeResponse

# TODO update the JSON string below
json = "{}"
# create an instance of PbAllocateComputeResponse from a JSON string
pb_allocate_compute_response_instance = PbAllocateComputeResponse.from_json(json)
# print the JSON string representation of the object
print(PbAllocateComputeResponse.to_json())

# convert the object into a dict
pb_allocate_compute_response_dict = pb_allocate_compute_response_instance.to_dict()
# create an instance of PbAllocateComputeResponse from a dict
pb_allocate_compute_response_from_dict = PbAllocateComputeResponse.from_dict(pb_allocate_compute_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


