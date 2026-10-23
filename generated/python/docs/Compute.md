# Compute

A virtualised compute resource: a virtual machine.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compute_id** | **str** | Identifier of the compute resource: a UUID assigned by the VIM. | [optional] [readonly] 
**compute_name** | **str** | Name of the compute resource, given by the consumer; not unique. | [optional] 
**disks** | [**List[Volume]**](Volume.md) | Disks of the compute resource: its own disks (the boot disk and the disks from the flavour, identified by disk name, deleted with the compute resource and not in the storage API) and attached volumes (identified by their volume_id). | [optional] [readonly] 
**flavour_id** | **str** | Flavour the compute resource is sized by; it needs exactly one boot storage attribute. | 
**host_id** | **str** | Host the compute resource runs on (the Kubernetes node). | [optional] [readonly] 
**image_id** | **str** | Software image the compute resource boots from. | 
**metadata** | **Dict[str, str]** | Metadata set by the consumer and by the VIM; for keys under \&quot;kubevim.kubenfv.io/\&quot; that the VIM maintains, its value wins. | [optional] 
**network_interfaces** | [**List[VirtualNetworkInterface]**](VirtualNetworkInterface.md) | Network interfaces, the management interface first; no other order is guaranteed, in the response or in the guest. Match interfaces by name, MAC address or guest PCI address. | [optional] [readonly] 
**operational_state** | [**OperationalState**](OperationalState.md) | Operational state of the compute resource. | [optional] [readonly] 
**placement_group_ids** | **List[str]** | Placement groups the compute resource belongs to, joined at creation. | [optional] 
**restart_required** | **bool** | A change (a flavour beyond the hotplug limits, or interfaces of a VM that cannot live-migrate) applies on the next restart. | [optional] [readonly] 
**running_state** | [**ComputeRunningState**](ComputeRunningState.md) | Running state of the compute resource. | [optional] [readonly] 
**status_detail** | **str** | Human-readable reason when the running state is FAILED. | [optional] [readonly] 
**virtual_cpu** | [**VirtualCpu**](VirtualCpu.md) | Virtual CPUs of the compute resource. | [optional] [readonly] 
**virtual_memory** | [**VirtualMemory**](VirtualMemory.md) | Virtual memory of the compute resource. | [optional] [readonly] 
**zone_id** | **str** | Resource zone of the compute resource. | [optional] [readonly] 

## Example

```python
from kubevim_client.models.compute import Compute

# TODO update the JSON string below
json = "{}"
# create an instance of Compute from a JSON string
compute_instance = Compute.from_json(json)
# print the JSON string representation of the object
print(Compute.to_json())

# convert the object into a dict
compute_dict = compute_instance.to_dict()
# create an instance of Compute from a dict
compute_from_dict = Compute.from_dict(compute_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


