# ViVnfmOperateVirtualisedComputeResourceBody


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compute_id** | **object** | Identifier of the virtualised compute resource successfully terminated. | 
**compute_operation_input_data** | **Dict[str, str]** | Additional parameters associated to the operation to be performed. For example, if the operation is \&quot;delete snapshot\&quot;, information identifying the snapshot to be deleted is provided. | [optional] 

## Example

```python
from kubevim_vivnfm_client.models.vi_vnfm_operate_virtualised_compute_resource_body import ViVnfmOperateVirtualisedComputeResourceBody

# TODO update the JSON string below
json = "{}"
# create an instance of ViVnfmOperateVirtualisedComputeResourceBody from a JSON string
vi_vnfm_operate_virtualised_compute_resource_body_instance = ViVnfmOperateVirtualisedComputeResourceBody.from_json(json)
# print the JSON string representation of the object
print(ViVnfmOperateVirtualisedComputeResourceBody.to_json())

# convert the object into a dict
vi_vnfm_operate_virtualised_compute_resource_body_dict = vi_vnfm_operate_virtualised_compute_resource_body_instance.to_dict()
# create an instance of ViVnfmOperateVirtualisedComputeResourceBody from a dict
vi_vnfm_operate_virtualised_compute_resource_body_from_dict = ViVnfmOperateVirtualisedComputeResourceBody.from_dict(vi_vnfm_operate_virtualised_compute_resource_body_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


