# coding: utf-8

"""
    vi-vnfm.proto

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

    The version of the OpenAPI document: version not set
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from kubevim_vivnfm_client.models.virtual_cpu_data import VirtualCpuData

class TestVirtualCpuData(unittest.TestCase):
    """VirtualCpuData unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> VirtualCpuData:
        """Test VirtualCpuData
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `VirtualCpuData`
        """
        model = VirtualCpuData()
        if include_optional:
            return VirtualCpuData(
                cpu_architecture = '',
                num_virtual_cpu = 56,
                cpu_clock = 1.337,
                virtual_cpu_oversubscription_policy = '',
                virtual_cpu_pinning = kubevim_vivnfm_client.models.virtual_cpu_data_virtual_cpu_pinning_data.VirtualCpuDataVirtualCpuPinningData(
                    virtual_cpu_pinning_policy = 'STATIC', 
                    virtual_cpu_pinning_rules = [
                        kubevim_vivnfm_client.models.virtual_cpu_pinning_data_virtual_cpu_pinning_rule.VirtualCpuPinningDataVirtualCpuPinningRule(
                            cores = 1.337, 
                            sockets = 1.337, 
                            threads = 1.337, )
                        ], ),
                power_state_reqs = ''
            )
        else:
            return VirtualCpuData(
        )
        """

    def testVirtualCpuData(self):
        """Test VirtualCpuData"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
