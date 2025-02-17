# AffinityOrAntiAffinityConstraintForCompute


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | [**TypeOfAffinityOrAntiAffinityConstraint**](TypeOfAffinityOrAntiAffinityConstraint.md) |  | [optional] [default to TypeOfAffinityOrAntiAffinityConstraint.AFFINITY]
**scope** | [**AffinityOrAntiAffinityConstraintForComputeScopeOfAffinityOrAntiAffinityConstraintForCompute**](AffinityOrAntiAffinityConstraintForComputeScopeOfAffinityOrAntiAffinityConstraintForCompute.md) |  | [optional] [default to AffinityOrAntiAffinityConstraintForComputeScopeOfAffinityOrAntiAffinityConstraintForCompute.NFVI_NODE]
**affinity_or_anti_affinity_resource_list** | [**AffinityOrAntiAffinityConstraintForComputeAffinityOrAntiAffinityResourceList**](AffinityOrAntiAffinityConstraintForComputeAffinityOrAntiAffinityResourceList.md) |  | [optional] 
**affinity_or_anti_affinity_resource_group_id** | [**Identifier**](Identifier.md) |  | [optional] 

## Example

```python
from vivnfm_client.models.affinity_or_anti_affinity_constraint_for_compute import AffinityOrAntiAffinityConstraintForCompute

# TODO update the JSON string below
json = "{}"
# create an instance of AffinityOrAntiAffinityConstraintForCompute from a JSON string
affinity_or_anti_affinity_constraint_for_compute_instance = AffinityOrAntiAffinityConstraintForCompute.from_json(json)
# print the JSON string representation of the object
print(AffinityOrAntiAffinityConstraintForCompute.to_json())

# convert the object into a dict
affinity_or_anti_affinity_constraint_for_compute_dict = affinity_or_anti_affinity_constraint_for_compute_instance.to_dict()
# create an instance of AffinityOrAntiAffinityConstraintForCompute from a dict
affinity_or_anti_affinity_constraint_for_compute_from_dict = AffinityOrAntiAffinityConstraintForCompute.from_dict(affinity_or_anti_affinity_constraint_for_compute_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


