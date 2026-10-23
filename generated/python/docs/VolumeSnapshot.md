# VolumeSnapshot

A snapshot of a volume (Kubernetes VolumeSnapshot).

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**created_at** | **datetime** | Time the snapshot was taken. | [optional] [readonly] 
**metadata** | **Dict[str, str]** | Metadata set by the consumer and by the VIM; for keys under \&quot;kubevim.kubenfv.io/\&quot; that the VIM maintains, its value wins. | [optional] 
**snapshot_id** | **str** | Identifier of the snapshot: a UUID assigned by the VIM. | [optional] [readonly] 
**snapshot_name** | **str** | Name of the snapshot, given by the consumer; not unique. | [optional] 
**status** | [**SnapshotStatus**](SnapshotStatus.md) | Status of the snapshot. | [optional] [readonly] 
**volume_id** | **str** | Volume the snapshot was taken of. The snapshot is kept when the volume is deleted; the VIM then sets the metadata key \&quot;kubevim.kubenfv.io/source-deleted\&quot; to \&quot;true\&quot;. | [optional] [readonly] 

## Example

```python
from kubevim_client.models.volume_snapshot import VolumeSnapshot

# TODO update the JSON string below
json = "{}"
# create an instance of VolumeSnapshot from a JSON string
volume_snapshot_instance = VolumeSnapshot.from_json(json)
# print the JSON string representation of the object
print(VolumeSnapshot.to_json())

# convert the object into a dict
volume_snapshot_dict = volume_snapshot_instance.to_dict()
# create an instance of VolumeSnapshot from a dict
volume_snapshot_from_dict = VolumeSnapshot.from_dict(volume_snapshot_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


