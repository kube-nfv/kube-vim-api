# SoftwareImageInformation


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**software_image_id** | [**Identifier**](Identifier.md) |  | 
**name** | **str** | Name of this software image. | 
**provider** | **str** | Provider of this software image. | [optional] 
**version** | **str** | Version of the software image file. | [optional] 
**checksum** | **str** | Checksum of the software image file. | [optional] 
**container_format** | **str** | Container format indicates whether the software image is in a file format that also contains metadata about the actual software. | [optional] 
**disk_format** | **str** | Disk format of a software image is the format of the underlying disk image. | [optional] 
**created_at** | **object** | Time this software image was created. | 
**updated_at** | **object** | Time this software image was last updated. | 
**min_disk** | **str** | Minimal disk size for this software image. | [optional] 
**min_ram** | **str** | Minimal RAM size for this software image. | [optional] 
**size** | **str** | Size of this software image. | 
**status** | **str** |  | 
**metadata** | [**Metadata**](Metadata.md) |  | [optional] 

## Example

```python
from kubevim_vivnfm_client.models.software_image_information import SoftwareImageInformation

# TODO update the JSON string below
json = "{}"
# create an instance of SoftwareImageInformation from a JSON string
software_image_information_instance = SoftwareImageInformation.from_json(json)
# print the JSON string representation of the object
print(SoftwareImageInformation.to_json())

# convert the object into a dict
software_image_information_dict = software_image_information_instance.to_dict()
# create an instance of SoftwareImageInformation from a dict
software_image_information_from_dict = SoftwareImageInformation.from_dict(software_image_information_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


