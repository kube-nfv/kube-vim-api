# coding: utf-8

"""
    vi-vnfm.proto

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

    The version of the OpenAPI document: version not set
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from vivnfm_client.models.network_qo_s import NetworkQoS

class TestNetworkQoS(unittest.TestCase):
    """NetworkQoS unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> NetworkQoS:
        """Test NetworkQoS
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `NetworkQoS`
        """
        model = NetworkQoS()
        if include_optional:
            return NetworkQoS(
                qos_name = '',
                qos_value = ''
            )
        else:
            return NetworkQoS(
        )
        """

    def testNetworkQoS(self):
        """Test NetworkQoS"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
