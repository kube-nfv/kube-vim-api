# Volume

A virtualised storage resource, e.g. a disk of a compute resource.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compute_id** | **str** | Compute resource that uses the volume; empty when not attached. | [optional] [readonly] 
**host_id** | **str** | Host of the volume; empty for distributed storage. | [optional] [readonly] 
**metadata** | **Dict[str, str]** | Metadata set by the consumer and by the VIM; for keys under \&quot;kubevim.kubenfv.io/\&quot; that the VIM maintains, its value wins. | [optional] 
**operational_state** | [**OperationalState**](OperationalState.md) | Operational state of the volume. | [optional] [readonly] 
**storage_data** | [**VirtualStorageData**](VirtualStorageData.md) | Size, type and capabilities of the volume. | 
**volume_id** | **str** | Identifier of the volume: a UUID assigned by the VIM for volumes created with the storage API. The own disks of a compute resource (its boot disk and the disks from its flavour) are not resources of their own and use their disk name instead. | [optional] [readonly] 
**volume_name** | **str** | Name of the volume, given by the consumer; not unique. | [optional] 
**zone_id** | **str** | Resource zone of the volume. | [optional] [readonly] 

## Example

```python
from kubevim_client.models.volume import Volume

# TODO update the JSON string below
json = "{}"
# create an instance of Volume from a JSON string
volume_instance = Volume.from_json(json)
# print the JSON string representation of the object
print(Volume.to_json())

# convert the object into a dict
volume_dict = volume_instance.to_dict()
# create an instance of Volume from a dict
volume_from_dict = Volume.from_dict(volume_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


