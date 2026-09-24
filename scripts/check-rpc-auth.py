#!/usr/bin/env python3
"""Exercise lifecycle authorization through Nakama's HTTP authentication layer."""
import base64
import json
import sys
import urllib.error
import urllib.parse
import urllib.request
import uuid

base = sys.argv[1]
def request(path, body, headers=None):
    data = json.dumps(body).encode()
    req = urllib.request.Request(base + path, data, {"Content-Type": "application/json", **(headers or {})})
    try:
        with urllib.request.urlopen(req, timeout=10) as response:
            return response.status, json.load(response)
    except urllib.error.HTTPError as error:
        return error.code, json.load(error)

status, account = request("/v2/account/authenticate/device?create=true",
                          {"id": "rpc-check-" + uuid.uuid4().hex},
                          {"Authorization": "Basic " + base64.b64encode(b"defaultkey:").decode()})
assert status == 200, (status, account)
for rpc in ("update_instance_info", "delete_instance_info"):
    status, body = request("/v2/rpc/" + rpc, "not-json",
                           {"Authorization": "Bearer " + account["token"]})
    assert status == 403 and body["code"] == 7, (rpc, status, body)
    status, body = request("/v2/rpc/" + rpc + "?http_key=local-test", "not-json")
    assert status == 400 and body["code"] == 3, (rpc, status, body)
    status, body = request("/v2/rpc/" + rpc + "?http_key=wrong-key", "not-json")
    assert status == 401, (rpc, status, body)
print("Lifecycle RPCs: player rejected; valid HTTP key reaches validation; invalid key rejected.")
