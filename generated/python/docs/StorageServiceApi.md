# kubevim_client.StorageServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_volume**](StorageServiceApi.md#create_volume) | **POST** /storage/v1/volumes | Creates a volume.
[**create_volume_snapshot**](StorageServiceApi.md#create_volume_snapshot) | **POST** /storage/v1/volumes/{volume_id}/snapshots | Creates a snapshot of a volume.
[**delete_volume**](StorageServiceApi.md#delete_volume) | **DELETE** /storage/v1/volumes/{volume_id} | Deletes a volume; may fail with RESOURCE_IN_USE while the volume is attached. Snapshots are kept.
[**delete_volume_snapshot**](StorageServiceApi.md#delete_volume_snapshot) | **DELETE** /storage/v1/volumes/{volume_id}/snapshots/{snapshot_id} | Deletes a snapshot of a volume.
[**get_volume**](StorageServiceApi.md#get_volume) | **GET** /storage/v1/volumes/{volume_id} | Returns one volume.
[**get_volume_snapshot**](StorageServiceApi.md#get_volume_snapshot) | **GET** /storage/v1/volumes/{volume_id}/snapshots/{snapshot_id} | Returns one snapshot of a volume, also after the volume is deleted.
[**list_volume_snapshots**](StorageServiceApi.md#list_volume_snapshots) | **GET** /storage/v1/volumes/{volume_id}/snapshots | Returns the snapshots of a volume, or of all volumes with volume_id \&quot;-\&quot;.
[**list_volumes**](StorageServiceApi.md#list_volumes) | **GET** /storage/v1/volumes | Returns the volumes matching a filter; the own disks of compute resources are not volumes (see Compute.disks).
[**migrate_volume**](StorageServiceApi.md#migrate_volume) | **POST** /storage/v1/volumes/{volume_id}/migrate | Moves a volume to another storage class.
[**revert_volume_snapshot**](StorageServiceApi.md#revert_volume_snapshot) | **POST** /storage/v1/volumes/{volume_id}/snapshots/{snapshot_id}/revert | Reverts a volume that no running compute resource uses to a snapshot; not possible once the volume is deleted.
[**scale_volume**](StorageServiceApi.md#scale_volume) | **POST** /storage/v1/volumes/{volume_id}/scale | Grows a volume.
[**update_volume**](StorageServiceApi.md#update_volume) | **PATCH** /storage/v1/volumes/{volume_id} | Updates the name and metadata of a volume.


# **create_volume**
> Volume create_volume(create_volume_request)

Creates a volume.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.create_volume_request import CreateVolumeRequest
from kubevim_client.models.volume import Volume
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
    api_instance = kubevim_client.StorageServiceApi(api_client)
    create_volume_request = kubevim_client.CreateVolumeRequest() # CreateVolumeRequest | 

    try:
        # Creates a volume.
        api_response = api_instance.create_volume(create_volume_request)
        print("The response of StorageServiceApi->create_volume:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling StorageServiceApi->create_volume: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **create_volume_request** | [**CreateVolumeRequest**](CreateVolumeRequest.md)|  | 

### Return type

[**Volume**](Volume.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | A virtualised storage resource, e.g. a disk of a compute resource. |  * Location - URI of the created resource. <br>  |
**0** | Error (SOL 013 6). |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_volume_snapshot**
> VolumeSnapshot create_volume_snapshot(volume_id, create_volume_snapshot_request)

Creates a snapshot of a volume.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.create_volume_snapshot_request import CreateVolumeSnapshotRequest
from kubevim_client.models.volume_snapshot import VolumeSnapshot
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
    api_instance = kubevim_client.StorageServiceApi(api_client)
    volume_id = 'volume_id_example' # str | Volume to take the snapshot of.
    create_volume_snapshot_request = kubevim_client.CreateVolumeSnapshotRequest() # CreateVolumeSnapshotRequest | 

    try:
        # Creates a snapshot of a volume.
        api_response = api_instance.create_volume_snapshot(volume_id, create_volume_snapshot_request)
        print("The response of StorageServiceApi->create_volume_snapshot:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling StorageServiceApi->create_volume_snapshot: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volume_id** | **str**| Volume to take the snapshot of. | 
 **create_volume_snapshot_request** | [**CreateVolumeSnapshotRequest**](CreateVolumeSnapshotRequest.md)|  | 

### Return type

[**VolumeSnapshot**](VolumeSnapshot.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | A snapshot of a volume (Kubernetes VolumeSnapshot). |  * Location - URI of the created resource. <br>  |
**0** | Error (SOL 013 6). |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_volume**
> delete_volume(volume_id)

Deletes a volume; may fail with RESOURCE_IN_USE while the volume is attached. Snapshots are kept.

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
    api_instance = kubevim_client.StorageServiceApi(api_client)
    volume_id = 'volume_id_example' # str | Identifier of the volume.

    try:
        # Deletes a volume; may fail with RESOURCE_IN_USE while the volume is attached. Snapshots are kept.
        api_instance.delete_volume(volume_id)
    except Exception as e:
        print("Exception when calling StorageServiceApi->delete_volume: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volume_id** | **str**| Identifier of the volume. | 

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

# **delete_volume_snapshot**
> delete_volume_snapshot(volume_id, snapshot_id)

Deletes a snapshot of a volume.

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
    api_instance = kubevim_client.StorageServiceApi(api_client)
    volume_id = 'volume_id_example' # str | Volume the snapshot belongs to.
    snapshot_id = 'snapshot_id_example' # str | Identifier of the snapshot.

    try:
        # Deletes a snapshot of a volume.
        api_instance.delete_volume_snapshot(volume_id, snapshot_id)
    except Exception as e:
        print("Exception when calling StorageServiceApi->delete_volume_snapshot: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volume_id** | **str**| Volume the snapshot belongs to. | 
 **snapshot_id** | **str**| Identifier of the snapshot. | 

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

# **get_volume**
> Volume get_volume(volume_id)

Returns one volume.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.volume import Volume
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
    api_instance = kubevim_client.StorageServiceApi(api_client)
    volume_id = 'volume_id_example' # str | Identifier of the volume.

    try:
        # Returns one volume.
        api_response = api_instance.get_volume(volume_id)
        print("The response of StorageServiceApi->get_volume:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling StorageServiceApi->get_volume: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volume_id** | **str**| Identifier of the volume. | 

### Return type

[**Volume**](Volume.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A virtualised storage resource, e.g. a disk of a compute resource. |  -  |
**0** | Error (SOL 013 6). |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_volume_snapshot**
> VolumeSnapshot get_volume_snapshot(volume_id, snapshot_id)

Returns one snapshot of a volume, also after the volume is deleted.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.volume_snapshot import VolumeSnapshot
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
    api_instance = kubevim_client.StorageServiceApi(api_client)
    volume_id = 'volume_id_example' # str | Volume the snapshot belongs to.
    snapshot_id = 'snapshot_id_example' # str | Identifier of the snapshot.

    try:
        # Returns one snapshot of a volume, also after the volume is deleted.
        api_response = api_instance.get_volume_snapshot(volume_id, snapshot_id)
        print("The response of StorageServiceApi->get_volume_snapshot:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling StorageServiceApi->get_volume_snapshot: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volume_id** | **str**| Volume the snapshot belongs to. | 
 **snapshot_id** | **str**| Identifier of the snapshot. | 

### Return type

[**VolumeSnapshot**](VolumeSnapshot.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A snapshot of a volume (Kubernetes VolumeSnapshot). |  -  |
**0** | Error (SOL 013 6). |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_volume_snapshots**
> List[VolumeSnapshot] list_volume_snapshots(volume_id, filter=filter, nextpage_opaque_marker=nextpage_opaque_marker)

Returns the snapshots of a volume, or of all volumes with volume_id \"-\".

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.volume_snapshot import VolumeSnapshot
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
    api_instance = kubevim_client.StorageServiceApi(api_client)
    volume_id = 'volume_id_example' # str | Volume the snapshots belong to, or \"-\" for all volumes (AIP-159).
    filter = 'filter_example' # str | Attribute-based filter expression (SOL 013 5.2). (optional)
    nextpage_opaque_marker = 'nextpage_opaque_marker_example' # str | Marker of the page to return, from a previous response (SOL 013 5.4). (optional)

    try:
        # Returns the snapshots of a volume, or of all volumes with volume_id \"-\".
        api_response = api_instance.list_volume_snapshots(volume_id, filter=filter, nextpage_opaque_marker=nextpage_opaque_marker)
        print("The response of StorageServiceApi->list_volume_snapshots:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling StorageServiceApi->list_volume_snapshots: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volume_id** | **str**| Volume the snapshots belong to, or \&quot;-\&quot; for all volumes (AIP-159). | 
 **filter** | **str**| Attribute-based filter expression (SOL 013 5.2). | [optional] 
 **nextpage_opaque_marker** | **str**| Marker of the page to return, from a previous response (SOL 013 5.4). | [optional] 

### Return type

[**List[VolumeSnapshot]**](VolumeSnapshot.md)

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

# **list_volumes**
> List[Volume] list_volumes(filter=filter, nextpage_opaque_marker=nextpage_opaque_marker)

Returns the volumes matching a filter; the own disks of compute resources are not volumes (see Compute.disks).

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.volume import Volume
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
    api_instance = kubevim_client.StorageServiceApi(api_client)
    filter = 'filter_example' # str | Attribute-based filter expression (SOL 013 5.2). (optional)
    nextpage_opaque_marker = 'nextpage_opaque_marker_example' # str | Marker of the page to return, from a previous response (SOL 013 5.4). (optional)

    try:
        # Returns the volumes matching a filter; the own disks of compute resources are not volumes (see Compute.disks).
        api_response = api_instance.list_volumes(filter=filter, nextpage_opaque_marker=nextpage_opaque_marker)
        print("The response of StorageServiceApi->list_volumes:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling StorageServiceApi->list_volumes: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **str**| Attribute-based filter expression (SOL 013 5.2). | [optional] 
 **nextpage_opaque_marker** | **str**| Marker of the page to return, from a previous response (SOL 013 5.4). | [optional] 

### Return type

[**List[Volume]**](Volume.md)

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

# **migrate_volume**
> Volume migrate_volume(volume_id, migrate_volume_request)

Moves a volume to another storage class.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.migrate_volume_request import MigrateVolumeRequest
from kubevim_client.models.volume import Volume
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
    api_instance = kubevim_client.StorageServiceApi(api_client)
    volume_id = 'volume_id_example' # str | Identifier of the volume.
    migrate_volume_request = kubevim_client.MigrateVolumeRequest() # MigrateVolumeRequest | 

    try:
        # Moves a volume to another storage class.
        api_response = api_instance.migrate_volume(volume_id, migrate_volume_request)
        print("The response of StorageServiceApi->migrate_volume:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling StorageServiceApi->migrate_volume: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volume_id** | **str**| Identifier of the volume. | 
 **migrate_volume_request** | [**MigrateVolumeRequest**](MigrateVolumeRequest.md)|  | 

### Return type

[**Volume**](Volume.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A virtualised storage resource, e.g. a disk of a compute resource. |  -  |
**0** | Error (SOL 013 6). |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **revert_volume_snapshot**
> Volume revert_volume_snapshot(volume_id, snapshot_id, body)

Reverts a volume that no running compute resource uses to a snapshot; not possible once the volume is deleted.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.volume import Volume
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
    api_instance = kubevim_client.StorageServiceApi(api_client)
    volume_id = 'volume_id_example' # str | Volume to revert; no running compute resource may use it.
    snapshot_id = 'snapshot_id_example' # str | Identifier of the snapshot to revert to.
    body = None # object | 

    try:
        # Reverts a volume that no running compute resource uses to a snapshot; not possible once the volume is deleted.
        api_response = api_instance.revert_volume_snapshot(volume_id, snapshot_id, body)
        print("The response of StorageServiceApi->revert_volume_snapshot:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling StorageServiceApi->revert_volume_snapshot: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volume_id** | **str**| Volume to revert; no running compute resource may use it. | 
 **snapshot_id** | **str**| Identifier of the snapshot to revert to. | 
 **body** | **object**|  | 

### Return type

[**Volume**](Volume.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A virtualised storage resource, e.g. a disk of a compute resource. |  -  |
**0** | Error (SOL 013 6). |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **scale_volume**
> Volume scale_volume(volume_id, scale_volume_request)

Grows a volume.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.scale_volume_request import ScaleVolumeRequest
from kubevim_client.models.volume import Volume
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
    api_instance = kubevim_client.StorageServiceApi(api_client)
    volume_id = 'volume_id_example' # str | Identifier of the volume.
    scale_volume_request = kubevim_client.ScaleVolumeRequest() # ScaleVolumeRequest | 

    try:
        # Grows a volume.
        api_response = api_instance.scale_volume(volume_id, scale_volume_request)
        print("The response of StorageServiceApi->scale_volume:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling StorageServiceApi->scale_volume: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volume_id** | **str**| Identifier of the volume. | 
 **scale_volume_request** | [**ScaleVolumeRequest**](ScaleVolumeRequest.md)|  | 

### Return type

[**Volume**](Volume.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A virtualised storage resource, e.g. a disk of a compute resource. |  -  |
**0** | Error (SOL 013 6). |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_volume**
> Volume update_volume(volume_id, volume_update)

Updates the name and metadata of a volume.

### Example

* OAuth Authentication (oauth2):

```python
import kubevim_client
from kubevim_client.models.volume import Volume
from kubevim_client.models.volume_update import VolumeUpdate
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
    api_instance = kubevim_client.StorageServiceApi(api_client)
    volume_id = 'volume_id_example' # str | Identifier of the volume.
    volume_update = kubevim_client.VolumeUpdate() # VolumeUpdate | 

    try:
        # Updates the name and metadata of a volume.
        api_response = api_instance.update_volume(volume_id, volume_update)
        print("The response of StorageServiceApi->update_volume:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling StorageServiceApi->update_volume: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volume_id** | **str**| Identifier of the volume. | 
 **volume_update** | [**VolumeUpdate**](VolumeUpdate.md)|  | 

### Return type

[**Volume**](Volume.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A virtualised storage resource, e.g. a disk of a compute resource. |  -  |
**0** | Error (SOL 013 6). |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

