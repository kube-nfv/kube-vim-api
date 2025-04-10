# PbCreateComputeResourceAffinityOrAntiAffinityConstraintsGroupRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**group_name** | **str** |  | 
**type** | [**TypeOfAffinityOrAntiAffinityConstraint**](TypeOfAffinityOrAntiAffinityConstraint.md) |  | [default to TypeOfAffinityOrAntiAffinityConstraint.AFFINITY]
**scope** | [**ScopeOfAffinityOrAntiAffinityConstraintForCompute**](ScopeOfAffinityOrAntiAffinityConstraintForCompute.md) |  | [optional] [default to ScopeOfAffinityOrAntiAffinityConstraintForCompute.NFVI_NODE]

## Example

```python
from kubevim_vivnfm_client.models.pb_create_compute_resource_affinity_or_anti_affinity_constraints_group_request import PbCreateComputeResourceAffinityOrAntiAffinityConstraintsGroupRequest

# TODO update the JSON string below
json = "{}"
# create an instance of PbCreateComputeResourceAffinityOrAntiAffinityConstraintsGroupRequest from a JSON string
pb_create_compute_resource_affinity_or_anti_affinity_constraints_group_request_instance = PbCreateComputeResourceAffinityOrAntiAffinityConstraintsGroupRequest.from_json(json)
# print the JSON string representation of the object
print(PbCreateComputeResourceAffinityOrAntiAffinityConstraintsGroupRequest.to_json())

# convert the object into a dict
pb_create_compute_resource_affinity_or_anti_affinity_constraints_group_request_dict = pb_create_compute_resource_affinity_or_anti_affinity_constraints_group_request_instance.to_dict()
# create an instance of PbCreateComputeResourceAffinityOrAntiAffinityConstraintsGroupRequest from a dict
pb_create_compute_resource_affinity_or_anti_affinity_constraints_group_request_from_dict = PbCreateComputeResourceAffinityOrAntiAffinityConstraintsGroupRequest.from_dict(pb_create_compute_resource_affinity_or_anti_affinity_constraints_group_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


