# kubevim_client.ImageServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_image**](ImageServiceApi.md#create_image) | **POST** /image/v1/images | Adds a software image to the image repository managed by the VIM.
[**delete_image**](ImageServiceApi.md#delete_image) | **DELETE** /image/v1/images/{image_id} | Deletes a software image and its imported copies; compute resources created from it keep their disks.
[**get_image**](ImageServiceApi.md#get_image) | **GET** /image/v1/images/{image_id} | Returns one software image.
[**list_image_sources**](ImageServiceApi.md#list_image_sources) | **GET** /image/v1/imageSources | Returns the image source types this VIM accepts.
[**list_images**](ImageServiceApi.md#list_images) | **GET** /image/v1/images | Returns the software images matching a filter, including images of catalog backends.
[**update_image**](ImageServiceApi.md#update_image) | **PATCH** /image/v1/images/{image_id} | Updates the name and metadata of a software image.


# **create_image**
> Image create_image(create_image_request)

Adds a software image to the image repository managed by the VIM.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.create_image_request import CreateImageRequest
from kubevim_client.models.image import Image
from kubevim_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = kubevim_client.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

configuration.access_token = os.environ["ACCESS_TOKEN"]

# Enter a context with an instance of the API client
with kubevim_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = kubevim_client.ImageServiceApi(api_client)
    create_image_request = kubevim_client.CreateImageRequest() # CreateImageRequest | 

    try:
        # Adds a software image to the image repository managed by the VIM.
        api_response = api_instance.create_image(create_image_request)
        print("The response of ImageServiceApi->create_image:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ImageServiceApi->create_image: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **create_image_request** | [**CreateImageRequest**](CreateImageRequest.md)|  | 

### Return type

[**Image**](Image.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | A software image in the image repository managed by the VIM. |  * Location - URI of the created resource. <br>  |
**0** | Error (SOL 013 6). |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_image**
> delete_image(image_id)

Deletes a software image and its imported copies; compute resources created from it keep their disks.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = kubevim_client.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

configuration.access_token = os.environ["ACCESS_TOKEN"]

# Enter a context with an instance of the API client
with kubevim_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = kubevim_client.ImageServiceApi(api_client)
    image_id = 'image_id_example' # str | Identifier of the software image.

    try:
        # Deletes a software image and its imported copies; compute resources created from it keep their disks.
        api_instance.delete_image(image_id)
    except Exception as e:
        print("Exception when calling ImageServiceApi->delete_image: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **image_id** | **str**| Identifier of the software image. | 

### Return type

void (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**204** | No Content |  -  |
**0** | Error (SOL 013 6). |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_image**
> Image get_image(image_id)

Returns one software image.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.image import Image
from kubevim_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = kubevim_client.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

configuration.access_token = os.environ["ACCESS_TOKEN"]

# Enter a context with an instance of the API client
with kubevim_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = kubevim_client.ImageServiceApi(api_client)
    image_id = 'image_id_example' # str | Identifier of the software image.

    try:
        # Returns one software image.
        api_response = api_instance.get_image(image_id)
        print("The response of ImageServiceApi->get_image:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ImageServiceApi->get_image: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **image_id** | **str**| Identifier of the software image. | 

### Return type

[**Image**](Image.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A software image in the image repository managed by the VIM. |  -  |
**0** | Error (SOL 013 6). |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_image_sources**
> List[ImageSourceType] list_image_sources()

Returns the image source types this VIM accepts.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.image_source_type import ImageSourceType
from kubevim_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = kubevim_client.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

configuration.access_token = os.environ["ACCESS_TOKEN"]

# Enter a context with an instance of the API client
with kubevim_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = kubevim_client.ImageServiceApi(api_client)

    try:
        # Returns the image source types this VIM accepts.
        api_response = api_instance.list_image_sources()
        print("The response of ImageServiceApi->list_image_sources:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ImageServiceApi->list_image_sources: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**List[ImageSourceType]**](ImageSourceType.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | Error (SOL 013 6). |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_images**
> List[Image] list_images(filter=filter, nextpage_opaque_marker=nextpage_opaque_marker)

Returns the software images matching a filter, including images of catalog backends.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.image import Image
from kubevim_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = kubevim_client.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

configuration.access_token = os.environ["ACCESS_TOKEN"]

# Enter a context with an instance of the API client
with kubevim_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = kubevim_client.ImageServiceApi(api_client)
    filter = 'filter_example' # str | Attribute-based filter expression (SOL 013 5.2). (optional)
    nextpage_opaque_marker = 'nextpage_opaque_marker_example' # str | Marker of the page to return, from a previous response (SOL 013 5.4). (optional)

    try:
        # Returns the software images matching a filter, including images of catalog backends.
        api_response = api_instance.list_images(filter=filter, nextpage_opaque_marker=nextpage_opaque_marker)
        print("The response of ImageServiceApi->list_images:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ImageServiceApi->list_images: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **str**| Attribute-based filter expression (SOL 013 5.2). | [optional] 
 **nextpage_opaque_marker** | **str**| Marker of the page to return, from a previous response (SOL 013 5.4). | [optional] 

### Return type

[**List[Image]**](Image.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  * Link - URI of the next page, rel&#x3D;\&quot;next\&quot;, carrying nextpage_opaque_marker; absent on the last page (SOL 013 5.4.2). <br>  |
**0** | Error (SOL 013 6). |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_image**
> Image update_image(image_id, image_update)

Updates the name and metadata of a software image.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.image import Image
from kubevim_client.models.image_update import ImageUpdate
from kubevim_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = kubevim_client.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

configuration.access_token = os.environ["ACCESS_TOKEN"]

# Enter a context with an instance of the API client
with kubevim_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = kubevim_client.ImageServiceApi(api_client)
    image_id = 'image_id_example' # str | Identifier of the software image.
    image_update = kubevim_client.ImageUpdate() # ImageUpdate | 

    try:
        # Updates the name and metadata of a software image.
        api_response = api_instance.update_image(image_id, image_update)
        print("The response of ImageServiceApi->update_image:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ImageServiceApi->update_image: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **image_id** | **str**| Identifier of the software image. | 
 **image_update** | [**ImageUpdate**](ImageUpdate.md)|  | 

### Return type

[**Image**](Image.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A software image in the image repository managed by the VIM. |  -  |
**0** | Error (SOL 013 6). |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

