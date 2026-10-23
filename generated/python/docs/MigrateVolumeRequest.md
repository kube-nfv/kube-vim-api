# MigrateVolumeRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**storage_class** | **str** | Storage class to move the volume to, one labelled kubevim.kubenfv.io/storage&#x3D;true. | 

## Example

```python
from kubevim_client.models.migrate_volume_request import MigrateVolumeRequest

# TODO update the JSON string below
json = "{}"
# create an instance of MigrateVolumeRequest from a JSON string
migrate_volume_request_instance = MigrateVolumeRequest.from_json(json)
# print the JSON string representation of the object
print(MigrateVolumeRequest.to_json())

# convert the object into a dict
migrate_volume_request_dict = migrate_volume_request_instance.to_dict()
# create an instance of MigrateVolumeRequest from a dict
migrate_volume_request_from_dict = MigrateVolumeRequest.from_dict(migrate_volume_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


