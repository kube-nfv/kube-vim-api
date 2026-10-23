# CreateComputeSnapshotRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**request_id** | **str** | Client-chosen UUID that makes the request idempotent (AIP-155). | [optional] 
**snapshot** | [**ComputeSnapshot**](ComputeSnapshot.md) | Snapshot to create: name and metadata. | [optional] 

## Example

```python
from kubevim_client.models.create_compute_snapshot_request import CreateComputeSnapshotRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CreateComputeSnapshotRequest from a JSON string
create_compute_snapshot_request_instance = CreateComputeSnapshotRequest.from_json(json)
# print the JSON string representation of the object
print(CreateComputeSnapshotRequest.to_json())

# convert the object into a dict
create_compute_snapshot_request_dict = create_compute_snapshot_request_instance.to_dict()
# create an instance of CreateComputeSnapshotRequest from a dict
create_compute_snapshot_request_from_dict = CreateComputeSnapshotRequest.from_dict(create_compute_snapshot_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


