# ProblemDetails

Error details (SOL 013 6.3, RFC 9457).

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**detail** | **str** |  | 
**instance** | **str** |  | [optional] 
**status** | **int** |  | 
**title** | **str** |  | [optional] 
**type** | **str** | https://kubevim.kubenfv.io/problems/&lt;reason&gt;, the error reason (ErrorReason in the protos) in kebab case, e.g. .../problems/resource-in-use; about:blank for authentication and authorisation errors. Clients should accept unknown types. | [optional] 

## Example

```python
from kubevim_client.models.problem_details import ProblemDetails

# TODO update the JSON string below
json = "{}"
# create an instance of ProblemDetails from a JSON string
problem_details_instance = ProblemDetails.from_json(json)
# print the JSON string representation of the object
print(ProblemDetails.to_json())

# convert the object into a dict
problem_details_dict = problem_details_instance.to_dict()
# create an instance of ProblemDetails from a dict
problem_details_from_dict = ProblemDetails.from_dict(problem_details_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


