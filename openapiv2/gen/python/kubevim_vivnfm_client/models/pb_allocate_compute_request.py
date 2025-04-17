# coding: utf-8

"""
    vi-vnfm.proto

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

    The version of the OpenAPI document: version not set
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json

from pydantic import BaseModel, ConfigDict, Field, StrictStr
from typing import Any, ClassVar, Dict, List, Optional
from kubevim_vivnfm_client.models.affinity_or_anti_affinity_constraint_for_compute import AffinityOrAntiAffinityConstraintForCompute
from kubevim_vivnfm_client.models.identifier import Identifier
from kubevim_vivnfm_client.models.metadata import Metadata
from kubevim_vivnfm_client.models.user_data import UserData
from kubevim_vivnfm_client.models.virtual_network_interface_data import VirtualNetworkInterfaceData
from kubevim_vivnfm_client.models.virtual_network_interface_ipam import VirtualNetworkInterfaceIPAM
from typing import Optional, Set
from typing_extensions import Self

class PbAllocateComputeRequest(BaseModel):
    """
    PbAllocateComputeRequest
    """ # noqa: E501
    compute_name: Optional[StrictStr] = Field(default=None, description="Name provided by the consumer for the virtualised compute resource to be allocated. It can be used for identifying resources from consumer side.", alias="computeName")
    reservation_id: Optional[Identifier] = Field(default=None, alias="reservationId")
    affinity_or_anti_affinity_constraints: Optional[List[AffinityOrAntiAffinityConstraintForCompute]] = Field(default=None, description="List of elements with affinity or anti affinity (see clause 8.4.8.2) information of the virtualised compute resource to be allocated. All the listed constraints shall be fulfilled for a successful operation.", alias="affinityOrAntiAffinityConstraints")
    compute_flavour_id: Identifier = Field(alias="computeFlavourId")
    vc_image_id: Optional[Identifier] = Field(default=None, alias="vcImageId")
    interface_data: Optional[List[VirtualNetworkInterfaceData]] = Field(default=None, description="Note: That is out of the ETSI GS NFV-IFA 006 scope. Traditionaly VirtualNetworkInterfaceData specified in the virtualComputeFlavour, but it is reduce flexibility, since the flavor contains virtual compute related networks, and network configuration for it (eg. QoS). Descided to move it in the AllocateComputeRequest.", alias="interfaceData")
    interface_ipam: Optional[List[VirtualNetworkInterfaceIPAM]] = Field(default=None, description="IPAM Data of network interfaces which are specific to a Virtual Compute Resource instance. See clause 8.4.3.7.", alias="interfaceIPAM")
    meta_data: Optional[Metadata] = Field(default=None, alias="metaData")
    resource_group_id: Optional[Identifier] = Field(default=None, alias="resourceGroupId")
    user_data: Optional[UserData] = Field(default=None, alias="userData")
    __properties: ClassVar[List[str]] = ["computeName", "reservationId", "affinityOrAntiAffinityConstraints", "computeFlavourId", "vcImageId", "interfaceData", "interfaceIPAM", "metaData", "resourceGroupId", "userData"]

    model_config = ConfigDict(
        populate_by_name=True,
        validate_assignment=True,
        protected_namespaces=(),
    )


    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.model_dump(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        # TODO: pydantic v2: use .model_dump_json(by_alias=True, exclude_unset=True) instead
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> Optional[Self]:
        """Create an instance of PbAllocateComputeRequest from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        """
        excluded_fields: Set[str] = set([
        ])

        _dict = self.model_dump(
            by_alias=True,
            exclude=excluded_fields,
            exclude_none=True,
        )
        # override the default output from pydantic by calling `to_dict()` of reservation_id
        if self.reservation_id:
            _dict['reservationId'] = self.reservation_id.to_dict()
        # override the default output from pydantic by calling `to_dict()` of each item in affinity_or_anti_affinity_constraints (list)
        _items = []
        if self.affinity_or_anti_affinity_constraints:
            for _item_affinity_or_anti_affinity_constraints in self.affinity_or_anti_affinity_constraints:
                if _item_affinity_or_anti_affinity_constraints:
                    _items.append(_item_affinity_or_anti_affinity_constraints.to_dict())
            _dict['affinityOrAntiAffinityConstraints'] = _items
        # override the default output from pydantic by calling `to_dict()` of compute_flavour_id
        if self.compute_flavour_id:
            _dict['computeFlavourId'] = self.compute_flavour_id.to_dict()
        # override the default output from pydantic by calling `to_dict()` of vc_image_id
        if self.vc_image_id:
            _dict['vcImageId'] = self.vc_image_id.to_dict()
        # override the default output from pydantic by calling `to_dict()` of each item in interface_data (list)
        _items = []
        if self.interface_data:
            for _item_interface_data in self.interface_data:
                if _item_interface_data:
                    _items.append(_item_interface_data.to_dict())
            _dict['interfaceData'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in interface_ipam (list)
        _items = []
        if self.interface_ipam:
            for _item_interface_ipam in self.interface_ipam:
                if _item_interface_ipam:
                    _items.append(_item_interface_ipam.to_dict())
            _dict['interfaceIPAM'] = _items
        # override the default output from pydantic by calling `to_dict()` of meta_data
        if self.meta_data:
            _dict['metaData'] = self.meta_data.to_dict()
        # override the default output from pydantic by calling `to_dict()` of resource_group_id
        if self.resource_group_id:
            _dict['resourceGroupId'] = self.resource_group_id.to_dict()
        # override the default output from pydantic by calling `to_dict()` of user_data
        if self.user_data:
            _dict['userData'] = self.user_data.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of PbAllocateComputeRequest from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "computeName": obj.get("computeName"),
            "reservationId": Identifier.from_dict(obj["reservationId"]) if obj.get("reservationId") is not None else None,
            "affinityOrAntiAffinityConstraints": [AffinityOrAntiAffinityConstraintForCompute.from_dict(_item) for _item in obj["affinityOrAntiAffinityConstraints"]] if obj.get("affinityOrAntiAffinityConstraints") is not None else None,
            "computeFlavourId": Identifier.from_dict(obj["computeFlavourId"]) if obj.get("computeFlavourId") is not None else None,
            "vcImageId": Identifier.from_dict(obj["vcImageId"]) if obj.get("vcImageId") is not None else None,
            "interfaceData": [VirtualNetworkInterfaceData.from_dict(_item) for _item in obj["interfaceData"]] if obj.get("interfaceData") is not None else None,
            "interfaceIPAM": [VirtualNetworkInterfaceIPAM.from_dict(_item) for _item in obj["interfaceIPAM"]] if obj.get("interfaceIPAM") is not None else None,
            "metaData": Metadata.from_dict(obj["metaData"]) if obj.get("metaData") is not None else None,
            "resourceGroupId": Identifier.from_dict(obj["resourceGroupId"]) if obj.get("resourceGroupId") is not None else None,
            "userData": UserData.from_dict(obj["userData"]) if obj.get("userData") is not None else None
        })
        return _obj


