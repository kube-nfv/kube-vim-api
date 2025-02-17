# coding: utf-8

"""
    vi-vnfm.proto

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

    The version of the OpenAPI document: version not set
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from kubevim_vivnfm_client.models.pb_query_images_response import PbQueryImagesResponse

class TestPbQueryImagesResponse(unittest.TestCase):
    """PbQueryImagesResponse unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> PbQueryImagesResponse:
        """Test PbQueryImagesResponse
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `PbQueryImagesResponse`
        """
        model = PbQueryImagesResponse()
        if include_optional:
            return PbQueryImagesResponse(
                software_images_information = [
                    kubevim_vivnfm_client.models.this_information_element_represents_software_image_information.This information element represents Software Image Information(
                        software_image_id = kubevim_vivnfm_client.models.identifier.Identifier(
                            value = '', ), 
                        name = '', 
                        provider = '', 
                        version = '', 
                        checksum = '', 
                        container_format = '', 
                        disk_format = '', 
                        created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                        updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                        min_disk = kubevim_vivnfm_client.models.resource_quantity.resourceQuantity(
                            string = '', ), 
                        min_ram = kubevim_vivnfm_client.models.resource_quantity.resourceQuantity(
                            string = '', ), 
                        size = , 
                        status = '', 
                        metadata = kubevim_vivnfm_client.models.metadata.Metadata(
                            fields = {
                                'key' : ''
                                }, ), )
                    ]
            )
        else:
            return PbQueryImagesResponse(
        )
        """

    def testPbQueryImagesResponse(self):
        """Test PbQueryImagesResponse"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
