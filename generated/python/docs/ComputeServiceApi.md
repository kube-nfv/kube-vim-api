# kubevim_client.ComputeServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**attach_volume**](ComputeServiceApi.md#attach_volume) | **POST** /compute/v1/computes/{compute_id}/attachVolume | Attaches a volume to a compute resource.
[**compute_get_api_major_versions**](ComputeServiceApi.md#compute_get_api_major_versions) | **GET** /compute/v1/api_versions | API versions this VIM serves (SOL 013 9.3).
[**compute_get_api_versions**](ComputeServiceApi.md#compute_get_api_versions) | **GET** /compute/api_versions | API versions this VIM serves (SOL 013 9.3).
[**create_compute**](ComputeServiceApi.md#create_compute) | **POST** /compute/v1/computes | Creates a compute resource.
[**create_compute_snapshot**](ComputeServiceApi.md#create_compute_snapshot) | **POST** /compute/v1/computes/{compute_id}/snapshots | Creates a snapshot of a compute resource.
[**create_placement_group**](ComputeServiceApi.md#create_placement_group) | **POST** /compute/v1/placementGroups | Creates a placement group.
[**delete_compute**](ComputeServiceApi.md#delete_compute) | **DELETE** /compute/v1/computes/{compute_id} | Deletes a compute resource with its boot disk and the disks from its flavour. Attached volumes are detached and kept unless delete_volumes is set; snapshots are kept.
[**delete_compute_snapshot**](ComputeServiceApi.md#delete_compute_snapshot) | **DELETE** /compute/v1/computes/{compute_id}/snapshots/{snapshot_id} | Deletes a snapshot of a compute resource.
[**delete_placement_group**](ComputeServiceApi.md#delete_placement_group) | **DELETE** /compute/v1/placementGroups/{placement_group_id} | Deletes a placement group that has no members.
[**detach_volume**](ComputeServiceApi.md#detach_volume) | **POST** /compute/v1/computes/{compute_id}/detachVolume | Detaches a volume from a compute resource.
[**get_compute**](ComputeServiceApi.md#get_compute) | **GET** /compute/v1/computes/{compute_id} | Returns one compute resource.
[**get_compute_snapshot**](ComputeServiceApi.md#get_compute_snapshot) | **GET** /compute/v1/computes/{compute_id}/snapshots/{snapshot_id} | Returns one snapshot of a compute resource, also after the compute resource is deleted.
[**get_placement_group**](ComputeServiceApi.md#get_placement_group) | **GET** /compute/v1/placementGroups/{placement_group_id} | Returns one placement group.
[**list_compute_snapshots**](ComputeServiceApi.md#list_compute_snapshots) | **GET** /compute/v1/computes/{compute_id}/snapshots | Returns the snapshots of a compute resource, or of all compute resources with compute_id \&quot;-\&quot;.
[**list_computes**](ComputeServiceApi.md#list_computes) | **GET** /compute/v1/computes | Returns the compute resources matching a filter.
[**list_placement_groups**](ComputeServiceApi.md#list_placement_groups) | **GET** /compute/v1/placementGroups | Returns the placement groups.
[**migrate_compute**](ComputeServiceApi.md#migrate_compute) | **POST** /compute/v1/computes/{compute_id}/migrate | Moves a running compute resource to another host (live migration).
[**operate_compute**](ComputeServiceApi.md#operate_compute) | **POST** /compute/v1/computes/{compute_id}/operate | Changes the running state of a compute resource, e.g. start, stop or reboot.
[**revert_compute_snapshot**](ComputeServiceApi.md#revert_compute_snapshot) | **POST** /compute/v1/computes/{compute_id}/snapshots/{snapshot_id}/revert | Reverts a stopped compute resource to a snapshot; not possible once the compute resource is deleted.
[**scale_compute**](ComputeServiceApi.md#scale_compute) | **POST** /compute/v1/computes/{compute_id}/scale | Changes the flavour of a compute resource; live where KubeVirt can hotplug, otherwise restart_required is set. Disks that are larger in the new flavour are grown; shrinking a disk or changing the number, type, boot flag or storage class of the disks is rejected.
[**update_compute**](ComputeServiceApi.md#update_compute) | **PATCH** /compute/v1/computes/{compute_id} | Updates a compute resource: its name and metadata, and adds or removes network interfaces.


# **attach_volume**
> Compute attach_volume(compute_id, attach_volume_request, version=version)

Attaches a volume to a compute resource.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.attach_volume_request import AttachVolumeRequest
from kubevim_client.models.compute import Compute
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
    api_instance = kubevim_client.ComputeServiceApi(api_client)
    compute_id = 'compute_id_example' # str | Identifier of the compute resource.
    attach_volume_request = kubevim_client.AttachVolumeRequest() # AttachVolumeRequest | 
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Attaches a volume to a compute resource.
        api_response = api_instance.attach_volume(compute_id, attach_volume_request, version=version)
        print("The response of ComputeServiceApi->attach_volume:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputeServiceApi->attach_volume: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compute_id** | **str**| Identifier of the compute resource. | 
 **attach_volume_request** | [**AttachVolumeRequest**](AttachVolumeRequest.md)|  | 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**Compute**](Compute.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A virtualised compute resource: a virtual machine. |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |
**0** | Error (SOL 013 6). |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **compute_get_api_major_versions**
> ApiVersionInformation compute_get_api_major_versions()

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
    api_instance = kubevim_client.ComputeServiceApi(api_client)

    try:
        # API versions this VIM serves (SOL 013 9.3).
        api_response = api_instance.compute_get_api_major_versions()
        print("The response of ComputeServiceApi->compute_get_api_major_versions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputeServiceApi->compute_get_api_major_versions: %s\n" % e)
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

# **compute_get_api_versions**
> ApiVersionInformation compute_get_api_versions()

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
    api_instance = kubevim_client.ComputeServiceApi(api_client)

    try:
        # API versions this VIM serves (SOL 013 9.3).
        api_response = api_instance.compute_get_api_versions()
        print("The response of ComputeServiceApi->compute_get_api_versions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputeServiceApi->compute_get_api_versions: %s\n" % e)
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

# **create_compute**
> Compute create_compute(create_compute_request, version=version)

Creates a compute resource.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.compute import Compute
from kubevim_client.models.create_compute_request import CreateComputeRequest
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
    api_instance = kubevim_client.ComputeServiceApi(api_client)
    create_compute_request = kubevim_client.CreateComputeRequest() # CreateComputeRequest | 
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Creates a compute resource.
        api_response = api_instance.create_compute(create_compute_request, version=version)
        print("The response of ComputeServiceApi->create_compute:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputeServiceApi->create_compute: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **create_compute_request** | [**CreateComputeRequest**](CreateComputeRequest.md)|  | 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**Compute**](Compute.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | A virtualised compute resource: a virtual machine. |  * Location - URI of the created resource. <br>  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |
**0** | Error (SOL 013 6). |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_compute_snapshot**
> ComputeSnapshot create_compute_snapshot(compute_id, create_compute_snapshot_request, version=version)

Creates a snapshot of a compute resource.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.compute_snapshot import ComputeSnapshot
from kubevim_client.models.create_compute_snapshot_request import CreateComputeSnapshotRequest
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
    api_instance = kubevim_client.ComputeServiceApi(api_client)
    compute_id = 'compute_id_example' # str | Compute resource to take the snapshot of.
    create_compute_snapshot_request = kubevim_client.CreateComputeSnapshotRequest() # CreateComputeSnapshotRequest | 
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Creates a snapshot of a compute resource.
        api_response = api_instance.create_compute_snapshot(compute_id, create_compute_snapshot_request, version=version)
        print("The response of ComputeServiceApi->create_compute_snapshot:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputeServiceApi->create_compute_snapshot: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compute_id** | **str**| Compute resource to take the snapshot of. | 
 **create_compute_snapshot_request** | [**CreateComputeSnapshotRequest**](CreateComputeSnapshotRequest.md)|  | 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**ComputeSnapshot**](ComputeSnapshot.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | A snapshot of a compute resource (KubeVirt VirtualMachineSnapshot). |  * Location - URI of the created resource. <br>  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |
**0** | Error (SOL 013 6). |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_placement_group**
> PlacementGroup create_placement_group(create_placement_group_request, version=version)

Creates a placement group.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.create_placement_group_request import CreatePlacementGroupRequest
from kubevim_client.models.placement_group import PlacementGroup
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
    api_instance = kubevim_client.ComputeServiceApi(api_client)
    create_placement_group_request = kubevim_client.CreatePlacementGroupRequest() # CreatePlacementGroupRequest | 
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Creates a placement group.
        api_response = api_instance.create_placement_group(create_placement_group_request, version=version)
        print("The response of ComputeServiceApi->create_placement_group:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputeServiceApi->create_placement_group: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **create_placement_group_request** | [**CreatePlacementGroupRequest**](CreatePlacementGroupRequest.md)|  | 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**PlacementGroup**](PlacementGroup.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | A group of compute resources placed together or apart (decision 0005). |  * Location - URI of the created resource. <br>  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |
**0** | Error (SOL 013 6). |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_compute**
> delete_compute(compute_id, delete_volumes=delete_volumes, version=version)

Deletes a compute resource with its boot disk and the disks from its flavour. Attached volumes are detached and kept unless delete_volumes is set; snapshots are kept.

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
    api_instance = kubevim_client.ComputeServiceApi(api_client)
    compute_id = 'compute_id_example' # str | Identifier of the compute resource.
    delete_volumes = True # bool | Also delete the attached volumes created with the storage API; they are detached and kept when unset. (optional)
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Deletes a compute resource with its boot disk and the disks from its flavour. Attached volumes are detached and kept unless delete_volumes is set; snapshots are kept.
        api_instance.delete_compute(compute_id, delete_volumes=delete_volumes, version=version)
    except Exception as e:
        print("Exception when calling ComputeServiceApi->delete_compute: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compute_id** | **str**| Identifier of the compute resource. | 
 **delete_volumes** | **bool**| Also delete the attached volumes created with the storage API; they are detached and kept when unset. | [optional] 
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

# **delete_compute_snapshot**
> delete_compute_snapshot(compute_id, snapshot_id, version=version)

Deletes a snapshot of a compute resource.

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
    api_instance = kubevim_client.ComputeServiceApi(api_client)
    compute_id = 'compute_id_example' # str | Compute resource the snapshot belongs to.
    snapshot_id = 'snapshot_id_example' # str | Identifier of the snapshot.
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Deletes a snapshot of a compute resource.
        api_instance.delete_compute_snapshot(compute_id, snapshot_id, version=version)
    except Exception as e:
        print("Exception when calling ComputeServiceApi->delete_compute_snapshot: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compute_id** | **str**| Compute resource the snapshot belongs to. | 
 **snapshot_id** | **str**| Identifier of the snapshot. | 
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

# **delete_placement_group**
> delete_placement_group(placement_group_id, version=version)

Deletes a placement group that has no members.

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
    api_instance = kubevim_client.ComputeServiceApi(api_client)
    placement_group_id = 'placement_group_id_example' # str | Identifier of the group.
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Deletes a placement group that has no members.
        api_instance.delete_placement_group(placement_group_id, version=version)
    except Exception as e:
        print("Exception when calling ComputeServiceApi->delete_placement_group: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **placement_group_id** | **str**| Identifier of the group. | 
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

# **detach_volume**
> Compute detach_volume(compute_id, detach_volume_request, version=version)

Detaches a volume from a compute resource.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.compute import Compute
from kubevim_client.models.detach_volume_request import DetachVolumeRequest
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
    api_instance = kubevim_client.ComputeServiceApi(api_client)
    compute_id = 'compute_id_example' # str | Identifier of the compute resource.
    detach_volume_request = kubevim_client.DetachVolumeRequest() # DetachVolumeRequest | 
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Detaches a volume from a compute resource.
        api_response = api_instance.detach_volume(compute_id, detach_volume_request, version=version)
        print("The response of ComputeServiceApi->detach_volume:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputeServiceApi->detach_volume: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compute_id** | **str**| Identifier of the compute resource. | 
 **detach_volume_request** | [**DetachVolumeRequest**](DetachVolumeRequest.md)|  | 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**Compute**](Compute.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A virtualised compute resource: a virtual machine. |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |
**0** | Error (SOL 013 6). |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_compute**
> Compute get_compute(compute_id, version=version)

Returns one compute resource.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.compute import Compute
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
    api_instance = kubevim_client.ComputeServiceApi(api_client)
    compute_id = 'compute_id_example' # str | Identifier of the compute resource.
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Returns one compute resource.
        api_response = api_instance.get_compute(compute_id, version=version)
        print("The response of ComputeServiceApi->get_compute:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputeServiceApi->get_compute: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compute_id** | **str**| Identifier of the compute resource. | 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**Compute**](Compute.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A virtualised compute resource: a virtual machine. |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |
**0** | Error (SOL 013 6). |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_compute_snapshot**
> ComputeSnapshot get_compute_snapshot(compute_id, snapshot_id, version=version)

Returns one snapshot of a compute resource, also after the compute resource is deleted.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.compute_snapshot import ComputeSnapshot
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
    api_instance = kubevim_client.ComputeServiceApi(api_client)
    compute_id = 'compute_id_example' # str | Compute resource the snapshot belongs to.
    snapshot_id = 'snapshot_id_example' # str | Identifier of the snapshot.
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Returns one snapshot of a compute resource, also after the compute resource is deleted.
        api_response = api_instance.get_compute_snapshot(compute_id, snapshot_id, version=version)
        print("The response of ComputeServiceApi->get_compute_snapshot:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputeServiceApi->get_compute_snapshot: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compute_id** | **str**| Compute resource the snapshot belongs to. | 
 **snapshot_id** | **str**| Identifier of the snapshot. | 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**ComputeSnapshot**](ComputeSnapshot.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A snapshot of a compute resource (KubeVirt VirtualMachineSnapshot). |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |
**0** | Error (SOL 013 6). |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_placement_group**
> PlacementGroup get_placement_group(placement_group_id, version=version)

Returns one placement group.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.placement_group import PlacementGroup
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
    api_instance = kubevim_client.ComputeServiceApi(api_client)
    placement_group_id = 'placement_group_id_example' # str | Identifier of the group.
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Returns one placement group.
        api_response = api_instance.get_placement_group(placement_group_id, version=version)
        print("The response of ComputeServiceApi->get_placement_group:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputeServiceApi->get_placement_group: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **placement_group_id** | **str**| Identifier of the group. | 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**PlacementGroup**](PlacementGroup.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A group of compute resources placed together or apart (decision 0005). |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |
**0** | Error (SOL 013 6). |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_compute_snapshots**
> List[ComputeSnapshot] list_compute_snapshots(compute_id, filter=filter, nextpage_opaque_marker=nextpage_opaque_marker, version=version)

Returns the snapshots of a compute resource, or of all compute resources with compute_id \"-\".

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.compute_snapshot import ComputeSnapshot
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
    api_instance = kubevim_client.ComputeServiceApi(api_client)
    compute_id = 'compute_id_example' # str | Compute resource the snapshots belong to, or \"-\" for all compute resources (AIP-159).
    filter = 'filter_example' # str | Attribute-based filter expression (SOL 013 5.2). (optional)
    nextpage_opaque_marker = 'nextpage_opaque_marker_example' # str | Marker of the page to return, from a previous response (SOL 013 5.4). (optional)
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Returns the snapshots of a compute resource, or of all compute resources with compute_id \"-\".
        api_response = api_instance.list_compute_snapshots(compute_id, filter=filter, nextpage_opaque_marker=nextpage_opaque_marker, version=version)
        print("The response of ComputeServiceApi->list_compute_snapshots:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputeServiceApi->list_compute_snapshots: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compute_id** | **str**| Compute resource the snapshots belong to, or \&quot;-\&quot; for all compute resources (AIP-159). | 
 **filter** | **str**| Attribute-based filter expression (SOL 013 5.2). | [optional] 
 **nextpage_opaque_marker** | **str**| Marker of the page to return, from a previous response (SOL 013 5.4). | [optional] 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**List[ComputeSnapshot]**](ComputeSnapshot.md)

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

# **list_computes**
> List[Compute] list_computes(filter=filter, nextpage_opaque_marker=nextpage_opaque_marker, version=version)

Returns the compute resources matching a filter.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.compute import Compute
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
    api_instance = kubevim_client.ComputeServiceApi(api_client)
    filter = 'filter_example' # str | Attribute-based filter expression (SOL 013 5.2). (optional)
    nextpage_opaque_marker = 'nextpage_opaque_marker_example' # str | Marker of the page to return, from a previous response (SOL 013 5.4). (optional)
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Returns the compute resources matching a filter.
        api_response = api_instance.list_computes(filter=filter, nextpage_opaque_marker=nextpage_opaque_marker, version=version)
        print("The response of ComputeServiceApi->list_computes:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputeServiceApi->list_computes: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **str**| Attribute-based filter expression (SOL 013 5.2). | [optional] 
 **nextpage_opaque_marker** | **str**| Marker of the page to return, from a previous response (SOL 013 5.4). | [optional] 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**List[Compute]**](Compute.md)

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

# **list_placement_groups**
> List[PlacementGroup] list_placement_groups(filter=filter, nextpage_opaque_marker=nextpage_opaque_marker, version=version)

Returns the placement groups.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.placement_group import PlacementGroup
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
    api_instance = kubevim_client.ComputeServiceApi(api_client)
    filter = 'filter_example' # str | Attribute-based filter expression (SOL 013 5.2). (optional)
    nextpage_opaque_marker = 'nextpage_opaque_marker_example' # str | Marker of the page to return, from a previous response (SOL 013 5.4). (optional)
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Returns the placement groups.
        api_response = api_instance.list_placement_groups(filter=filter, nextpage_opaque_marker=nextpage_opaque_marker, version=version)
        print("The response of ComputeServiceApi->list_placement_groups:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputeServiceApi->list_placement_groups: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **str**| Attribute-based filter expression (SOL 013 5.2). | [optional] 
 **nextpage_opaque_marker** | **str**| Marker of the page to return, from a previous response (SOL 013 5.4). | [optional] 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**List[PlacementGroup]**](PlacementGroup.md)

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

# **migrate_compute**
> Compute migrate_compute(compute_id, body, version=version)

Moves a running compute resource to another host (live migration).

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.compute import Compute
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
    api_instance = kubevim_client.ComputeServiceApi(api_client)
    compute_id = 'compute_id_example' # str | Identifier of the compute resource.
    body = None # object | 
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Moves a running compute resource to another host (live migration).
        api_response = api_instance.migrate_compute(compute_id, body, version=version)
        print("The response of ComputeServiceApi->migrate_compute:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputeServiceApi->migrate_compute: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compute_id** | **str**| Identifier of the compute resource. | 
 **body** | **object**|  | 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**Compute**](Compute.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A virtualised compute resource: a virtual machine. |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |
**0** | Error (SOL 013 6). |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **operate_compute**
> Compute operate_compute(compute_id, operate_compute_request, version=version)

Changes the running state of a compute resource, e.g. start, stop or reboot.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.compute import Compute
from kubevim_client.models.operate_compute_request import OperateComputeRequest
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
    api_instance = kubevim_client.ComputeServiceApi(api_client)
    compute_id = 'compute_id_example' # str | Identifier of the compute resource.
    operate_compute_request = kubevim_client.OperateComputeRequest() # OperateComputeRequest | 
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Changes the running state of a compute resource, e.g. start, stop or reboot.
        api_response = api_instance.operate_compute(compute_id, operate_compute_request, version=version)
        print("The response of ComputeServiceApi->operate_compute:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputeServiceApi->operate_compute: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compute_id** | **str**| Identifier of the compute resource. | 
 **operate_compute_request** | [**OperateComputeRequest**](OperateComputeRequest.md)|  | 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**Compute**](Compute.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A virtualised compute resource: a virtual machine. |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |
**0** | Error (SOL 013 6). |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **revert_compute_snapshot**
> Compute revert_compute_snapshot(compute_id, snapshot_id, body, version=version)

Reverts a stopped compute resource to a snapshot; not possible once the compute resource is deleted.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.compute import Compute
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
    api_instance = kubevim_client.ComputeServiceApi(api_client)
    compute_id = 'compute_id_example' # str | Compute resource to revert; it must be stopped.
    snapshot_id = 'snapshot_id_example' # str | Identifier of the snapshot to revert to.
    body = None # object | 
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Reverts a stopped compute resource to a snapshot; not possible once the compute resource is deleted.
        api_response = api_instance.revert_compute_snapshot(compute_id, snapshot_id, body, version=version)
        print("The response of ComputeServiceApi->revert_compute_snapshot:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputeServiceApi->revert_compute_snapshot: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compute_id** | **str**| Compute resource to revert; it must be stopped. | 
 **snapshot_id** | **str**| Identifier of the snapshot to revert to. | 
 **body** | **object**|  | 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**Compute**](Compute.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A virtualised compute resource: a virtual machine. |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |
**0** | Error (SOL 013 6). |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **scale_compute**
> Compute scale_compute(compute_id, scale_compute_request, version=version)

Changes the flavour of a compute resource; live where KubeVirt can hotplug, otherwise restart_required is set. Disks that are larger in the new flavour are grown; shrinking a disk or changing the number, type, boot flag or storage class of the disks is rejected.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.compute import Compute
from kubevim_client.models.scale_compute_request import ScaleComputeRequest
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
    api_instance = kubevim_client.ComputeServiceApi(api_client)
    compute_id = 'compute_id_example' # str | Identifier of the compute resource.
    scale_compute_request = kubevim_client.ScaleComputeRequest() # ScaleComputeRequest | 
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Changes the flavour of a compute resource; live where KubeVirt can hotplug, otherwise restart_required is set. Disks that are larger in the new flavour are grown; shrinking a disk or changing the number, type, boot flag or storage class of the disks is rejected.
        api_response = api_instance.scale_compute(compute_id, scale_compute_request, version=version)
        print("The response of ComputeServiceApi->scale_compute:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputeServiceApi->scale_compute: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compute_id** | **str**| Identifier of the compute resource. | 
 **scale_compute_request** | [**ScaleComputeRequest**](ScaleComputeRequest.md)|  | 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**Compute**](Compute.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A virtualised compute resource: a virtual machine. |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |
**0** | Error (SOL 013 6). |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_compute**
> Compute update_compute(compute_id, compute_update, version=version)

Updates a compute resource: its name and metadata, and adds or removes network interfaces.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.compute import Compute
from kubevim_client.models.compute_update import ComputeUpdate
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
    api_instance = kubevim_client.ComputeServiceApi(api_client)
    compute_id = 'compute_id_example' # str | Identifier of the compute resource.
    compute_update = kubevim_client.ComputeUpdate() # ComputeUpdate | 
    version = 'version_example' # str | API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. (optional)

    try:
        # Updates a compute resource: its name and metadata, and adds or removes network interfaces.
        api_response = api_instance.update_compute(compute_id, compute_update, version=version)
        print("The response of ComputeServiceApi->update_compute:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputeServiceApi->update_compute: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compute_id** | **str**| Identifier of the compute resource. | 
 **compute_update** | [**ComputeUpdate**](ComputeUpdate.md)|  | 
 **version** | **str**| API version requested by the consumer, MAJOR.MINOR.PATCH (SOL 013 9.4); any version with the served major version is accepted. Required by the server; set it once on the client, e.g. as a default header. | [optional] 

### Return type

[**Compute**](Compute.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A virtualised compute resource: a virtual machine. |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |
**0** | Error (SOL 013 6). |  * Version - API version served, with the implementation suffix, e.g. 1.0.0-impl:kubenfv.io:kube-vim:0.4.0 (SOL 013 9.1.2, 9.4). <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

