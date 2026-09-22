#!/usr/bin/env sh
set -eu
root="$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)"
project="i3d-smoke-$(date +%s)-$$"
compose() { docker compose -p "$project" -f "$root/_infra/docker-compose.test.yml" -f "$root/_infra/docker-compose.smoke.yml" "$@"; }
cleanup() {
  code=$?
  if [ "$code" -ne 0 ]; then compose logs --no-color --tail=100 nakama mock_server; fi
  compose down --volumes --remove-orphans
}
trap cleanup EXIT
compose up --build --wait --wait-timeout 180 nakama
address="$(compose port nakama 7350)"
provider="$(compose port mock_server 8080)"
python3 "$root/scripts/check-rpc-auth.py" "http://$address"
cd "$root/nakama_test_client/src"
GOTOOLCHAIN=go1.24.1 go run -mod=readonly ./cmd/smoke "http://$address" "http://$provider"
