# PbQueryComputeResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**query_result** | [**List[VirtualCompute]**](VirtualCompute.md) | Contains information about the virtual compute resource(s) matching the filter. The cardinality can be 0 if no matching compute resources exist. | 

## Example

```python
from kubevim_vivnfm_client.models.pb_query_compute_response import PbQueryComputeResponse

# TODO update the JSON string below
json = "{}"
# create an instance of PbQueryComputeResponse from a JSON string
pb_query_compute_response_instance = PbQueryComputeResponse.from_json(json)
# print the JSON string representation of the object
print(PbQueryComputeResponse.to_json())

# convert the object into a dict
pb_query_compute_response_dict = pb_query_compute_response_instance.to_dict()
# create an instance of PbQueryComputeResponse from a dict
pb_query_compute_response_from_dict = PbQueryComputeResponse.from_dict(pb_query_compute_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


