# coding: utf-8

"""
    vi-vnfm.proto

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

    The version of the OpenAPI document: version not set
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from kubevim_vivnfm_client.models.pb_create_compute_resource_affinity_or_anti_affinity_constraints_group_response import PbCreateComputeResourceAffinityOrAntiAffinityConstraintsGroupResponse

class TestPbCreateComputeResourceAffinityOrAntiAffinityConstraintsGroupResponse(unittest.TestCase):
    """PbCreateComputeResourceAffinityOrAntiAffinityConstraintsGroupResponse unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> PbCreateComputeResourceAffinityOrAntiAffinityConstraintsGroupResponse:
        """Test PbCreateComputeResourceAffinityOrAntiAffinityConstraintsGroupResponse
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `PbCreateComputeResourceAffinityOrAntiAffinityConstraintsGroupResponse`
        """
        model = PbCreateComputeResourceAffinityOrAntiAffinityConstraintsGroupResponse()
        if include_optional:
            return PbCreateComputeResourceAffinityOrAntiAffinityConstraintsGroupResponse(
                group_id = kubevim_vivnfm_client.models.identifier.Identifier(
                    value = '', )
            )
        else:
            return PbCreateComputeResourceAffinityOrAntiAffinityConstraintsGroupResponse(
        )
        """

    def testPbCreateComputeResourceAffinityOrAntiAffinityConstraintsGroupResponse(self):
        """Test PbCreateComputeResourceAffinityOrAntiAffinityConstraintsGroupResponse"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
