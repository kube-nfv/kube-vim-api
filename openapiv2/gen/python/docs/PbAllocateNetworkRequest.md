# PbAllocateNetworkRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**network_resource_name** | **str** | Name provided by the consumer for the virtualised network resource to be allocated. It can be used for identifying resources from consumer side. | [optional] 
**reservation_id** | [**Identifier**](Identifier.md) |  | [optional] 
**network_resource_type** | [**NetworkResourceType**](NetworkResourceType.md) |  | [optional] [default to NetworkResourceType.NETWORK]
**type_network_data** | [**VirtualNetworkData**](VirtualNetworkData.md) |  | [optional] 
**type_subnet_data** | [**NetworkSubnetData**](NetworkSubnetData.md) |  | [optional] 
**type_network_port_data** | **object** |  | [optional] 
**type_trunk_data** | **object** |  | [optional] 
**affinity_or_anti_affinity_constraints** | **List[object]** | List of elements with affinity or anti affinity (see clause 8.4.8.2) information of the virtualised network resource to be allocated. All the listed constraints shall be fulfilled for a successful operation. | [optional] 
**location_constraints_for_network** | **str** | If present, it defines location constraints for the resource(s) to be allocated, e.g. in what particular resource zone. | [optional] 
**meta_data** | [**Dict[str, ProtobufAny]**](ProtobufAny.md) | List of metadata key-value pairs used by the consumer to associate meaningful metadata to the related virtualised resource. | [optional] 
**resource_group_id** | [**Identifier**](Identifier.md) |  | [optional] 

## Example

```python
from vivnfm_client.models.pb_allocate_network_request import PbAllocateNetworkRequest

# TODO update the JSON string below
json = "{}"
# create an instance of PbAllocateNetworkRequest from a JSON string
pb_allocate_network_request_instance = PbAllocateNetworkRequest.from_json(json)
# print the JSON string representation of the object
print(PbAllocateNetworkRequest.to_json())

# convert the object into a dict
pb_allocate_network_request_dict = pb_allocate_network_request_instance.to_dict()
# create an instance of PbAllocateNetworkRequest from a dict
pb_allocate_network_request_from_dict = PbAllocateNetworkRequest.from_dict(pb_allocate_network_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


