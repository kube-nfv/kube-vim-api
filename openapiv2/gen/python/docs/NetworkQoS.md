# NetworkQoS

This clause describes the attributes for the NetworkQoS information element. This type gives QoS options to be supported on the virtualised network, e.g. latency, jitter, etc.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**qos_name** | **str** | Name given to the QoS parameter. | [optional] 
**qos_value** | **str** | Value of the QoS parameter. | [optional] 

## Example

```python
from kubevim_vivnfm_client.models.network_qo_s import NetworkQoS

# TODO update the JSON string below
json = "{}"
# create an instance of NetworkQoS from a JSON string
network_qo_s_instance = NetworkQoS.from_json(json)
# print the JSON string representation of the object
print(NetworkQoS.to_json())

# convert the object into a dict
network_qo_s_dict = network_qo_s_instance.to_dict()
# create an instance of NetworkQoS from a dict
network_qo_s_from_dict = NetworkQoS.from_dict(network_qo_s_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


