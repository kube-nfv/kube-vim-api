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
from vivnfm_client.models.identifier import Identifier
from typing import Optional, Set
from typing_extensions import Self

class AffinityOrAntiAffinityConstraintForComputeAffinityOrAntiAffinityResourceList(BaseModel):
    """
    The AffinityOrAntiAffinityResourceList information element defines an explicit list of resources to express affinity or anti-affinity between these resources and a current resource. The scope of the affinity or anti-affinity can also be defined.
    """ # noqa: E501
    resource_id: Optional[List[Identifier]] = Field(default=None, description="List of identifiers of virtualised resources.", alias="resourceId")
    __properties: ClassVar[List[str]] = ["resourceId"]

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
        """Create an instance of AffinityOrAntiAffinityConstraintForComputeAffinityOrAntiAffinityResourceList from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in resource_id (list)
        _items = []
        if self.resource_id:
            for _item_resource_id in self.resource_id:
                if _item_resource_id:
                    _items.append(_item_resource_id.to_dict())
            _dict['resourceId'] = _items
        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of AffinityOrAntiAffinityConstraintForComputeAffinityOrAntiAffinityResourceList from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "resourceId": [Identifier.from_dict(_item) for _item in obj["resourceId"]] if obj.get("resourceId") is not None else None
        })
        return _obj


