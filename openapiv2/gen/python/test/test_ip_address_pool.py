# coding: utf-8

"""
    vi-vnfm.proto

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

    The version of the OpenAPI document: version not set
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from vivnfm_client.models.ip_address_pool import IPAddressPool

class TestIPAddressPool(unittest.TestCase):
    """IPAddressPool unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> IPAddressPool:
        """Test IPAddressPool
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `IPAddressPool`
        """
        model = IPAddressPool()
        if include_optional:
            return IPAddressPool(
                start_ip = vivnfm_client.models.ip_address.IPAddress(
                    ip = '', ),
                end_ip = vivnfm_client.models.ip_address.IPAddress(
                    ip = '', )
            )
        else:
            return IPAddressPool(
        )
        """

    def testIPAddressPool(self):
        """Test IPAddressPool"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
