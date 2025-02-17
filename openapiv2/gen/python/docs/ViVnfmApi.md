# kubevim_vivnfm_client.ViVnfmApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**vi_vnfm_allocate_virtualised_compute_resource**](ViVnfmApi.md#vi_vnfm_allocate_virtualised_compute_resource) | **POST** /vivnfm/v5/compute | This operation allows requesting the allocation of virtualised compute resources as indicated by the consumer functional block. Result: After successful operation, the VIM has created the internal management objects for the virtualised compute resource and allocated this resource according to the input requirements and constraints. In addition, the VIM shall return to the VNFM information on the newly instantiated virtualised compute resource plus any additional information about the allocate request operation. The VIM may also return intermediate status reports during the allocation process. If the operation was not successful, the VIM shall return to the VNFM appropriate error information.
[**vi_vnfm_allocate_virtualised_network_resource**](ViVnfmApi.md#vi_vnfm_allocate_virtualised_network_resource) | **POST** /vivnfm/v5/networks | This operation allows requesting the allocation of virtualised network resources as indicated by the consumer functional block. Result: After successful operation, the VIM has created the internal management objects for the virtualised network resource and allocated this resource. In addition, the VIM shall return to the VNFM information on the newly instantiated virtualised network resource plus any additional information about the allocate request operation. The VIM may also return intermediate status reports during the allocation process. If the operation was not successful, the VIM shall return to the VNFM appropriate error information.
[**vi_vnfm_create_compute_flavour**](ViVnfmApi.md#vi_vnfm_create_compute_flavour) | **POST** /vivnfm/v5/flavours | This operation allows requesting the creation of a flavour as indicated by the consumer functional block. Result: After successful operation, the VIM has created the Compute Flavour. In addition, the VIM shall return to the VNFM information on the newly created Compute Flavour. If the operation was not successful, the VIM shall return to the VNFM appropriate error information.
[**vi_vnfm_delete_compute_flavour**](ViVnfmApi.md#vi_vnfm_delete_compute_flavour) | **DELETE** /vivnfm/v5/flavours/{computeFlavourId.value} | This operation allows deleting a Compute Flavour. Result: After successful operation, the VIM has deleted the Compute Flavour, so no new Virtualised Compute Resource can be allocated based on it. The already allocated Virtualised Compute Resources are not affected. If the operation was not successful, the VIM shall return to the VNFM appropriate error information.
[**vi_vnfm_query_compute_flavour**](ViVnfmApi.md#vi_vnfm_query_compute_flavour) | **GET** /vivnfm/v5/flavours | This operation allows querying information about created Compute Flavours. Result: After successful operation, the VIM has queried the internal management objects for the Compute Flavours. The result of the query shall indicate with a standard success/error result if the query has been processed correctly. For a particular query, information about the Compute Flavours that the VNFM has access to and that are matching the filter shall be returned.
[**vi_vnfm_query_image**](ViVnfmApi.md#vi_vnfm_query_image) | **GET** /vivnfm/v5/images/{softwareImageId.value} | This operation allows querying the information about a specific software image in the image repository managed by the VIM. Result: As a result of this operation, the producer (VIM) shall indicate to the consumer (VNFM) whether or not it was possible to process the query.
[**vi_vnfm_query_image2**](ViVnfmApi.md#vi_vnfm_query_image2) | **POST** /vivnfm/v5/images | This operation allows querying the information about a specific software image in the image repository managed by the VIM. Result: As a result of this operation, the producer (VIM) shall indicate to the consumer (VNFM) whether or not it was possible to process the query.
[**vi_vnfm_query_images**](ViVnfmApi.md#vi_vnfm_query_images) | **GET** /vivnfm/v5/images | This operation allows querying the information of software images in the image repository managed by the VIM. Result: As a result of this operation, the producer (VIM) shall indicate to the consumer (VNFM) whether or not it was possible to process the query
[**vi_vnfm_query_virtualised_network_resource**](ViVnfmApi.md#vi_vnfm_query_virtualised_network_resource) | **GET** /vivnfm/v5/networks | This operation allows querying information about instantiated virtualised network resources. Result: After successful operation, the VIM has queried the internal management objects for the virtualised network resources. The result of the query shall indicate with a standard success/error result if the query has been processed correctly. For a particular query, information about the network resources that the VNFM has access to and that are matching the filter shall be returned.
[**vi_vnfm_terminate_virtualised_network_resource**](ViVnfmApi.md#vi_vnfm_terminate_virtualised_network_resource) | **DELETE** /vivnfm/v5/networks/{networkResourceId.value} | This operation allows de-allocating and terminating one or more an instantiated virtualised network resource(s). When the operation is done on multiple ids, it is assumed to be best-effort, i.e. it can succeed for a subset of the ids, and fail for the remaining ones. Result: After successful operation, the VIM has terminated the virtualised network resources and removed the internal management objects for those resources. In addition, the VIM shall return to the VNFM information on the terminated virtualised network resource plus any additional information about the terminate request operation. If the operation was not successful, the VIM shall return to the VNFM appropriate error information.


# **vi_vnfm_allocate_virtualised_compute_resource**
> PbAllocateComputeResponse vi_vnfm_allocate_virtualised_compute_resource(body)

This operation allows requesting the allocation of virtualised compute resources as indicated by the consumer functional block. Result: After successful operation, the VIM has created the internal management objects for the virtualised compute resource and allocated this resource according to the input requirements and constraints. In addition, the VIM shall return to the VNFM information on the newly instantiated virtualised compute resource plus any additional information about the allocate request operation. The VIM may also return intermediate status reports during the allocation process. If the operation was not successful, the VIM shall return to the VNFM appropriate error information.

### Example


```python
import kubevim_vivnfm_client
from kubevim_vivnfm_client.models.pb_allocate_compute_request import PbAllocateComputeRequest
from kubevim_vivnfm_client.models.pb_allocate_compute_response import PbAllocateComputeResponse
from kubevim_vivnfm_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = kubevim_vivnfm_client.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with kubevim_vivnfm_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = kubevim_vivnfm_client.ViVnfmApi(api_client)
    body = kubevim_vivnfm_client.PbAllocateComputeRequest() # PbAllocateComputeRequest | 

    try:
        # This operation allows requesting the allocation of virtualised compute resources as indicated by the consumer functional block. Result: After successful operation, the VIM has created the internal management objects for the virtualised compute resource and allocated this resource according to the input requirements and constraints. In addition, the VIM shall return to the VNFM information on the newly instantiated virtualised compute resource plus any additional information about the allocate request operation. The VIM may also return intermediate status reports during the allocation process. If the operation was not successful, the VIM shall return to the VNFM appropriate error information.
        api_response = api_instance.vi_vnfm_allocate_virtualised_compute_resource(body)
        print("The response of ViVnfmApi->vi_vnfm_allocate_virtualised_compute_resource:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ViVnfmApi->vi_vnfm_allocate_virtualised_compute_resource: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PbAllocateComputeRequest**](PbAllocateComputeRequest.md)|  | 

### Return type

[**PbAllocateComputeResponse**](PbAllocateComputeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **vi_vnfm_allocate_virtualised_network_resource**
> PbAllocateNetworkResponse vi_vnfm_allocate_virtualised_network_resource(body)

This operation allows requesting the allocation of virtualised network resources as indicated by the consumer functional block. Result: After successful operation, the VIM has created the internal management objects for the virtualised network resource and allocated this resource. In addition, the VIM shall return to the VNFM information on the newly instantiated virtualised network resource plus any additional information about the allocate request operation. The VIM may also return intermediate status reports during the allocation process. If the operation was not successful, the VIM shall return to the VNFM appropriate error information.

### Example


```python
import kubevim_vivnfm_client
from kubevim_vivnfm_client.models.pb_allocate_network_request import PbAllocateNetworkRequest
from kubevim_vivnfm_client.models.pb_allocate_network_response import PbAllocateNetworkResponse
from kubevim_vivnfm_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = kubevim_vivnfm_client.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with kubevim_vivnfm_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = kubevim_vivnfm_client.ViVnfmApi(api_client)
    body = kubevim_vivnfm_client.PbAllocateNetworkRequest() # PbAllocateNetworkRequest | 

    try:
        # This operation allows requesting the allocation of virtualised network resources as indicated by the consumer functional block. Result: After successful operation, the VIM has created the internal management objects for the virtualised network resource and allocated this resource. In addition, the VIM shall return to the VNFM information on the newly instantiated virtualised network resource plus any additional information about the allocate request operation. The VIM may also return intermediate status reports during the allocation process. If the operation was not successful, the VIM shall return to the VNFM appropriate error information.
        api_response = api_instance.vi_vnfm_allocate_virtualised_network_resource(body)
        print("The response of ViVnfmApi->vi_vnfm_allocate_virtualised_network_resource:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ViVnfmApi->vi_vnfm_allocate_virtualised_network_resource: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PbAllocateNetworkRequest**](PbAllocateNetworkRequest.md)|  | 

### Return type

[**PbAllocateNetworkResponse**](PbAllocateNetworkResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **vi_vnfm_create_compute_flavour**
> PbCreateComputeFlavourResponse vi_vnfm_create_compute_flavour(body)

This operation allows requesting the creation of a flavour as indicated by the consumer functional block. Result: After successful operation, the VIM has created the Compute Flavour. In addition, the VIM shall return to the VNFM information on the newly created Compute Flavour. If the operation was not successful, the VIM shall return to the VNFM appropriate error information.

### Example


```python
import kubevim_vivnfm_client
from kubevim_vivnfm_client.models.pb_create_compute_flavour_request import PbCreateComputeFlavourRequest
from kubevim_vivnfm_client.models.pb_create_compute_flavour_response import PbCreateComputeFlavourResponse
from kubevim_vivnfm_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = kubevim_vivnfm_client.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with kubevim_vivnfm_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = kubevim_vivnfm_client.ViVnfmApi(api_client)
    body = kubevim_vivnfm_client.PbCreateComputeFlavourRequest() # PbCreateComputeFlavourRequest | 

    try:
        # This operation allows requesting the creation of a flavour as indicated by the consumer functional block. Result: After successful operation, the VIM has created the Compute Flavour. In addition, the VIM shall return to the VNFM information on the newly created Compute Flavour. If the operation was not successful, the VIM shall return to the VNFM appropriate error information.
        api_response = api_instance.vi_vnfm_create_compute_flavour(body)
        print("The response of ViVnfmApi->vi_vnfm_create_compute_flavour:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ViVnfmApi->vi_vnfm_create_compute_flavour: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PbCreateComputeFlavourRequest**](PbCreateComputeFlavourRequest.md)|  | 

### Return type

[**PbCreateComputeFlavourResponse**](PbCreateComputeFlavourResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **vi_vnfm_delete_compute_flavour**
> object vi_vnfm_delete_compute_flavour(compute_flavour_id_value)

This operation allows deleting a Compute Flavour. Result: After successful operation, the VIM has deleted the Compute Flavour, so no new Virtualised Compute Resource can be allocated based on it. The already allocated Virtualised Compute Resources are not affected. If the operation was not successful, the VIM shall return to the VNFM appropriate error information.

### Example


```python
import kubevim_vivnfm_client
from kubevim_vivnfm_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = kubevim_vivnfm_client.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with kubevim_vivnfm_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = kubevim_vivnfm_client.ViVnfmApi(api_client)
    compute_flavour_id_value = 'compute_flavour_id_value_example' # str | UUID Identifier representation

    try:
        # This operation allows deleting a Compute Flavour. Result: After successful operation, the VIM has deleted the Compute Flavour, so no new Virtualised Compute Resource can be allocated based on it. The already allocated Virtualised Compute Resources are not affected. If the operation was not successful, the VIM shall return to the VNFM appropriate error information.
        api_response = api_instance.vi_vnfm_delete_compute_flavour(compute_flavour_id_value)
        print("The response of ViVnfmApi->vi_vnfm_delete_compute_flavour:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ViVnfmApi->vi_vnfm_delete_compute_flavour: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compute_flavour_id_value** | **str**| UUID Identifier representation | 

### Return type

**object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **vi_vnfm_query_compute_flavour**
> PbQueryComputeFlavourResponse vi_vnfm_query_compute_flavour(query_compute_flavour_filter_value=query_compute_flavour_filter_value)

This operation allows querying information about created Compute Flavours. Result: After successful operation, the VIM has queried the internal management objects for the Compute Flavours. The result of the query shall indicate with a standard success/error result if the query has been processed correctly. For a particular query, information about the Compute Flavours that the VNFM has access to and that are matching the filter shall be returned.

### Example


```python
import kubevim_vivnfm_client
from kubevim_vivnfm_client.models.pb_query_compute_flavour_response import PbQueryComputeFlavourResponse
from kubevim_vivnfm_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = kubevim_vivnfm_client.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with kubevim_vivnfm_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = kubevim_vivnfm_client.ViVnfmApi(api_client)
    query_compute_flavour_filter_value = 'query_compute_flavour_filter_value_example' # str |  (optional)

    try:
        # This operation allows querying information about created Compute Flavours. Result: After successful operation, the VIM has queried the internal management objects for the Compute Flavours. The result of the query shall indicate with a standard success/error result if the query has been processed correctly. For a particular query, information about the Compute Flavours that the VNFM has access to and that are matching the filter shall be returned.
        api_response = api_instance.vi_vnfm_query_compute_flavour(query_compute_flavour_filter_value=query_compute_flavour_filter_value)
        print("The response of ViVnfmApi->vi_vnfm_query_compute_flavour:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ViVnfmApi->vi_vnfm_query_compute_flavour: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query_compute_flavour_filter_value** | **str**|  | [optional] 

### Return type

[**PbQueryComputeFlavourResponse**](PbQueryComputeFlavourResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **vi_vnfm_query_image**
> PbQueryImageResponse vi_vnfm_query_image(software_image_id_value)

This operation allows querying the information about a specific software image in the image repository managed by the VIM. Result: As a result of this operation, the producer (VIM) shall indicate to the consumer (VNFM) whether or not it was possible to process the query.

### Example


```python
import kubevim_vivnfm_client
from kubevim_vivnfm_client.models.pb_query_image_response import PbQueryImageResponse
from kubevim_vivnfm_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = kubevim_vivnfm_client.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with kubevim_vivnfm_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = kubevim_vivnfm_client.ViVnfmApi(api_client)
    software_image_id_value = 'software_image_id_value_example' # str | UUID Identifier representation

    try:
        # This operation allows querying the information about a specific software image in the image repository managed by the VIM. Result: As a result of this operation, the producer (VIM) shall indicate to the consumer (VNFM) whether or not it was possible to process the query.
        api_response = api_instance.vi_vnfm_query_image(software_image_id_value)
        print("The response of ViVnfmApi->vi_vnfm_query_image:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ViVnfmApi->vi_vnfm_query_image: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **software_image_id_value** | **str**| UUID Identifier representation | 

### Return type

[**PbQueryImageResponse**](PbQueryImageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **vi_vnfm_query_image2**
> PbQueryImageResponse vi_vnfm_query_image2(body)

This operation allows querying the information about a specific software image in the image repository managed by the VIM. Result: As a result of this operation, the producer (VIM) shall indicate to the consumer (VNFM) whether or not it was possible to process the query.

### Example


```python
import kubevim_vivnfm_client
from kubevim_vivnfm_client.models.pb_query_image_request import PbQueryImageRequest
from kubevim_vivnfm_client.models.pb_query_image_response import PbQueryImageResponse
from kubevim_vivnfm_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = kubevim_vivnfm_client.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with kubevim_vivnfm_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = kubevim_vivnfm_client.ViVnfmApi(api_client)
    body = kubevim_vivnfm_client.PbQueryImageRequest() # PbQueryImageRequest | 

    try:
        # This operation allows querying the information about a specific software image in the image repository managed by the VIM. Result: As a result of this operation, the producer (VIM) shall indicate to the consumer (VNFM) whether or not it was possible to process the query.
        api_response = api_instance.vi_vnfm_query_image2(body)
        print("The response of ViVnfmApi->vi_vnfm_query_image2:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ViVnfmApi->vi_vnfm_query_image2: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PbQueryImageRequest**](PbQueryImageRequest.md)|  | 

### Return type

[**PbQueryImageResponse**](PbQueryImageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **vi_vnfm_query_images**
> PbQueryImagesResponse vi_vnfm_query_images(image_query_filter_value=image_query_filter_value)

This operation allows querying the information of software images in the image repository managed by the VIM. Result: As a result of this operation, the producer (VIM) shall indicate to the consumer (VNFM) whether or not it was possible to process the query

### Example


```python
import kubevim_vivnfm_client
from kubevim_vivnfm_client.models.pb_query_images_response import PbQueryImagesResponse
from kubevim_vivnfm_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = kubevim_vivnfm_client.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with kubevim_vivnfm_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = kubevim_vivnfm_client.ViVnfmApi(api_client)
    image_query_filter_value = 'image_query_filter_value_example' # str |  (optional)

    try:
        # This operation allows querying the information of software images in the image repository managed by the VIM. Result: As a result of this operation, the producer (VIM) shall indicate to the consumer (VNFM) whether or not it was possible to process the query
        api_response = api_instance.vi_vnfm_query_images(image_query_filter_value=image_query_filter_value)
        print("The response of ViVnfmApi->vi_vnfm_query_images:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ViVnfmApi->vi_vnfm_query_images: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **image_query_filter_value** | **str**|  | [optional] 

### Return type

[**PbQueryImagesResponse**](PbQueryImagesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **vi_vnfm_query_virtualised_network_resource**
> PbQueryNetworkResponse vi_vnfm_query_virtualised_network_resource(query_network_filter_value=query_network_filter_value, network_resource_type=network_resource_type)

This operation allows querying information about instantiated virtualised network resources. Result: After successful operation, the VIM has queried the internal management objects for the virtualised network resources. The result of the query shall indicate with a standard success/error result if the query has been processed correctly. For a particular query, information about the network resources that the VNFM has access to and that are matching the filter shall be returned.

### Example


```python
import kubevim_vivnfm_client
from kubevim_vivnfm_client.models.pb_query_network_response import PbQueryNetworkResponse
from kubevim_vivnfm_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = kubevim_vivnfm_client.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with kubevim_vivnfm_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = kubevim_vivnfm_client.ViVnfmApi(api_client)
    query_network_filter_value = 'query_network_filter_value_example' # str |  (optional)
    network_resource_type = NETWORK # str | Note: this message goes out of ETSI GS NFV-IFA 006 reference but it is required to identify network resource type while performing query. Later the filter will be applied to that network resource type. (optional) (default to NETWORK)

    try:
        # This operation allows querying information about instantiated virtualised network resources. Result: After successful operation, the VIM has queried the internal management objects for the virtualised network resources. The result of the query shall indicate with a standard success/error result if the query has been processed correctly. For a particular query, information about the network resources that the VNFM has access to and that are matching the filter shall be returned.
        api_response = api_instance.vi_vnfm_query_virtualised_network_resource(query_network_filter_value=query_network_filter_value, network_resource_type=network_resource_type)
        print("The response of ViVnfmApi->vi_vnfm_query_virtualised_network_resource:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ViVnfmApi->vi_vnfm_query_virtualised_network_resource: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query_network_filter_value** | **str**|  | [optional] 
 **network_resource_type** | **str**| Note: this message goes out of ETSI GS NFV-IFA 006 reference but it is required to identify network resource type while performing query. Later the filter will be applied to that network resource type. | [optional] [default to NETWORK]

### Return type

[**PbQueryNetworkResponse**](PbQueryNetworkResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **vi_vnfm_terminate_virtualised_network_resource**
> PbTerminateNetworkResponse vi_vnfm_terminate_virtualised_network_resource(network_resource_id_value)

This operation allows de-allocating and terminating one or more an instantiated virtualised network resource(s). When the operation is done on multiple ids, it is assumed to be best-effort, i.e. it can succeed for a subset of the ids, and fail for the remaining ones. Result: After successful operation, the VIM has terminated the virtualised network resources and removed the internal management objects for those resources. In addition, the VIM shall return to the VNFM information on the terminated virtualised network resource plus any additional information about the terminate request operation. If the operation was not successful, the VIM shall return to the VNFM appropriate error information.

Note(dmalovan): ETSI GS NFV-IFA 006 (7.4.1.5.4) Operation result attached above is not coresponds to the reality, since Ouput parameters defined in the (7.4.1.5.3) Output parameters block are not contains any (C) \"additional information about the terminated request operation\" and (C) \"appropriate error information\"

### Example


```python
import kubevim_vivnfm_client
from kubevim_vivnfm_client.models.pb_terminate_network_response import PbTerminateNetworkResponse
from kubevim_vivnfm_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = kubevim_vivnfm_client.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with kubevim_vivnfm_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = kubevim_vivnfm_client.ViVnfmApi(api_client)
    network_resource_id_value = 'network_resource_id_value_example' # str | UUID Identifier representation

    try:
        # This operation allows de-allocating and terminating one or more an instantiated virtualised network resource(s). When the operation is done on multiple ids, it is assumed to be best-effort, i.e. it can succeed for a subset of the ids, and fail for the remaining ones. Result: After successful operation, the VIM has terminated the virtualised network resources and removed the internal management objects for those resources. In addition, the VIM shall return to the VNFM information on the terminated virtualised network resource plus any additional information about the terminate request operation. If the operation was not successful, the VIM shall return to the VNFM appropriate error information.
        api_response = api_instance.vi_vnfm_terminate_virtualised_network_resource(network_resource_id_value)
        print("The response of ViVnfmApi->vi_vnfm_terminate_virtualised_network_resource:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ViVnfmApi->vi_vnfm_terminate_virtualised_network_resource: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **network_resource_id_value** | **str**| UUID Identifier representation | 

### Return type

[**PbTerminateNetworkResponse**](PbTerminateNetworkResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

