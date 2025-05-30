# PbAllocateComputeRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compute_name** | **str** | Name provided by the consumer for the virtualised compute resource to be allocated. It can be used for identifying resources from consumer side. | [optional] 
**reservation_id** | [**Identifier**](Identifier.md) |  | [optional] 
**affinity_or_anti_affinity_constraints** | [**List[AffinityOrAntiAffinityConstraintForCompute]**](AffinityOrAntiAffinityConstraintForCompute.md) | List of elements with affinity or anti affinity (see clause 8.4.8.2) information of the virtualised compute resource to be allocated. All the listed constraints shall be fulfilled for a successful operation. | [optional] 
**compute_flavour_id** | [**Identifier**](Identifier.md) |  | 
**vc_image_id** | [**Identifier**](Identifier.md) |  | [optional] 
**interface_data** | [**List[VirtualNetworkInterfaceData]**](VirtualNetworkInterfaceData.md) | Note: That is out of the ETSI GS NFV-IFA 006 scope. Traditionaly VirtualNetworkInterfaceData specified in the virtualComputeFlavour, but it is reduce flexibility, since the flavor contains virtual compute related networks, and network configuration for it (eg. QoS). Descided to move it in the AllocateComputeRequest. | [optional] 
**interface_ipam** | [**List[VirtualNetworkInterfaceIPAM]**](VirtualNetworkInterfaceIPAM.md) | IPAM Data of network interfaces which are specific to a Virtual Compute Resource instance. See clause 8.4.3.7. | [optional] 
**meta_data** | [**Metadata**](Metadata.md) |  | [optional] 
**resource_group_id** | [**Identifier**](Identifier.md) |  | [optional] 
**user_data** | [**UserData**](UserData.md) |  | [optional] 

## Example

```python
from kubevim_vivnfm_client.models.pb_allocate_compute_request import PbAllocateComputeRequest

# TODO update the JSON string below
json = "{}"
# create an instance of PbAllocateComputeRequest from a JSON string
pb_allocate_compute_request_instance = PbAllocateComputeRequest.from_json(json)
# print the JSON string representation of the object
print(PbAllocateComputeRequest.to_json())

# convert the object into a dict
pb_allocate_compute_request_dict = pb_allocate_compute_request_instance.to_dict()
# create an instance of PbAllocateComputeRequest from a dict
pb_allocate_compute_request_from_dict = PbAllocateComputeRequest.from_dict(pb_allocate_compute_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


