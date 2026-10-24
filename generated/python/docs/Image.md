# Image

A software image in the image repository managed by the VIM.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**container_format** | [**ContainerFormat**](ContainerFormat.md) | Container format of the software image file; unspecified when the VIM cannot tell. | [optional] [readonly] 
**created_at** | **datetime** | Time the software image was created. | [optional] [readonly] 
**disk_bus** | [**DiskBus**](DiskBus.md) | Bus the boot disk is attached to; virtio when unset. | [optional] 
**disk_format** | [**DiskFormat**](DiskFormat.md) | Disk format of the software image; unspecified when the VIM cannot tell. | [optional] [readonly] 
**firmware** | [**Firmware**](Firmware.md) | Firmware the image boots with; the VIM default (EFI) when unset. | [optional] 
**image_id** | **str** | Identifier of the software image: a UUID assigned by the VIM. | [optional] [readonly] 
**import_policy** | [**ImportPolicy**](ImportPolicy.md) | When the VIM imports the software image; IMPORT_POLICY_EAGER when unset on creation. | [optional] 
**metadata** | **Dict[str, str]** | Metadata set by the consumer and by the VIM; for keys under \&quot;kubevim.kubenfv.io/\&quot; that the VIM maintains, its value wins. | [optional] 
**min_disk** | **str** | Minimal disk size for the software image, in Kubernetes quantity format (e.g. \&quot;10Gi\&quot;); the VIM rejects a compute resource whose boot disk is smaller. | [optional] 
**min_ram** | **str** | Minimal RAM size for the software image, in Kubernetes quantity format (e.g. \&quot;2Gi\&quot;); the VIM rejects a compute resource with less memory. | [optional] 
**name** | **str** | Name of the software image, given by the consumer; not unique. | 
**progress** | **int** | Import progress in percent while the status is IMAGE_STATUS_IMPORTING. | [optional] [readonly] 
**provider** | **str** | Provider of the software image. | [optional] 
**size** | **str** | Size of the software image, in Kubernetes quantity format; empty until imported. | [optional] [readonly] 
**source** | [**ImageSource**](ImageSource.md) | Where the VIM imports the software image from; images of a catalog backend name their catalog entry. | 
**status** | [**ImageStatus**](ImageStatus.md) | Status of the software image. | [optional] [readonly] 
**status_detail** | **str** | Human-readable reason when the status is IMAGE_STATUS_FAILED. | [optional] [readonly] 
**updated_at** | **datetime** | Time the software image was last updated. | [optional] [readonly] 
**version** | **str** | Version of the software image. | [optional] 

## Example

```python
from kubevim_client.models.image import Image

# TODO update the JSON string below
json = "{}"
# create an instance of Image from a JSON string
image_instance = Image.from_json(json)
# print the JSON string representation of the object
print(Image.to_json())

# convert the object into a dict
image_dict = image_instance.to_dict()
# create an instance of Image from a dict
image_from_dict = Image.from_dict(image_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


