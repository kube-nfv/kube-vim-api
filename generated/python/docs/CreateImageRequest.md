# CreateImageRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**image** | [**Image**](Image.md) | Software image to add. | 
**request_id** | **str** | Client-chosen UUID that makes the request idempotent (AIP-155). | [optional] 

## Example

```python
from kubevim_client.models.create_image_request import CreateImageRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CreateImageRequest from a JSON string
create_image_request_instance = CreateImageRequest.from_json(json)
# print the JSON string representation of the object
print(CreateImageRequest.to_json())

# convert the object into a dict
create_image_request_dict = create_image_request_instance.to_dict()
# create an instance of CreateImageRequest from a dict
create_image_request_from_dict = CreateImageRequest.from_dict(create_image_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


