"""Stub HTTP server for the generated client: each test registers canned responses
and inspects the requests the client sent."""

import json
import threading
from http.server import BaseHTTPRequestHandler, HTTPServer
from pathlib import Path

import pytest

import kubevim_client

TESTDATA = Path(__file__).resolve().parent.parent / "testdata"


def golden(path):
    """Golden JSON from test/testdata, shared with the Go wire tests."""
    return json.loads((TESTDATA / path).read_text())


class Stub:
    def __init__(self):
        self.routes = {}
        self.requests = []

    def route(self, method, path, status=200, body=None, headers=None, content_type="application/json"):
        self.routes[(method, path)] = (status, body, headers or {}, content_type)


@pytest.fixture
def stub():
    state = Stub()

    class Handler(BaseHTTPRequestHandler):
        def handle_one(self):
            length = int(self.headers.get("Content-Length") or 0)
            body = self.rfile.read(length) if length else b""
            path = self.path.split("?")[0]
            state.requests.append({"method": self.command, "path": self.path, "headers": dict(self.headers), "body": body})
            status, payload, headers, content_type = state.routes.get((self.command, path), (404, None, {}, ""))
            data = b"" if payload is None else json.dumps(payload).encode()
            self.send_response(status)
            if data:
                self.send_header("Content-Type", content_type)
            self.send_header("Content-Length", str(len(data)))
            for name, value in headers.items():
                self.send_header(name, value)
            self.end_headers()
            self.wfile.write(data)

        do_GET = do_POST = do_PATCH = do_DELETE = handle_one

        def log_message(self, *args):
            pass

    server = HTTPServer(("127.0.0.1", 0), Handler)
    thread = threading.Thread(target=server.serve_forever, daemon=True)
    thread.start()
    state.url = f"http://127.0.0.1:{server.server_port}"
    yield state
    server.shutdown()


@pytest.fixture
def client(stub):
    config = kubevim_client.Configuration(host=stub.url, access_token="token")
    api_client = kubevim_client.ApiClient(config)
    api_client.set_default_header("Version", "1.0.0")
    return api_client
