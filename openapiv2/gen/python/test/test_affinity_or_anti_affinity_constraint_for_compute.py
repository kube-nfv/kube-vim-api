# coding: utf-8

"""
    vi-vnfm.proto

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

    The version of the OpenAPI document: version not set
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from vivnfm_client.models.affinity_or_anti_affinity_constraint_for_compute import AffinityOrAntiAffinityConstraintForCompute

class TestAffinityOrAntiAffinityConstraintForCompute(unittest.TestCase):
    """AffinityOrAntiAffinityConstraintForCompute unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> AffinityOrAntiAffinityConstraintForCompute:
        """Test AffinityOrAntiAffinityConstraintForCompute
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `AffinityOrAntiAffinityConstraintForCompute`
        """
        model = AffinityOrAntiAffinityConstraintForCompute()
        if include_optional:
            return AffinityOrAntiAffinityConstraintForCompute(
                type = 'AFFINITY',
                scope = 'NFVI_NODE',
                affinity_or_anti_affinity_resource_list = vivnfm_client.models.affinity_or_anti_affinity_constraint_for_compute_affinity_or_anti_affinity_resource_list.AffinityOrAntiAffinityConstraintForComputeAffinityOrAntiAffinityResourceList(
                    resource_id = [
                        vivnfm_client.models.identifier.Identifier(
                            value = '', )
                        ], ),
                affinity_or_anti_affinity_resource_group_id = vivnfm_client.models.identifier.Identifier(
                    value = '', )
            )
        else:
            return AffinityOrAntiAffinityConstraintForCompute(
        )
        """

    def testAffinityOrAntiAffinityConstraintForCompute(self):
        """Test AffinityOrAntiAffinityConstraintForCompute"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
