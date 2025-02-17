# coding: utf-8

"""
    vi-vnfm.proto

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

    The version of the OpenAPI document: version not set
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from kubevim_vivnfm_client.models.network_subnet import NetworkSubnet

class TestNetworkSubnet(unittest.TestCase):
    """NetworkSubnet unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> NetworkSubnet:
        """Test NetworkSubnet
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `NetworkSubnet`
        """
        model = NetworkSubnet()
        if include_optional:
            return NetworkSubnet(
                resource_id = kubevim_vivnfm_client.models.identifier.Identifier(
                    value = '', ),
                network_id = kubevim_vivnfm_client.models.identifier.Identifier(
                    value = '', ),
                ip_version = 'IPV4',
                gateway_ip = kubevim_vivnfm_client.models.ip_address.IPAddress(
                    ip = '', ),
                cidr = kubevim_vivnfm_client.models.ip_subnet_cidr.IPSubnetCIDR(
                    cidr = '', ),
                is_dhcp_enabled = True,
                address_pool = [
                    kubevim_vivnfm_client.models.todo:_might_be_few_ranges_specified_in_pool.TODO: Might be few ranges specified in pool(
                        start_ip = kubevim_vivnfm_client.models.ip_address.IPAddress(
                            ip = '', ), 
                        end_ip = kubevim_vivnfm_client.models.ip_address.IPAddress(
                            ip = '', ), )
                    ],
                metadata = kubevim_vivnfm_client.models.metadata.Metadata(
                    fields = {
                        'key' : ''
                        }, )
            )
        else:
            return NetworkSubnet(
        )
        """

    def testNetworkSubnet(self):
        """Test NetworkSubnet"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
