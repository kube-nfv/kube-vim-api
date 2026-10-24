# VirtualStorageData

Size, type and capabilities of a virtualised storage resource.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**is_boot** | **bool** | The software image is copied to this storage and the compute resource boots from it. | [optional] 
**size_of_storage** | **str** | Size of the storage resource, in Kubernetes quantity format (e.g. \&quot;20Gi\&quot;). | 
**storage_class** | **str** | Kubernetes storage class of the storage, one labelled kubevim.kubenfv.io/storage&#x3D;true; the VIM default when unset. | [optional] 
**type_of_storage** | [**TypeOfStorage**](TypeOfStorage.md) | Type of the storage resource. | 

## Example

```python
from kubevim_client.models.virtual_storage_data import VirtualStorageData

# TODO update the JSON string below
json = "{}"
# create an instance of VirtualStorageData from a JSON string
virtual_storage_data_instance = VirtualStorageData.from_json(json)
# print the JSON string representation of the object
print(VirtualStorageData.to_json())

# convert the object into a dict
virtual_storage_data_dict = virtual_storage_data_instance.to_dict()
# create an instance of VirtualStorageData from a dict
virtual_storage_data_from_dict = VirtualStorageData.from_dict(virtual_storage_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


