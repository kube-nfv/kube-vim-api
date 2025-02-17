# AffinityOrAntiAffinityConstraintForComputeAffinityOrAntiAffinityResourceList

The AffinityOrAntiAffinityResourceList information element defines an explicit list of resources to express affinity or anti-affinity between these resources and a current resource. The scope of the affinity or anti-affinity can also be defined.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**resource_id** | [**List[Identifier]**](Identifier.md) | List of identifiers of virtualised resources. | [optional] 

## Example

```python
from kubevim_vivnfm_client.models.affinity_or_anti_affinity_constraint_for_compute_affinity_or_anti_affinity_resource_list import AffinityOrAntiAffinityConstraintForComputeAffinityOrAntiAffinityResourceList

# TODO update the JSON string below
json = "{}"
# create an instance of AffinityOrAntiAffinityConstraintForComputeAffinityOrAntiAffinityResourceList from a JSON string
affinity_or_anti_affinity_constraint_for_compute_affinity_or_anti_affinity_resource_list_instance = AffinityOrAntiAffinityConstraintForComputeAffinityOrAntiAffinityResourceList.from_json(json)
# print the JSON string representation of the object
print(AffinityOrAntiAffinityConstraintForComputeAffinityOrAntiAffinityResourceList.to_json())

# convert the object into a dict
affinity_or_anti_affinity_constraint_for_compute_affinity_or_anti_affinity_resource_list_dict = affinity_or_anti_affinity_constraint_for_compute_affinity_or_anti_affinity_resource_list_instance.to_dict()
# create an instance of AffinityOrAntiAffinityConstraintForComputeAffinityOrAntiAffinityResourceList from a dict
affinity_or_anti_affinity_constraint_for_compute_affinity_or_anti_affinity_resource_list_from_dict = AffinityOrAntiAffinityConstraintForComputeAffinityOrAntiAffinityResourceList.from_dict(affinity_or_anti_affinity_constraint_for_compute_affinity_or_anti_affinity_resource_list_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


