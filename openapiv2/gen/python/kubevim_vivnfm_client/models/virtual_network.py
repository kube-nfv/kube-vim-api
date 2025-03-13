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

from pydantic import BaseModel, ConfigDict, Field, StrictBool, StrictFloat, StrictInt, StrictStr
from typing import Any, ClassVar, Dict, List, Optional, Union
from kubevim_vivnfm_client.models.identifier import Identifier
from kubevim_vivnfm_client.models.metadata import Metadata
from kubevim_vivnfm_client.models.network_qo_s import NetworkQoS
from kubevim_vivnfm_client.models.network_type import NetworkType
from kubevim_vivnfm_client.models.operational_state import OperationalState
from typing import Optional, Set
from typing_extensions import Self

class VirtualNetwork(BaseModel):
    """
    VirtualNetwork
    """ # noqa: E501
    network_resource_id: Optional[Identifier] = Field(default=None, alias="networkResourceId")
    network_resource_name: Optional[StrictStr] = Field(default=None, description="Name of the virtualised network resource.", alias="networkResourceName")
    subnet_id: Optional[List[Identifier]] = Field(default=None, description="References the network subnet. Only present if the network provides layer 3 connectivity.", alias="subnetId")
    network_port: Optional[List[Dict[str, Any]]] = Field(default=None, description="Provides information on an instantiated virtual network port.", alias="networkPort")
    bandwidth: Optional[Union[StrictFloat, StrictInt]] = None
    network_type: Optional[NetworkType] = Field(default=NetworkType.OVERLAY, alias="networkType")
    provider_network: Optional[StrictStr] = Field(default=None, description="Name of the infrastructure provider network used to realize the virtual network. Cardinality can be \"0\" to cover the case where virtual network is not based on infrastructure provider network.", alias="providerNetwork")
    segmentation_id: Optional[StrictStr] = Field(default=None, description="The segmentation identifier of the network that maps to the virtualised network, for which, the segmentation model is defined by the networkType attribute. For instance, for a \"vlan\" networkType, it corresponds to the vlan identifier; and for a \"gre\" networkType, it corresponds to a gre key. Cardinality can be \"0\" to cover the case where networkType is flat network without any specific segmentation.", alias="segmentationId")
    network_qo_s: Optional[List[NetworkQoS]] = Field(default=None, description="Provides information about Quality of Service attributes that the network supports. Cardinality can be \"0\" for virtual network without any QoS requirements.", alias="networkQoS")
    is_shared: Optional[StrictBool] = Field(default=None, description="Defines whether the virtualised network is shared among consumers.", alias="isShared")
    zone_id: Optional[Identifier] = Field(default=None, alias="zoneId")
    operational_state: Optional[OperationalState] = Field(default=OperationalState.ENABLED, alias="operationalState")
    metadata: Optional[Metadata] = None
    connected_networks: Optional[List[Identifier]] = Field(default=None, description="Specifies the virtual network resources to which the newly created virtual network is intended to be explicitly interconnected.", alias="connectedNetworks")
    __properties: ClassVar[List[str]] = ["networkResourceId", "networkResourceName", "subnetId", "networkPort", "bandwidth", "networkType", "providerNetwork", "segmentationId", "networkQoS", "isShared", "zoneId", "operationalState", "metadata", "connectedNetworks"]

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
        """Create an instance of VirtualNetwork from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of network_resource_id
        if self.network_resource_id:
            _dict['networkResourceId'] = self.network_resource_id.to_dict()
        # override the default output from pydantic by calling `to_dict()` of each item in subnet_id (list)
        _items = []
        if self.subnet_id:
            for _item_subnet_id in self.subnet_id:
                if _item_subnet_id:
                    _items.append(_item_subnet_id.to_dict())
            _dict['subnetId'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in network_qo_s (list)
        _items = []
        if self.network_qo_s:
            for _item_network_qo_s in self.network_qo_s:
                if _item_network_qo_s:
                    _items.append(_item_network_qo_s.to_dict())
            _dict['networkQoS'] = _items
        # override the default output from pydantic by calling `to_dict()` of zone_id
        if self.zone_id:
            _dict['zoneId'] = self.zone_id.to_dict()
        # override the default output from pydantic by calling `to_dict()` of metadata
        if self.metadata:
            _dict['metadata'] = self.metadata.to_dict()
        # override the default output from pydantic by calling `to_dict()` of each item in connected_networks (list)
        _items = []
        if self.connected_networks:
            for _item_connected_networks in self.connected_networks:
                if _item_connected_networks:
                    _items.append(_item_connected_networks.to_dict())
            _dict['connectedNetworks'] = _items
        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of VirtualNetwork from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "networkResourceId": Identifier.from_dict(obj["networkResourceId"]) if obj.get("networkResourceId") is not None else None,
            "networkResourceName": obj.get("networkResourceName"),
            "subnetId": [Identifier.from_dict(_item) for _item in obj["subnetId"]] if obj.get("subnetId") is not None else None,
            "networkPort": obj.get("networkPort"),
            "bandwidth": obj.get("bandwidth"),
            "networkType": obj.get("networkType") if obj.get("networkType") is not None else NetworkType.OVERLAY,
            "providerNetwork": obj.get("providerNetwork"),
            "segmentationId": obj.get("segmentationId"),
            "networkQoS": [NetworkQoS.from_dict(_item) for _item in obj["networkQoS"]] if obj.get("networkQoS") is not None else None,
            "isShared": obj.get("isShared"),
            "zoneId": Identifier.from_dict(obj["zoneId"]) if obj.get("zoneId") is not None else None,
            "operationalState": obj.get("operationalState") if obj.get("operationalState") is not None else OperationalState.ENABLED,
            "metadata": Metadata.from_dict(obj["metadata"]) if obj.get("metadata") is not None else None,
            "connectedNetworks": [Identifier.from_dict(_item) for _item in obj["connectedNetworks"]] if obj.get("connectedNetworks") is not None else None
        })
        return _obj


