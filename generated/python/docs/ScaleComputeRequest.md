# ScaleComputeRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**flavour_id** | **str** | Flavour to resize to. | 

## Example

```python
from kubevim_client.models.scale_compute_request import ScaleComputeRequest

# TODO update the JSON string below
json = "{}"
# create an instance of ScaleComputeRequest from a JSON string
scale_compute_request_instance = ScaleComputeRequest.from_json(json)
# print the JSON string representation of the object
print(ScaleComputeRequest.to_json())

# convert the object into a dict
scale_compute_request_dict = scale_compute_request_instance.to_dict()
# create an instance of ScaleComputeRequest from a dict
scale_compute_request_from_dict = ScaleComputeRequest.from_dict(scale_compute_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


