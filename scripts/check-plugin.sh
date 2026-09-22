#!/usr/bin/env sh
set -eu
root="$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)"
project="i3d-check-$(date +%s)-$$"
compose() { docker compose -p "$project" -f "$root/_infra/docker-compose.test.yml" "$@"; }
cleanup() { compose down --volumes --remove-orphans; }
trap cleanup EXIT
compose up --build --wait --wait-timeout 180 nakama
compose logs --no-color nakama
address="$(compose port nakama 7350)"
python3 "$root/scripts/check-rpc-auth.py" "http://$address"
