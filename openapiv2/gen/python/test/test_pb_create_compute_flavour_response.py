# coding: utf-8

"""
    vi-vnfm.proto

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

    The version of the OpenAPI document: version not set
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from kubevim_vivnfm_client.models.pb_create_compute_flavour_response import PbCreateComputeFlavourResponse

class TestPbCreateComputeFlavourResponse(unittest.TestCase):
    """PbCreateComputeFlavourResponse unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> PbCreateComputeFlavourResponse:
        """Test PbCreateComputeFlavourResponse
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `PbCreateComputeFlavourResponse`
        """
        model = PbCreateComputeFlavourResponse()
        if include_optional:
            return PbCreateComputeFlavourResponse(
                flavour_id = kubevim_vivnfm_client.models.identifier.Identifier(
                    value = '', )
            )
        else:
            return PbCreateComputeFlavourResponse(
        )
        """

    def testPbCreateComputeFlavourResponse(self):
        """Test PbCreateComputeFlavourResponse"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
