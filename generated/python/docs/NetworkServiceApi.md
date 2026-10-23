# kubevim_client.NetworkServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_network**](NetworkServiceApi.md#create_network) | **POST** /network/v1/networks | Creates a network.
[**create_subnet**](NetworkServiceApi.md#create_subnet) | **POST** /network/v1/subnets | Creates a subnet.
[**delete_network**](NetworkServiceApi.md#delete_network) | **DELETE** /network/v1/networks/{network_id} | Deletes a network and its subnets; may fail with RESOURCE_IN_USE while interfaces use the network.
[**delete_subnet**](NetworkServiceApi.md#delete_subnet) | **DELETE** /network/v1/subnets/{subnet_id} | Deletes a subnet: one IP version of a network. May fail with RESOURCE_IN_USE while interfaces use the network; the last subnet of a Kube-OVN network cannot be deleted (delete the network instead).
[**get_network**](NetworkServiceApi.md#get_network) | **GET** /network/v1/networks/{network_id} | Returns one network.
[**get_subnet**](NetworkServiceApi.md#get_subnet) | **GET** /network/v1/subnets/{subnet_id} | Returns one subnet.
[**list_networks**](NetworkServiceApi.md#list_networks) | **GET** /network/v1/networks | Returns the networks matching a filter.
[**list_subnets**](NetworkServiceApi.md#list_subnets) | **GET** /network/v1/subnets | Returns the subnets matching a filter.
[**network_get_api_major_versions**](NetworkServiceApi.md#network_get_api_major_versions) | **GET** /network/v1/api_versions | API versions this VIM serves (SOL 013 9.3).
[**network_get_api_versions**](NetworkServiceApi.md#network_get_api_versions) | **GET** /network/api_versions | API versions this VIM serves (SOL 013 9.3).
[**update_network**](NetworkServiceApi.md#update_network) | **PATCH** /network/v1/networks/{network_id} | Updates the name and metadata of a network.
[**update_subnet**](NetworkServiceApi.md#update_subnet) | **PATCH** /network/v1/subnets/{subnet_id} | Updates the name and metadata of a subnet.


# **create_network**
> Network create_network(create_network_request, version=version)

Creates a network.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.create_network_request import CreateNetworkRequest
from kubevim_client.models.network import Network
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
    api_instance = kubevim_client.NetworkServiceApi(api_client)
    create_network_request = kubevim_client.CreateNetworkRequest() # CreateNetworkRequest | 
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Creates a network.
        api_response = api_instance.create_network(create_network_request, version=version)
        print("The response of NetworkServiceApi->create_network:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling NetworkServiceApi->create_network: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **create_network_request** | [**CreateNetworkRequest**](CreateNetworkRequest.md)|  | 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**Network**](Network.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | A virtualised network: one layer 2 segment, with at most one IPv4 and one IPv6 subnet. |  * Location - URI of the created resource. <br>  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |
**0** | Error (SOL 013 6). |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_subnet**
> Subnet create_subnet(create_subnet_request, version=version)

Creates a subnet.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.create_subnet_request import CreateSubnetRequest
from kubevim_client.models.subnet import Subnet
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
    api_instance = kubevim_client.NetworkServiceApi(api_client)
    create_subnet_request = kubevim_client.CreateSubnetRequest() # CreateSubnetRequest | 
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Creates a subnet.
        api_response = api_instance.create_subnet(create_subnet_request, version=version)
        print("The response of NetworkServiceApi->create_subnet:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling NetworkServiceApi->create_subnet: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **create_subnet_request** | [**CreateSubnetRequest**](CreateSubnetRequest.md)|  | 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**Subnet**](Subnet.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | A virtualised subnet: one IP version of a network. |  * Location - URI of the created resource. <br>  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |
**0** | Error (SOL 013 6). |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_network**
> delete_network(network_id, version=version)

Deletes a network and its subnets; may fail with RESOURCE_IN_USE while interfaces use the network.

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
    api_instance = kubevim_client.NetworkServiceApi(api_client)
    network_id = 'network_id_example' # str | Identifier of the network.
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Deletes a network and its subnets; may fail with RESOURCE_IN_USE while interfaces use the network.
        api_instance.delete_network(network_id, version=version)
    except Exception as e:
        print("Exception when calling NetworkServiceApi->delete_network: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **network_id** | **str**| Identifier of the network. | 
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

# **delete_subnet**
> delete_subnet(subnet_id, version=version)

Deletes a subnet: one IP version of a network. May fail with RESOURCE_IN_USE while interfaces use the network; the last subnet of a Kube-OVN network cannot be deleted (delete the network instead).

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
    api_instance = kubevim_client.NetworkServiceApi(api_client)
    subnet_id = 'subnet_id_example' # str | Identifier of the subnet.
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Deletes a subnet: one IP version of a network. May fail with RESOURCE_IN_USE while interfaces use the network; the last subnet of a Kube-OVN network cannot be deleted (delete the network instead).
        api_instance.delete_subnet(subnet_id, version=version)
    except Exception as e:
        print("Exception when calling NetworkServiceApi->delete_subnet: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subnet_id** | **str**| Identifier of the subnet. | 
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

# **get_network**
> Network get_network(network_id, version=version)

Returns one network.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.network import Network
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
    api_instance = kubevim_client.NetworkServiceApi(api_client)
    network_id = 'network_id_example' # str | Identifier of the network.
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Returns one network.
        api_response = api_instance.get_network(network_id, version=version)
        print("The response of NetworkServiceApi->get_network:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling NetworkServiceApi->get_network: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **network_id** | **str**| Identifier of the network. | 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**Network**](Network.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A virtualised network: one layer 2 segment, with at most one IPv4 and one IPv6 subnet. |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |
**0** | Error (SOL 013 6). |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_subnet**
> Subnet get_subnet(subnet_id, version=version)

Returns one subnet.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.subnet import Subnet
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
    api_instance = kubevim_client.NetworkServiceApi(api_client)
    subnet_id = 'subnet_id_example' # str | Identifier of the subnet.
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Returns one subnet.
        api_response = api_instance.get_subnet(subnet_id, version=version)
        print("The response of NetworkServiceApi->get_subnet:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling NetworkServiceApi->get_subnet: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subnet_id** | **str**| Identifier of the subnet. | 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**Subnet**](Subnet.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A virtualised subnet: one IP version of a network. |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |
**0** | Error (SOL 013 6). |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_networks**
> List[Network] list_networks(filter=filter, nextpage_opaque_marker=nextpage_opaque_marker, version=version)

Returns the networks matching a filter.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.network import Network
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
    api_instance = kubevim_client.NetworkServiceApi(api_client)
    filter = 'filter_example' # str | Attribute-based filter expression (SOL 013 5.2). (optional)
    nextpage_opaque_marker = 'nextpage_opaque_marker_example' # str | Marker of the page to return, from a previous response (SOL 013 5.4). (optional)
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Returns the networks matching a filter.
        api_response = api_instance.list_networks(filter=filter, nextpage_opaque_marker=nextpage_opaque_marker, version=version)
        print("The response of NetworkServiceApi->list_networks:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling NetworkServiceApi->list_networks: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **str**| Attribute-based filter expression (SOL 013 5.2). | [optional] 
 **nextpage_opaque_marker** | **str**| Marker of the page to return, from a previous response (SOL 013 5.4). | [optional] 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**List[Network]**](Network.md)

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

# **list_subnets**
> List[Subnet] list_subnets(filter=filter, nextpage_opaque_marker=nextpage_opaque_marker, version=version)

Returns the subnets matching a filter.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.subnet import Subnet
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
    api_instance = kubevim_client.NetworkServiceApi(api_client)
    filter = 'filter_example' # str | Attribute-based filter expression (SOL 013 5.2). (optional)
    nextpage_opaque_marker = 'nextpage_opaque_marker_example' # str | Marker of the page to return, from a previous response (SOL 013 5.4). (optional)
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Returns the subnets matching a filter.
        api_response = api_instance.list_subnets(filter=filter, nextpage_opaque_marker=nextpage_opaque_marker, version=version)
        print("The response of NetworkServiceApi->list_subnets:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling NetworkServiceApi->list_subnets: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **str**| Attribute-based filter expression (SOL 013 5.2). | [optional] 
 **nextpage_opaque_marker** | **str**| Marker of the page to return, from a previous response (SOL 013 5.4). | [optional] 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**List[Subnet]**](Subnet.md)

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

# **network_get_api_major_versions**
> ApiVersionInformation network_get_api_major_versions()

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
    api_instance = kubevim_client.NetworkServiceApi(api_client)

    try:
        # API versions this VIM serves (SOL 013 9.3).
        api_response = api_instance.network_get_api_major_versions()
        print("The response of NetworkServiceApi->network_get_api_major_versions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling NetworkServiceApi->network_get_api_major_versions: %s\n" % e)
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

# **network_get_api_versions**
> ApiVersionInformation network_get_api_versions()

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
    api_instance = kubevim_client.NetworkServiceApi(api_client)

    try:
        # API versions this VIM serves (SOL 013 9.3).
        api_response = api_instance.network_get_api_versions()
        print("The response of NetworkServiceApi->network_get_api_versions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling NetworkServiceApi->network_get_api_versions: %s\n" % e)
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

# **update_network**
> Network update_network(network_id, network_update, version=version)

Updates the name and metadata of a network.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.network import Network
from kubevim_client.models.network_update import NetworkUpdate
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
    api_instance = kubevim_client.NetworkServiceApi(api_client)
    network_id = 'network_id_example' # str | Identifier of the network.
    network_update = kubevim_client.NetworkUpdate() # NetworkUpdate | 
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Updates the name and metadata of a network.
        api_response = api_instance.update_network(network_id, network_update, version=version)
        print("The response of NetworkServiceApi->update_network:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling NetworkServiceApi->update_network: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **network_id** | **str**| Identifier of the network. | 
 **network_update** | [**NetworkUpdate**](NetworkUpdate.md)|  | 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**Network**](Network.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A virtualised network: one layer 2 segment, with at most one IPv4 and one IPv6 subnet. |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |
**0** | Error (SOL 013 6). |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_subnet**
> Subnet update_subnet(subnet_id, subnet_update, version=version)

Updates the name and metadata of a subnet.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.subnet import Subnet
from kubevim_client.models.subnet_update import SubnetUpdate
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
    api_instance = kubevim_client.NetworkServiceApi(api_client)
    subnet_id = 'subnet_id_example' # str | Identifier of the subnet.
    subnet_update = kubevim_client.SubnetUpdate() # SubnetUpdate | 
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Updates the name and metadata of a subnet.
        api_response = api_instance.update_subnet(subnet_id, subnet_update, version=version)
        print("The response of NetworkServiceApi->update_subnet:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling NetworkServiceApi->update_subnet: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subnet_id** | **str**| Identifier of the subnet. | 
 **subnet_update** | [**SubnetUpdate**](SubnetUpdate.md)|  | 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**Subnet**](Subnet.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A virtualised subnet: one IP version of a network. |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |
**0** | Error (SOL 013 6). |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

