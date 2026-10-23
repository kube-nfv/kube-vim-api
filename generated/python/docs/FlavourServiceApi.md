# kubevim_client.FlavourServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_flavour**](FlavourServiceApi.md#create_flavour) | **POST** /flavour/v1/flavours | Creates a compute flavour.
[**delete_flavour**](FlavourServiceApi.md#delete_flavour) | **DELETE** /flavour/v1/flavours/{flavour_id} | Deletes a compute flavour. Compute resources already created from it are not affected.
[**flavour_get_api_major_versions**](FlavourServiceApi.md#flavour_get_api_major_versions) | **GET** /flavour/v1/api_versions | API versions this VIM serves (SOL 013 9.3).
[**flavour_get_api_versions**](FlavourServiceApi.md#flavour_get_api_versions) | **GET** /flavour/api_versions | API versions this VIM serves (SOL 013 9.3).
[**get_flavour**](FlavourServiceApi.md#get_flavour) | **GET** /flavour/v1/flavours/{flavour_id} | Returns one compute flavour.
[**list_flavours**](FlavourServiceApi.md#list_flavours) | **GET** /flavour/v1/flavours | Returns the compute flavours matching a filter.


# **create_flavour**
> Flavour create_flavour(create_flavour_request, version=version)

Creates a compute flavour.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.create_flavour_request import CreateFlavourRequest
from kubevim_client.models.flavour import Flavour
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
    api_instance = kubevim_client.FlavourServiceApi(api_client)
    create_flavour_request = kubevim_client.CreateFlavourRequest() # CreateFlavourRequest | 
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Creates a compute flavour.
        api_response = api_instance.create_flavour(create_flavour_request, version=version)
        print("The response of FlavourServiceApi->create_flavour:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FlavourServiceApi->create_flavour: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **create_flavour_request** | [**CreateFlavourRequest**](CreateFlavourRequest.md)|  | 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**Flavour**](Flavour.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Memory, CPU and storage of the compute resources created from it. |  * Location - URI of the created resource. <br>  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |
**0** | Error (SOL 013 6). |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_flavour**
> delete_flavour(flavour_id, version=version)

Deletes a compute flavour. Compute resources already created from it are not affected.

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
    api_instance = kubevim_client.FlavourServiceApi(api_client)
    flavour_id = 'flavour_id_example' # str | Identifier of the flavour.
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Deletes a compute flavour. Compute resources already created from it are not affected.
        api_instance.delete_flavour(flavour_id, version=version)
    except Exception as e:
        print("Exception when calling FlavourServiceApi->delete_flavour: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flavour_id** | **str**| Identifier of the flavour. | 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

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
**204** | No Content |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |
**0** | Error (SOL 013 6). |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **flavour_get_api_major_versions**
> ApiVersionInformation flavour_get_api_major_versions()

API versions this VIM serves (SOL 013 9.3).

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.api_version_information import ApiVersionInformation
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
    api_instance = kubevim_client.FlavourServiceApi(api_client)

    try:
        # API versions this VIM serves (SOL 013 9.3).
        api_response = api_instance.flavour_get_api_major_versions()
        print("The response of FlavourServiceApi->flavour_get_api_major_versions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FlavourServiceApi->flavour_get_api_major_versions: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**ApiVersionInformation**](ApiVersionInformation.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**0** | Error (SOL 013 6). |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **flavour_get_api_versions**
> ApiVersionInformation flavour_get_api_versions()

API versions this VIM serves (SOL 013 9.3).

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.api_version_information import ApiVersionInformation
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
    api_instance = kubevim_client.FlavourServiceApi(api_client)

    try:
        # API versions this VIM serves (SOL 013 9.3).
        api_response = api_instance.flavour_get_api_versions()
        print("The response of FlavourServiceApi->flavour_get_api_versions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FlavourServiceApi->flavour_get_api_versions: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**ApiVersionInformation**](ApiVersionInformation.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**0** | Error (SOL 013 6). |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_flavour**
> Flavour get_flavour(flavour_id, version=version)

Returns one compute flavour.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.flavour import Flavour
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
    api_instance = kubevim_client.FlavourServiceApi(api_client)
    flavour_id = 'flavour_id_example' # str | Identifier of the flavour.
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Returns one compute flavour.
        api_response = api_instance.get_flavour(flavour_id, version=version)
        print("The response of FlavourServiceApi->get_flavour:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FlavourServiceApi->get_flavour: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flavour_id** | **str**| Identifier of the flavour. | 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**Flavour**](Flavour.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Memory, CPU and storage of the compute resources created from it. |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |
**0** | Error (SOL 013 6). |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_flavours**
> List[Flavour] list_flavours(filter=filter, nextpage_opaque_marker=nextpage_opaque_marker, version=version)

Returns the compute flavours matching a filter.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.flavour import Flavour
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
    api_instance = kubevim_client.FlavourServiceApi(api_client)
    filter = 'filter_example' # str | Attribute-based filter expression (SOL 013 5.2). (optional)
    nextpage_opaque_marker = 'nextpage_opaque_marker_example' # str | Marker of the page to return, from a previous response (SOL 013 5.4). (optional)
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Returns the compute flavours matching a filter.
        api_response = api_instance.list_flavours(filter=filter, nextpage_opaque_marker=nextpage_opaque_marker, version=version)
        print("The response of FlavourServiceApi->list_flavours:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FlavourServiceApi->list_flavours: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **str**| Attribute-based filter expression (SOL 013 5.2). | [optional] 
 **nextpage_opaque_marker** | **str**| Marker of the page to return, from a previous response (SOL 013 5.4). | [optional] 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**List[Flavour]**](Flavour.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  * Link - URI of the next page, rel&#x3D;\&quot;next\&quot;, carrying nextpage_opaque_marker; absent on the last page (SOL 013 5.4.2). <br>  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |
**0** | Error (SOL 013 6). |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

