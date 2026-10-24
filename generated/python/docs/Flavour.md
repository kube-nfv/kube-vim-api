# Flavour

Memory, CPU and storage of the compute resources created from it.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**flavour_id** | **str** | Identifier of the flavour: a UUID assigned by the VIM. | [optional] [readonly] 
**flavour_name** | **str** | Name of the flavour, given by the consumer; not unique. | [optional] 
**metadata** | **Dict[str, str]** | Metadata set by the consumer and by the VIM; for keys under \&quot;kubevim.kubenfv.io/\&quot; that the VIM maintains, its value wins. | [optional] 
**storage_attributes** | [**List[VirtualStorageData]**](VirtualStorageData.md) | Storage of the compute resource: the boot disk (is_boot) and further disks, created with each compute resource. A compute resource needs exactly one boot disk. | [optional] 
**virtual_cpu** | [**VirtualCpuData**](VirtualCpuData.md) | Virtual CPUs. | 
**virtual_memory** | [**VirtualMemoryData**](VirtualMemoryData.md) | Virtual memory. | 

## Example

```python
from kubevim_client.models.flavour import Flavour

# TODO update the JSON string below
json = "{}"
# create an instance of Flavour from a JSON string
flavour_instance = Flavour.from_json(json)
# print the JSON string representation of the object
print(Flavour.to_json())

# convert the object into a dict
flavour_dict = flavour_instance.to_dict()
# create an instance of Flavour from a dict
flavour_from_dict = Flavour.from_dict(flavour_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


