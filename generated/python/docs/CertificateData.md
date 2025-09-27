# CertificateData

Note: Either set of \"privatekey\" and \"certificateFile\" or \"keystoreFile\" but not both shall be present. certificateFile is optional only if keystoreFile is present. certSubjectData is required only when a certificate needs to be generated or signed (i.e., not when a complete keystore is provided). certificateProfileName helps automate signing via predefined templates or CA policies.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**private_key** | **List[str]** | Private key paired with signed public key. VNFM shall generate both private key and public key and set this attribute. | [optional] 
**certificate_file** | **List[str]** | Signed certificate including public key and certificate chain. | [optional] 
**keystore_file** | **List[str]** | Keystore which includes the private key, signed certificate and certificate chain, e.g. pkcs#12, pfx. Credentials to read this file shall be provided to the VNF instance by outbound. | [optional] 
**cert_subject_data** | **List[str]** | Subject to be signed. | [optional] 
**certifiate_profile_name** | **List[str]** | Name of certificate profile to be signed. | [optional] 

## Example

```python
from kubevim_vivnfm_client.models.certificate_data import CertificateData

# TODO update the JSON string below
json = "{}"
# create an instance of CertificateData from a JSON string
certificate_data_instance = CertificateData.from_json(json)
# print the JSON string representation of the object
print(CertificateData.to_json())

# convert the object into a dict
certificate_data_dict = certificate_data_instance.to_dict()
# create an instance of CertificateData from a dict
certificate_data_from_dict = CertificateData.from_dict(certificate_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


