# coding: utf-8

"""
    vi-vnfm.proto

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

    The version of the OpenAPI document: version not set
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from kubevim_vivnfm_client.api.vi_vnfm_api import ViVnfmApi


class TestViVnfmApi(unittest.TestCase):
    """ViVnfmApi unit test stubs"""

    def setUp(self) -> None:
        self.api = ViVnfmApi()

    def tearDown(self) -> None:
        pass

    def test_vi_vnfm_allocate_virtualised_compute_resource(self) -> None:
        """Test case for vi_vnfm_allocate_virtualised_compute_resource

        This operation allows requesting the allocation of virtualised compute resources as indicated by the consumer functional block. Result: After successful operation, the VIM has created the internal management objects for the virtualised compute resource and allocated this resource according to the input requirements and constraints. In addition, the VIM shall return to the VNFM information on the newly instantiated virtualised compute resource plus any additional information about the allocate request operation. The VIM may also return intermediate status reports during the allocation process. If the operation was not successful, the VIM shall return to the VNFM appropriate error information.
        """
        pass

    def test_vi_vnfm_allocate_virtualised_network_resource(self) -> None:
        """Test case for vi_vnfm_allocate_virtualised_network_resource

        This operation allows requesting the allocation of virtualised network resources as indicated by the consumer functional block. Result: After successful operation, the VIM has created the internal management objects for the virtualised network resource and allocated this resource. In addition, the VIM shall return to the VNFM information on the newly instantiated virtualised network resource plus any additional information about the allocate request operation. The VIM may also return intermediate status reports during the allocation process. If the operation was not successful, the VIM shall return to the VNFM appropriate error information.
        """
        pass

    def test_vi_vnfm_create_compute_flavour(self) -> None:
        """Test case for vi_vnfm_create_compute_flavour

        This operation allows requesting the creation of a flavour as indicated by the consumer functional block. Result: After successful operation, the VIM has created the Compute Flavour. In addition, the VIM shall return to the VNFM information on the newly created Compute Flavour. If the operation was not successful, the VIM shall return to the VNFM appropriate error information.
        """
        pass

    def test_vi_vnfm_delete_compute_flavour(self) -> None:
        """Test case for vi_vnfm_delete_compute_flavour

        This operation allows deleting a Compute Flavour. Result: After successful operation, the VIM has deleted the Compute Flavour, so no new Virtualised Compute Resource can be allocated based on it. The already allocated Virtualised Compute Resources are not affected. If the operation was not successful, the VIM shall return to the VNFM appropriate error information.
        """
        pass

    def test_vi_vnfm_query_compute_flavour(self) -> None:
        """Test case for vi_vnfm_query_compute_flavour

        This operation allows querying information about created Compute Flavours. Result: After successful operation, the VIM has queried the internal management objects for the Compute Flavours. The result of the query shall indicate with a standard success/error result if the query has been processed correctly. For a particular query, information about the Compute Flavours that the VNFM has access to and that are matching the filter shall be returned.
        """
        pass

    def test_vi_vnfm_query_image(self) -> None:
        """Test case for vi_vnfm_query_image

        This operation allows querying the information about a specific software image in the image repository managed by the VIM. Result: As a result of this operation, the producer (VIM) shall indicate to the consumer (VNFM) whether or not it was possible to process the query.
        """
        pass

    def test_vi_vnfm_query_image2(self) -> None:
        """Test case for vi_vnfm_query_image2

        This operation allows querying the information about a specific software image in the image repository managed by the VIM. Result: As a result of this operation, the producer (VIM) shall indicate to the consumer (VNFM) whether or not it was possible to process the query.
        """
        pass

    def test_vi_vnfm_query_images(self) -> None:
        """Test case for vi_vnfm_query_images

        This operation allows querying the information of software images in the image repository managed by the VIM. Result: As a result of this operation, the producer (VIM) shall indicate to the consumer (VNFM) whether or not it was possible to process the query
        """
        pass

    def test_vi_vnfm_query_virtualised_network_resource(self) -> None:
        """Test case for vi_vnfm_query_virtualised_network_resource

        This operation allows querying information about instantiated virtualised network resources. Result: After successful operation, the VIM has queried the internal management objects for the virtualised network resources. The result of the query shall indicate with a standard success/error result if the query has been processed correctly. For a particular query, information about the network resources that the VNFM has access to and that are matching the filter shall be returned.
        """
        pass

    def test_vi_vnfm_terminate_virtualised_network_resource(self) -> None:
        """Test case for vi_vnfm_terminate_virtualised_network_resource

        This operation allows de-allocating and terminating one or more an instantiated virtualised network resource(s). When the operation is done on multiple ids, it is assumed to be best-effort, i.e. it can succeed for a subset of the ids, and fail for the remaining ones. Result: After successful operation, the VIM has terminated the virtualised network resources and removed the internal management objects for those resources. In addition, the VIM shall return to the VNFM information on the terminated virtualised network resource plus any additional information about the terminate request operation. If the operation was not successful, the VIM shall return to the VNFM appropriate error information.
        """
        pass


if __name__ == '__main__':
    unittest.main()
