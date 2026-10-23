# CreateVolumeRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**request_id** | **str** | Client-chosen UUID that makes the request idempotent (AIP-155). | [optional] 
**volume** | [**Volume**](Volume.md) | Volume to create: name, size, type, storage class and metadata. Only TYPE_OF_STORAGE_VOLUME, without is_boot: ephemeral and boot disks belong to a compute resource. | 

## Example

```python
from kubevim_client.models.create_volume_request import CreateVolumeRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CreateVolumeRequest from a JSON string
create_volume_request_instance = CreateVolumeRequest.from_json(json)
# print the JSON string representation of the object
print(CreateVolumeRequest.to_json())

# convert the object into a dict
create_volume_request_dict = create_volume_request_instance.to_dict()
# create an instance of CreateVolumeRequest from a dict
create_volume_request_from_dict = CreateVolumeRequest.from_dict(create_volume_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


