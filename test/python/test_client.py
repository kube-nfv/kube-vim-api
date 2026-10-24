"""The generated Python client against the REST conventions of hack/openapi:
status codes, List arrays and paging, ProblemDetails, oneof variants, open enums
and the token every request carries."""

import json

import pydantic
import pytest

import kubevim_client
from kubevim_client import (
    ComputeServiceApi,
    FlavourServiceApi,
    ImageServiceApi,
    NetworkServiceApi,
    StorageServiceApi,
)

from conftest import golden

IMAGE_ID = "5f0c9d2e-3a4b-4c1d-9e8f-0a1b2c3d4e5f"
COMPUTE_ID = "c0ffee00-1234-4abc-8def-0123456789ab"


def dump(model):
    """Every field as JSON, read-only ones included (to_dict drops them for requests)."""
    return unwrap(model.model_dump(by_alias=True, mode="json", exclude_none=True))


def unwrap(node):
    """Replaces oneOf wrappers ({"actual_instance": ..., "one_of_schemas": ...}) by their member."""
    if isinstance(node, dict):
        if "actual_instance" in node:
            return unwrap(node["actual_instance"])
        return {k: unwrap(v) for k, v in node.items()}
    if isinstance(node, list):
        return [unwrap(v) for v in node]
    return node


def test_create_returns_201_with_location(stub, client):
    image = golden("image/image.json")
    stub.route("POST", "/image/v1/images", 201, image, {"Location": f"/image/v1/images/{IMAGE_ID}"})
    req = kubevim_client.CreateImageRequest.from_dict(golden("image/create_image_request.json"))
    resp = ImageServiceApi(client).create_image_with_http_info(req)
    assert resp.status_code == 201
    assert resp.headers["Location"].endswith(IMAGE_ID)
    assert resp.data.image_id == image["imageId"]


def test_delete_returns_none_on_204(stub, client):
    stub.route("DELETE", f"/image/v1/images/{IMAGE_ID}", 204)
    assert ImageServiceApi(client).delete_image(IMAGE_ID) is None


def test_list_is_an_array_with_link_paging(stub, client):
    images = golden("image/list_images_response.json")["images"]
    link = '</image/v1/images?nextpage_opaque_marker=abc>; rel="next"'
    stub.route("GET", "/image/v1/images", 200, images, {"Link": link})
    resp = ImageServiceApi(client).list_images_with_http_info(nextpage_opaque_marker="prev")
    assert [i.image_id for i in resp.data] == [i["imageId"] for i in images]
    assert resp.headers["Link"] == link
    assert "nextpage_opaque_marker=prev" in stub.requests[-1]["path"]


@pytest.mark.parametrize(
    "api, method, path, file",
    [
        (ImageServiceApi, "get_image", f"/image/v1/images/{IMAGE_ID}", "image/image.json"),
        (FlavourServiceApi, "get_flavour", f"/flavour/v1/flavours/{IMAGE_ID}", "flavour/flavour.json"),
        (NetworkServiceApi, "get_network", f"/network/v1/networks/{IMAGE_ID}", "network/network.json"),
        (NetworkServiceApi, "get_subnet", f"/network/v1/subnets/{IMAGE_ID}", "network/subnet.json"),
        (ComputeServiceApi, "get_compute", f"/compute/v1/computes/{IMAGE_ID}", "compute/compute.json"),
        (ComputeServiceApi, "get_placement_group", f"/compute/v1/placementGroups/{IMAGE_ID}", "compute/placement_group.json"),
        (StorageServiceApi, "get_volume", f"/storage/v1/volumes/{IMAGE_ID}", "storage/volume.json"),
    ],
)
def test_real_responses_parse(stub, client, api, method, path, file):
    body = golden(file)
    stub.route("GET", path, 200, body)
    got = getattr(api(client), method)(IMAGE_ID)
    assert dump(got) == body


def test_snapshot_responses_parse(stub, client):
    body = golden("compute/compute_snapshot.json")
    stub.route("GET", f"/compute/v1/computes/{COMPUTE_ID}/snapshots/{IMAGE_ID}", 200, body)
    assert dump(ComputeServiceApi(client).get_compute_snapshot(COMPUTE_ID, IMAGE_ID)) == body
    body = golden("storage/volume_snapshot.json")
    stub.route("GET", f"/storage/v1/volumes/{COMPUTE_ID}/snapshots/{IMAGE_ID}", 200, body)
    assert dump(StorageServiceApi(client).get_volume_snapshot(COMPUTE_ID, IMAGE_ID)) == body


def test_problem_details_raise_with_type(stub, client):
    problem = {
        "type": "https://kubevim.kubenfv.io/problems/not-found",
        "title": "Not Found",
        "status": 404,
        "detail": "image not found",
    }
    stub.route("GET", f"/image/v1/images/{IMAGE_ID}", 404, problem, content_type="application/problem+json")
    with pytest.raises(kubevim_client.exceptions.NotFoundException) as err:
        ImageServiceApi(client).get_image(IMAGE_ID)
    assert json.loads(err.value.body)["type"] == problem["type"]


def test_unknown_enum_value_is_tolerated(stub, client):
    body = golden("compute/compute.json")
    body["runningState"] = "COMPUTE_RUNNING_STATE_HIBERNATED"
    stub.route("GET", f"/compute/v1/computes/{COMPUTE_ID}", 200, body)
    got = ComputeServiceApi(client).get_compute(COMPUTE_ID)
    assert got.running_state == kubevim_client.ComputeRunningState.UNKNOWN_DEFAULT_OPEN_API


def test_oneof_takes_exactly_one_member():
    kubevim_client.ImageSource.from_dict({"http": {"url": "https://example.com/a.img"}})
    with pytest.raises((ValueError, pydantic.ValidationError)):
        kubevim_client.ImageSource.from_dict({"http": {"url": "https://example.com/a.img"}, "glance": {"glanceImageId": "5f0c9d2e-3a4b-4c1d-9e8f-0a1b2c3d4e5f"}})
    with pytest.raises((ValueError, pydantic.ValidationError)):
        kubevim_client.ImageSource.from_dict({})


def test_every_request_carries_the_token(stub, client):
    stub.route("DELETE", f"/image/v1/images/{IMAGE_ID}", 204)
    stub.route("GET", "/image/v1/imageSources", 200, ["IMAGE_SOURCE_TYPE_HTTP"])
    api = ImageServiceApi(client)
    api.delete_image(IMAGE_ID)
    assert api.list_image_sources() == [kubevim_client.ImageSourceType.IMAGE_SOURCE_TYPE_HTTP]
    for req in stub.requests:
        assert req["headers"]["Authorization"] == "Bearer token"
