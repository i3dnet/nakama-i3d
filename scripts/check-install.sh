#!/usr/bin/env bash
# Verify a remotely available commit or tag through the public Go resolver.
set -euo pipefail
version="${1:?Usage: scripts/check-install.sh <published-commit-or-tag>}"
root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
tmp="$(mktemp -d)"
trap 'chmod -R u+w "$tmp"; rm -rf "$tmp"' EXIT
export GOWORK=off GOMODCACHE="$tmp/modcache" GOPRIVATE= GONOPROXY= GONOSUMDB=
mkdir "$tmp/consumer"
cd "$tmp/consumer"
go mod init example.com/i3d-consumer
go get "github.com/i3dnet/nakama-i3d@$version"
# Compile the real consumer example, with no replace directives or private access.
cp "$root/plugin/example/src/cmd/main/main.go" .
go mod tidy
go list -m github.com/i3dnet/nakama-i3d
go list github.com/i3dnet/nakama-i3d github.com/i3dnet/nakama-i3d/config
go test -mod=readonly ./...
go build -mod=readonly -trimpath -buildmode=plugin -o consumer.so .
