# ComputeSnapshot

A snapshot of a compute resource (KubeVirt VirtualMachineSnapshot).

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compute_id** | **str** | Compute resource the snapshot was taken of. The snapshot is kept when the compute resource is deleted; the VIM then sets the metadata key \&quot;kubevim.kubenfv.io/source-deleted\&quot; to \&quot;true\&quot;. | [optional] [readonly] 
**created_at** | **datetime** | Time the snapshot was taken. | [optional] [readonly] 
**metadata** | **Dict[str, str]** | Metadata set by the consumer and by the VIM; for keys under \&quot;kubevim.kubenfv.io/\&quot; that the VIM maintains, its value wins. | [optional] 
**snapshot_id** | **str** | Identifier of the snapshot: a UUID assigned by the VIM. | [optional] [readonly] 
**snapshot_name** | **str** | Name of the snapshot, given by the consumer; not unique. | [optional] 
**status** | [**SnapshotStatus**](SnapshotStatus.md) | Status of the snapshot. | [optional] [readonly] 

## Example

```python
from kubevim_client.models.compute_snapshot import ComputeSnapshot

# TODO update the JSON string below
json = "{}"
# create an instance of ComputeSnapshot from a JSON string
compute_snapshot_instance = ComputeSnapshot.from_json(json)
# print the JSON string representation of the object
print(ComputeSnapshot.to_json())

# convert the object into a dict
compute_snapshot_dict = compute_snapshot_instance.to_dict()
# create an instance of ComputeSnapshot from a dict
compute_snapshot_from_dict = ComputeSnapshot.from_dict(compute_snapshot_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


