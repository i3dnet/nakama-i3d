# Local verification

Run the root tests with Go 1.27.1: go test -mod=readonly -race -count=1 ./... and go vet -mod=readonly ./.... Run the same commands in plugin/example/src. The mock_api_server/src and nakama_test_client/src modules use Go 1.24.1.

sh scripts/check-plugin.sh builds and loads the normal plugin in Nakama 3.41.0 and tests lifecycle RPC authentication. sh scripts/smoke.sh runs the full scenario with Docker Compose, Python 3 and Go installed. Each script creates its own Compose project with random localhost ports and a temporary PostgreSQL database, and cleans up only that project.

The smoke script authenticates two clients, opens real sockets and submits one ticket each. It checks one allocation, metadata and connection notifications, storage before notification, forced native StorageWriteRetry conflicts, concurrent Join capacity, rejected player RPCs, trusted lifecycle update/restart, recovery after a missed update over multiple provider pages, invalid readiness and allocation timeout without blind retry. Unit tests deterministically cover late/duplicate completion and race boundaries.

The i3d_smoke build tag adds test-only server RPCs and the storage-before-notification assertion. The normal plugin build excludes them. Never distribute the smoke image. The mock service's /_test routes are fixture controls and use the test-token credential; they do not exist in the production adapter.

For an interactive local environment, run docker compose -f _infra/docker-compose.yml up --build --wait. Optional clients: docker compose -f _infra/docker-compose.yml --profile clients up --build --scale nakama-test-client=2. Services have health-based dependencies and no fixed container names. PostgreSQL 16 uses the new data_v16 volume; an old PostgreSQL 12 data volume is left untouched and is not automatically migrated. Do not run down --volumes against a development project containing data you need.

These tests use a local One API simulation. They do not establish real fleet credentials, Arcus readiness or provider propagation timings; those remain staging release checks.
