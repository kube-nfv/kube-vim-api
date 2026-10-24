# ImageSourceHttp


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**http** | [**HttpSource**](HttpSource.md) | Download over HTTP or HTTPS. | 

## Example

```python
from kubevim_client.models.image_source_http import ImageSourceHttp

# TODO update the JSON string below
json = "{}"
# create an instance of ImageSourceHttp from a JSON string
image_source_http_instance = ImageSourceHttp.from_json(json)
# print the JSON string representation of the object
print(ImageSourceHttp.to_json())

# convert the object into a dict
image_source_http_dict = image_source_http_instance.to_dict()
# create an instance of ImageSourceHttp from a dict
image_source_http_from_dict = ImageSourceHttp.from_dict(image_source_http_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


