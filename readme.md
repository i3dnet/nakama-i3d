# Nakama and i3D.net

Allocate existing online i3D game servers from Nakama matchmaking, deliver connection notifications, and maintain local session state. i3D manages hosts, deployment and available fleet capacity; Arcus carries allocation metadata to the headless server. This adapter implements Nakama's FleetManager interface.

This candidate is under review. Main at 9a20297 is the legacy Nakama 3.26 integration; the two review PRs target the exact combination below. No new tag or production release has been published.

| Component | Tested candidate |
| --- | --- |
| Nakama / pluginbuilder | 3.41.0 |
| nakama-common | 1.48.0 |
| Go | 1.27.1 |
| Shared protobuf | 1.36.12 |
| Implementation commit | 60cea4930eae7c464d6f4342663a6985ca9f0615 |

The Dockerfile pins the matching image digests. Other runtime/toolchain combinations need their own build/load test.

## Install and integrate

~~~sh
go mod init example.com/nakama-game
go get github.com/i3dnet/nakama-i3d@60cea4930eae7c464d6f4342663a6985ca9f0615
~~~

Import github.com/i3dnet/nakama-i3d (package fleetmanager) and github.com/i3dnet/nakama-i3d/config. External installation needs no clone or local replace. Pin the approved tag after release; @latest is not the candidate path while main retains the old layout.

Follow the complete [partner guide draft](online-docs.md) for compiling InitModule/matchmaking examples, Arcus responsibilities, filters, notifications and headless-server lifecycle calls. The [example source](plugin/example/src/cmd/main/main.go) is also built and loaded by CI.

## Run locally

Docker Compose, Go and Python 3 are needed for the smoke check.

~~~sh
sh scripts/smoke.sh
~~~

This builds a disposable Nakama/PostgreSQL/mock stack, authenticates two clients and checks the complete allocation/lifecycle/recovery flow. It uses random localhost ports and removes only its own test resources.

For an interactive environment:

~~~sh
docker compose -f _infra/docker-compose.yml up --build --wait
docker compose -f _infra/docker-compose.yml logs -f nakama
docker compose -f _infra/docker-compose.yml --profile clients up --build --scale nakama-test-client=2
~~~

The development stack uses a new PostgreSQL 16 data_v16 volume, preserving old PostgreSQL 12 data. See [testing](docs/testing.md) for prerequisites, module checks and test-only build flags.

## Behavior and migration

- Create returns (map[string]string, error); synchronous metadata is nil. Accepted work runs independently of the matchmaking hook, with a deadline and shutdown cancellation. Success follows storage persistence. Process-local callbacks are not durable after a crash.
- One API allocation must return one ALLOCATED instance with a usable public IP/port. Allocation, restart and writes are not blindly retried.
- Delete requests a provider restart. Reconciliation only repairs storage, after complete scoped scans and version checks.
- Lifecycle RPCs reject player-session callers. Trusted game servers use Nakama's runtime.http_key and report the absolute player_count.
- Join has versioned local capacity accounting and partial admission. It does not issue expiring reservations or player authentication tokens; authoritative count updates reset local deduplication.
- An empty List query pages through provider allocations. Nonempty queries use the Nakama index, e.g. +value.player_count:>=2. Preserve the query, limit and listing mode when following a cursor.

Read [compatibility/migration](docs/compatibility.md), [configuration](docs/configuration.md), [reconciliation](docs/reconciliation.md), [metrics](docs/observability.md) and the [release checklist](docs/release-checklist.md).

## Release evidence

The [release checklist](docs/release-checklist.md) links both review PRs and the historical stack and distinguishes unit/race checks, native plugin/smoke tests, clean-consumer installation and live-provider gates. The [Heroic Labs resolution](docs/reviews/heroiclabs-resolution.md) maps the supplied partner comments to changes and evidence. The user will send online-docs.md to Heroic Labs.

Licensed under [MIT](LICENSE).
