# coding: utf-8

"""
    vi-vnfm.proto

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

    The version of the OpenAPI document: version not set
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from vivnfm_client.models.pb_create_compute_flavour_request import PbCreateComputeFlavourRequest

class TestPbCreateComputeFlavourRequest(unittest.TestCase):
    """PbCreateComputeFlavourRequest unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> PbCreateComputeFlavourRequest:
        """Test PbCreateComputeFlavourRequest
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `PbCreateComputeFlavourRequest`
        """
        model = PbCreateComputeFlavourRequest()
        if include_optional:
            return PbCreateComputeFlavourRequest(
                flavour = vivnfm_client.models.virtual_compute_flavour.VirtualComputeFlavour(
                    flavour_id = vivnfm_client.models.identifier.Identifier(
                        value = '', ), 
                    is_public = True, 
                    virtual_memory = vivnfm_client.models.virtual_memory_data.VirtualMemoryData(
                        virtual_mem_size = 1.337, 
                        virtual_mem_oversubscription_policy = '', 
                        numa_enabled = True, ), 
                    virtual_cpu = vivnfm_client.models.virtual_cpu_data.VirtualCpuData(
                        cpu_architecture = '', 
                        num_virtual_cpu = 56, 
                        cpu_clock = 1.337, 
                        virtual_cpu_oversubscription_policy = '', 
                        virtual_cpu_pinning = vivnfm_client.models.virtual_cpu_data_virtual_cpu_pinning_data.VirtualCpuDataVirtualCpuPinningData(
                            virtual_cpu_pinning_policy = 'STATIC', 
                            virtual_cpu_pinning_rules = [
                                vivnfm_client.models.virtual_cpu_pinning_data_virtual_cpu_pinning_rule.VirtualCpuPinningDataVirtualCpuPinningRule(
                                    cores = 1.337, 
                                    sockets = 1.337, 
                                    threads = 1.337, )
                                ], ), 
                        power_state_reqs = '', ), 
                    storage_attributes = [
                        vivnfm_client.models.virtual_storage_data.VirtualStorageData(
                            type_of_storage = '', 
                            size_of_storage = 1.337, 
                            rdma_enabled = True, )
                        ], 
                    virtual_network_interface = [
                        vivnfm_client.models.virtual_network_interface_data.VirtualNetworkInterfaceData(
                            network_id = vivnfm_client.models.identifier.Identifier(
                                value = '', ), 
                            network_port_id = , 
                            bandwidth = 1.337, 
                            acceleration_capability_for_virtual_network_interface = [
                                ''
                                ], 
                            metadata = vivnfm_client.models.metadata.Metadata(
                                fields = {
                                    'key' : ''
                                    }, ), )
                        ], 
                    metadata = vivnfm_client.models.metadata.Metadata(), )
            )
        else:
            return PbCreateComputeFlavourRequest(
        )
        """

    def testPbCreateComputeFlavourRequest(self):
        """Test PbCreateComputeFlavourRequest"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
