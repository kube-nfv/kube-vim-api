# CreateVolumeSnapshotRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**request_id** | **str** | Client-chosen UUID that makes the request idempotent (AIP-155). | [optional] 
**snapshot** | [**VolumeSnapshot**](VolumeSnapshot.md) | Snapshot to create: name and metadata. | [optional] 

## Example

```python
from kubevim_client.models.create_volume_snapshot_request import CreateVolumeSnapshotRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CreateVolumeSnapshotRequest from a JSON string
create_volume_snapshot_request_instance = CreateVolumeSnapshotRequest.from_json(json)
# print the JSON string representation of the object
print(CreateVolumeSnapshotRequest.to_json())

# convert the object into a dict
create_volume_snapshot_request_dict = create_volume_snapshot_request_instance.to_dict()
# create an instance of CreateVolumeSnapshotRequest from a dict
create_volume_snapshot_request_from_dict = CreateVolumeSnapshotRequest.from_dict(create_volume_snapshot_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


