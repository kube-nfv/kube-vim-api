# UserData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**content** | **str** | Contains the user data to customize a virtualised compute resource at boot-time. | 
**method** | [**UserDataUserDataTransportationMethod**](UserDataUserDataTransportationMethod.md) |  | [optional] [default to UserDataUserDataTransportationMethod.CONFIG_DRIVE_PLAINTEXT]
**certificate_data** | [**List[CertificateData]**](CertificateData.md) | Contains the additional user data to store certificate data for the VNF composed of (fully or partially) virtualised compute resource at boot-time. Shall be present if delegation-mode is used. Otherwise it shall be absent. | [optional] 

## Example

```python
from kubevim_vivnfm_client.models.user_data import UserData

# TODO update the JSON string below
json = "{}"
# create an instance of UserData from a JSON string
user_data_instance = UserData.from_json(json)
# print the JSON string representation of the object
print(UserData.to_json())

# convert the object into a dict
user_data_dict = user_data_instance.to_dict()
# create an instance of UserData from a dict
user_data_from_dict = UserData.from_dict(user_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


