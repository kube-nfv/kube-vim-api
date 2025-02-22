# coding: utf-8

# flake8: noqa
"""
    vi-vnfm.proto

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

    The version of the OpenAPI document: version not set
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


# import models into model package
from kubevim_vivnfm_client.models.affinity_or_anti_affinity_constraint_for_compute import AffinityOrAntiAffinityConstraintForCompute
from kubevim_vivnfm_client.models.affinity_or_anti_affinity_constraint_for_compute_affinity_or_anti_affinity_resource_list import AffinityOrAntiAffinityConstraintForComputeAffinityOrAntiAffinityResourceList
from kubevim_vivnfm_client.models.compute_running_state import ComputeRunningState
from kubevim_vivnfm_client.models.filter import Filter
from kubevim_vivnfm_client.models.ip_address import IPAddress
from kubevim_vivnfm_client.models.ip_address_pool import IPAddressPool
from kubevim_vivnfm_client.models.ip_subnet_cidr import IPSubnetCIDR
from kubevim_vivnfm_client.models.ip_version import IPVersion
from kubevim_vivnfm_client.models.identifier import Identifier
from kubevim_vivnfm_client.models.mac_address import MacAddress
from kubevim_vivnfm_client.models.metadata import Metadata
from kubevim_vivnfm_client.models.network_qo_s import NetworkQoS
from kubevim_vivnfm_client.models.network_resource_type import NetworkResourceType
from kubevim_vivnfm_client.models.network_subnet import NetworkSubnet
from kubevim_vivnfm_client.models.network_subnet_data import NetworkSubnetData
from kubevim_vivnfm_client.models.operational_state import OperationalState
from kubevim_vivnfm_client.models.pb_allocate_compute_request import PbAllocateComputeRequest
from kubevim_vivnfm_client.models.pb_allocate_compute_response import PbAllocateComputeResponse
from kubevim_vivnfm_client.models.pb_allocate_network_request import PbAllocateNetworkRequest
from kubevim_vivnfm_client.models.pb_allocate_network_response import PbAllocateNetworkResponse
from kubevim_vivnfm_client.models.pb_create_compute_flavour_request import PbCreateComputeFlavourRequest
from kubevim_vivnfm_client.models.pb_create_compute_flavour_response import PbCreateComputeFlavourResponse
from kubevim_vivnfm_client.models.pb_create_compute_resource_affinity_or_anti_affinity_constraints_group_request import PbCreateComputeResourceAffinityOrAntiAffinityConstraintsGroupRequest
from kubevim_vivnfm_client.models.pb_create_compute_resource_affinity_or_anti_affinity_constraints_group_response import PbCreateComputeResourceAffinityOrAntiAffinityConstraintsGroupResponse
from kubevim_vivnfm_client.models.pb_query_compute_flavour_response import PbQueryComputeFlavourResponse
from kubevim_vivnfm_client.models.pb_query_compute_response import PbQueryComputeResponse
from kubevim_vivnfm_client.models.pb_query_image_request import PbQueryImageRequest
from kubevim_vivnfm_client.models.pb_query_image_response import PbQueryImageResponse
from kubevim_vivnfm_client.models.pb_query_images_response import PbQueryImagesResponse
from kubevim_vivnfm_client.models.pb_query_network_response import PbQueryNetworkResponse
from kubevim_vivnfm_client.models.pb_terminate_network_response import PbTerminateNetworkResponse
from kubevim_vivnfm_client.models.protobuf_any import ProtobufAny
from kubevim_vivnfm_client.models.resource_quantity import ResourceQuantity
from kubevim_vivnfm_client.models.rpc_status import RpcStatus
from kubevim_vivnfm_client.models.scope_of_affinity_or_anti_affinity_constraint_for_compute import ScopeOfAffinityOrAntiAffinityConstraintForCompute
from kubevim_vivnfm_client.models.software_image_information import SoftwareImageInformation
from kubevim_vivnfm_client.models.type_of_affinity_or_anti_affinity_constraint import TypeOfAffinityOrAntiAffinityConstraint
from kubevim_vivnfm_client.models.user_data import UserData
from kubevim_vivnfm_client.models.user_data_user_data_transportation_method import UserDataUserDataTransportationMethod
from kubevim_vivnfm_client.models.virtual_compute import VirtualCompute
from kubevim_vivnfm_client.models.virtual_compute_flavour import VirtualComputeFlavour
from kubevim_vivnfm_client.models.virtual_cpu_data import VirtualCpuData
from kubevim_vivnfm_client.models.virtual_cpu_data_virtual_cpu_pinning_data import VirtualCpuDataVirtualCpuPinningData
from kubevim_vivnfm_client.models.virtual_cpu_pinning_data_virtual_cpu_pinning_policy import VirtualCpuPinningDataVirtualCpuPinningPolicy
from kubevim_vivnfm_client.models.virtual_cpu_pinning_data_virtual_cpu_pinning_rule import VirtualCpuPinningDataVirtualCpuPinningRule
from kubevim_vivnfm_client.models.virtual_interface_data import VirtualInterfaceData
from kubevim_vivnfm_client.models.virtual_memory_data import VirtualMemoryData
from kubevim_vivnfm_client.models.virtual_network import VirtualNetwork
from kubevim_vivnfm_client.models.virtual_network_data import VirtualNetworkData
from kubevim_vivnfm_client.models.virtual_network_interface_data import VirtualNetworkInterfaceData
from kubevim_vivnfm_client.models.virtual_storage_data import VirtualStorageData
