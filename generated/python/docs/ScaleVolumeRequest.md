# ScaleVolumeRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**new_size** | **str** | New size in Kubernetes quantity format; volumes can only grow (persistent volume claim expansion). | 

## Example

```python
from kubevim_client.models.scale_volume_request import ScaleVolumeRequest

# TODO update the JSON string below
json = "{}"
# create an instance of ScaleVolumeRequest from a JSON string
scale_volume_request_instance = ScaleVolumeRequest.from_json(json)
# print the JSON string representation of the object
print(ScaleVolumeRequest.to_json())

# convert the object into a dict
scale_volume_request_dict = scale_volume_request_instance.to_dict()
# create an instance of ScaleVolumeRequest from a dict
scale_volume_request_from_dict = ScaleVolumeRequest.from_dict(scale_volume_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


