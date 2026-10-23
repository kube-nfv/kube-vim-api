# CpuTopology

Guest CPU topology.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cores** | **int** |  | 
**sockets** | **int** |  | 
**threads** | **int** |  | 

## Example

```python
from kubevim_client.models.cpu_topology import CpuTopology

# TODO update the JSON string below
json = "{}"
# create an instance of CpuTopology from a JSON string
cpu_topology_instance = CpuTopology.from_json(json)
# print the JSON string representation of the object
print(CpuTopology.to_json())

# convert the object into a dict
cpu_topology_dict = cpu_topology_instance.to_dict()
# create an instance of CpuTopology from a dict
cpu_topology_from_dict = CpuTopology.from_dict(cpu_topology_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


