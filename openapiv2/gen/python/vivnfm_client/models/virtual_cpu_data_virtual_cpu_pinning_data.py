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

from pydantic import BaseModel, ConfigDict, Field
from typing import Any, ClassVar, Dict, List, Optional
from vivnfm_client.models.virtual_cpu_pinning_data_virtual_cpu_pinning_policy import VirtualCpuPinningDataVirtualCpuPinningPolicy
from vivnfm_client.models.virtual_cpu_pinning_data_virtual_cpu_pinning_rule import VirtualCpuPinningDataVirtualCpuPinningRule
from typing import Optional, Set
from typing_extensions import Self

class VirtualCpuDataVirtualCpuPinningData(BaseModel):
    """
    Information describing CPU pinning policy and rules for virtual CPU to physical CPU mapping of the virtualised compute resource.
    """ # noqa: E501
    virtual_cpu_pinning_policy: Optional[VirtualCpuPinningDataVirtualCpuPinningPolicy] = Field(default=VirtualCpuPinningDataVirtualCpuPinningPolicy.STATIC, alias="virtualCpuPinningPolicy")
    virtual_cpu_pinning_rules: Optional[List[VirtualCpuPinningDataVirtualCpuPinningRule]] = Field(default=None, description="A list of rules that should be considered during the allocation of the virtual CPU-s to logical CPU-s in case of \"static\" virtualCpuPinningPolicy.", alias="virtualCpuPinningRules")
    __properties: ClassVar[List[str]] = ["virtualCpuPinningPolicy", "virtualCpuPinningRules"]

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
        """Create an instance of VirtualCpuDataVirtualCpuPinningData from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in virtual_cpu_pinning_rules (list)
        _items = []
        if self.virtual_cpu_pinning_rules:
            for _item_virtual_cpu_pinning_rules in self.virtual_cpu_pinning_rules:
                if _item_virtual_cpu_pinning_rules:
                    _items.append(_item_virtual_cpu_pinning_rules.to_dict())
            _dict['virtualCpuPinningRules'] = _items
        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of VirtualCpuDataVirtualCpuPinningData from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "virtualCpuPinningPolicy": obj.get("virtualCpuPinningPolicy") if obj.get("virtualCpuPinningPolicy") is not None else VirtualCpuPinningDataVirtualCpuPinningPolicy.STATIC,
            "virtualCpuPinningRules": [VirtualCpuPinningDataVirtualCpuPinningRule.from_dict(_item) for _item in obj["virtualCpuPinningRules"]] if obj.get("virtualCpuPinningRules") is not None else None
        })
        return _obj


