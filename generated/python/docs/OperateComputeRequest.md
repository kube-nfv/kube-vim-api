# OperateComputeRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compute_operation** | [**ComputeOperation**](ComputeOperation.md) | Operation to run. | 

## Example

```python
from kubevim_client.models.operate_compute_request import OperateComputeRequest

# TODO update the JSON string below
json = "{}"
# create an instance of OperateComputeRequest from a JSON string
operate_compute_request_instance = OperateComputeRequest.from_json(json)
# print the JSON string representation of the object
print(OperateComputeRequest.to_json())

# convert the object into a dict
operate_compute_request_dict = operate_compute_request_instance.to_dict()
# create an instance of OperateComputeRequest from a dict
operate_compute_request_from_dict = OperateComputeRequest.from_dict(operate_compute_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


