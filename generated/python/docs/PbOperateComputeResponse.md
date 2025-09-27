# PbOperateComputeResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compute_data** | [**VirtualCompute**](VirtualCompute.md) |  | 
**compute_operation_output_data** | **Dict[str, str]** | Set of output values depending on the type of operation. For instance, when a snapshot operation is requested, this field provides information about the identifier of the snapshot and its location. | [optional] 

## Example

```python
from kubevim_vivnfm_client.models.pb_operate_compute_response import PbOperateComputeResponse

# TODO update the JSON string below
json = "{}"
# create an instance of PbOperateComputeResponse from a JSON string
pb_operate_compute_response_instance = PbOperateComputeResponse.from_json(json)
# print the JSON string representation of the object
print(PbOperateComputeResponse.to_json())

# convert the object into a dict
pb_operate_compute_response_dict = pb_operate_compute_response_instance.to_dict()
# create an instance of PbOperateComputeResponse from a dict
pb_operate_compute_response_from_dict = PbOperateComputeResponse.from_dict(pb_operate_compute_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


