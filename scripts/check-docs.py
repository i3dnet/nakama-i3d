#!/usr/bin/env python3
"""Compile the actual partner-guide blocks in a consumer without replace.

Without an argument, expose this checkout through a temporary local Go proxy.
With a commit/tag argument, use the normal public Go resolver.
"""
import json
import os
from pathlib import Path
import re
import subprocess
import sys
import tempfile
import zipfile

root = Path(__file__).resolve().parent.parent
module = "github.com/i3dnet/nakama-i3d"
blocks = re.findall(r"~~~go\n(.*?)\n~~~", (root / "online-docs.md").read_text(), re.S)
complete = [block for block in blocks if block.startswith("package main")]
filters = [block for block in blocks if block.startswith("filters :=")]
imports = [block for block in blocks if block.startswith("import (")]
assert len(complete) == 2 and len(filters) == 1 and len(imports) == 1, "guide snippet layout changed"
assert len(blocks) == len(complete) + len(filters) + len(imports), "unverified Go block in guide"
with tempfile.TemporaryDirectory(prefix="i3d-docs-") as directory:
    tmp = Path(directory)
    env = {**os.environ, "GOWORK": "off", "GOPRIVATE": "", "GONOPROXY": ""}
    if len(sys.argv) > 1:
        version = sys.argv[1]
    else:
        version = "v0.0.0-local"
        proxy = tmp / "proxy" / module / "@v"
        proxy.mkdir(parents=True)
        (proxy / (version + ".mod")).write_bytes((root / "go.mod").read_bytes())
        (proxy / (version + ".info")).write_text(json.dumps({"Version": version, "Time": "2026-09-22T00:00:00Z"}))
        (proxy / "list").write_text(version + "\n")
        sources = [root / "go.mod", root / "go.sum", *root.glob("*.go")]
        for child in ("config", "internal"):
            sources.extend((root / child).rglob("*.go"))
        with zipfile.ZipFile(proxy / (version + ".zip"), "w", zipfile.ZIP_DEFLATED) as archive:
            for source in sources:
                archive.write(source, f"{module}@{version}/{source.relative_to(root)}")
        env["GOPROXY"] = (tmp / "proxy").as_uri() + ",https://proxy.golang.org"
        env["GONOSUMDB"] = module
        # Do not reuse a cached local-test version from another checkout.
        env["GOMODCACHE"] = str(tmp / "cache")
    consumer = tmp / "consumer"
    consumer.mkdir()
    def go(*args):
        subprocess.run(["go", *args], cwd=consumer, env=env, check=True)
    try:
        go("mod", "init", "example.com/i3d-guide")
        go("get", module + "@" + version)
        for index, block in enumerate(complete):
            (consumer / f"guide{index}.go").write_text(block + "\n")
        (consumer / "guide_imports.go").write_text(
            "package main\n" + imports[0] +
            "\nvar _ = fleetmanager.NewI3dFleetManager\nvar _ = fleetconfig.NewConfigFromRuntime\n")
        (consumer / "filters_test.go").write_text(
            'package main\nimport ("testing"; fleetmanager "github.com/i3dnet/nakama-i3d")\n'
            'func TestDocumentedFilters(t *testing.T) {\n' + filters[0] +
            '\nif metadata["i3dFilters"] == "" { t.Fatal("missing filters") }\n}\n')
        go("mod", "tidy")
        assert "replace " not in (consumer / "go.mod").read_text()
        go("test", "-mod=readonly", "./...")
        go("vet", "-mod=readonly", "./...")
        go("build", "-mod=readonly", "-trimpath", "-buildmode=plugin", "-o", "guide.so", ".")
        print("Partner guide compiled and filters ran without replace: " + version)
    finally:
        # Go marks its module-cache contents read-only.
        for path in tmp.rglob("*"):
            if not path.is_symlink():
                path.chmod(path.stat().st_mode | 0o700)
