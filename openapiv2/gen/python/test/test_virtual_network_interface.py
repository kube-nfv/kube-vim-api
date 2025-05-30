# coding: utf-8

"""
    vi-vnfm.proto

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

    The version of the OpenAPI document: version not set
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from kubevim_vivnfm_client.models.virtual_network_interface import VirtualNetworkInterface

class TestVirtualNetworkInterface(unittest.TestCase):
    """VirtualNetworkInterface unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> VirtualNetworkInterface:
        """Test VirtualNetworkInterface
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `VirtualNetworkInterface`
        """
        model = VirtualNetworkInterface()
        if include_optional:
            return VirtualNetworkInterface(
                resource_id = kubevim_vivnfm_client.models.identifier.Identifier(
                    value = '', ),
                network_id = kubevim_vivnfm_client.models.identifier.Identifier(
                    value = '', ),
                subnet_id = kubevim_vivnfm_client.models.identifier.Identifier(
                    value = '', ),
                network_port_id = kubevim_vivnfm_client.models.identifier.Identifier(
                    value = '', ),
                ip_address = [
                    kubevim_vivnfm_client.models.ip_address.IPAddress(
                        ip = '', )
                    ],
                type_virtual_nic = 'BRIDGE',
                type_configuration = [
                    ''
                    ],
                mac_address = kubevim_vivnfm_client.models.mac_address.MacAddress(
                    mac = '', ),
                bandwidth = 1.337,
                acceleration_capability = [
                    ''
                    ],
                operational_state = 'ENABLED',
                metadata = kubevim_vivnfm_client.models.metadata.Metadata(
                    fields = {
                        'key' : ''
                        }, )
            )
        else:
            return VirtualNetworkInterface(
                resource_id = kubevim_vivnfm_client.models.identifier.Identifier(
                    value = '', ),
                type_virtual_nic = 'BRIDGE',
                mac_address = kubevim_vivnfm_client.models.mac_address.MacAddress(
                    mac = '', ),
                bandwidth = 1.337,
                operational_state = 'ENABLED',
        )
        """

    def testVirtualNetworkInterface(self):
        """Test VirtualNetworkInterface"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
