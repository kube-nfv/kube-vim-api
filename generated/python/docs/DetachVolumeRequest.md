# DetachVolumeRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**volume_id** | **str** | Identifier of the volume to detach. | 

## Example

```python
from kubevim_client.models.detach_volume_request import DetachVolumeRequest

# TODO update the JSON string below
json = "{}"
# create an instance of DetachVolumeRequest from a JSON string
detach_volume_request_instance = DetachVolumeRequest.from_json(json)
# print the JSON string representation of the object
print(DetachVolumeRequest.to_json())

# convert the object into a dict
detach_volume_request_dict = detach_volume_request_instance.to_dict()
# create an instance of DetachVolumeRequest from a dict
detach_volume_request_from_dict = DetachVolumeRequest.from_dict(detach_volume_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


