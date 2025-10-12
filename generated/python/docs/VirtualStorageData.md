# VirtualStorageData

This clause describes the attributes for the VirtualStorageData information element.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type_of_storage** | **str** | Type of virtualised storage resource (e.g. volume, object). | 
**size_of_storage** | **float** | Size of virtualised storage resource (e.g. size of volume, in GB). | 
**rdma_enabled** | **bool** | Indicates if the storage supports RDMA. | [optional] 
**is_boot** | **bool** |  | [optional] [default to False]

## Example

```python
from kubevim_vivnfm_client.models.virtual_storage_data import VirtualStorageData

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


