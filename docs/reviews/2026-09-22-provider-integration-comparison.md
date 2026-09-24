# Provider integration comparison for Nakama–i3D

Reviewed 2026-09-22. This compares the infrastructure integrations linked from Heroic Labs' concepts index with i3D main at f26ac9d83ebb1a70f492db70b69eacf842403cb6. It informs the [implementation plan](../superpowers/plans/2026-09-22-nakama-i3d-compatibility-and-reliability.md); it does not implement changes.

The useful shared architecture is a public Go module, Nakama's FleetManager/callback interfaces, provider-specific allocation, and a stored view of sessions. The providers differ substantially in readiness, reservations and shutdown. Preserve i3D's allocation model while tightening its guarantees; do not copy another adapter wholesale.

## Sources and scope

Read the current [concepts index](https://heroiclabs.com/docs/nakama/guides/concepts/), its [GameLift](https://heroiclabs.com/docs/nakama/guides/concepts/gamelift-integration/), [Edgegap](https://heroiclabs.com/docs/nakama/guides/concepts/edgegap-integration/), [i3D](https://heroiclabs.com/docs/nakama/guides/concepts/i3d-integration/), [headless server authentication](https://heroiclabs.com/docs/nakama/guides/concepts/headless-server-auth/) and [PurrNet](https://heroiclabs.com/docs/nakama/guides/concepts/purrnet-integration/) pages. PurrNet is a client lobby/networking integration, not another infrastructure FleetManager; it was reviewed for scope, not used as a server allocation reference.

Source inspection was pinned to these repository revisions:

| Repository | Revision / commit date | Files inspected |
| --- | --- | --- |
| [heroiclabs/nakama-gamelift](https://github.com/heroiclabs/nakama-gamelift/tree/da1ffe91d8e13ab0160edc7d5d22c6269fc923ec) | da1ffe91, 2026-07-02 | Root module/build/example; complete fleet manager implementation |
| [edgegap/nakama-edgegap](https://github.com/edgegap/nakama-edgegap/tree/20d849e39d46f6fb1be849b6a15e3fb3375c0eb1) | 20d849e3, 2026-06-18 | Root module/build/example; fleet manager, provider requests, event handlers, storage and models |
| [edgegap/edgegap-server-nakama-plugin-unity](https://github.com/edgegap/edgegap-server-nakama-plugin-unity/tree/8c4c0bce2d7a0defa3ec55b678d7a8b4e42108fa) | 8c4c0bce, 2026-03-26 | ServerAgent and API transport |

Also checked the released [common 1.48 interface](https://github.com/heroiclabs/nakama-common/blob/v1.48.0/runtime/runtime.go), [Nakama 3.41 callback handler](https://github.com/heroiclabs/nakama/blob/v3.41.0/server/fleet_manager_callback_handler.go), and i3D's [allocation lifecycle](https://docs.i3d.net/game-hosting/game-integration/matchmaker-allocation) and [One API reference](https://docs.i3d.net/api/api_one).

This is source/documentation analysis. Peer integrations were not built or run against AWS, Edgegap or i3D. Observed source weaknesses below are not claims of confirmed production incidents.

## Primary reference: Heroic Labs' GameLift adapter

Per the user's direction, use this repository as the primary architectural reference, with Edgegap as a secondary comparison. The exact released nakama-common interface remains authoritative for the target runtime. Rechecked GameLift HEAD on 2026-09-22: still da1ffe91d8e13ab0160edc7d5d22c6269fc923ec.

The [June 19 migration](https://github.com/heroiclabs/nakama-gamelift/commit/6c64227164759d5532fa0ce535aefa8b011e8ed0) moved GameLift from Nakama 3.26.0/common 1.36.0 to 3.39.0/common 1.46.0. It changed Create to return metadata plus error, returned the AWS placement ID, updated the interface assertion, and changed the Go/dependency/image pins together. The core allocation flow remained largely unchanged. This is a useful migration precedent for i3D's same legacy baseline; it does not establish compatibility with our newer 3.41.0 target.

| GameLift responsibility | Corresponding i3D decision |
| --- | --- |
| Constructor validates configuration, registers storage index and lifecycle RPCs | Preserve i3D's existing separation of adapter, provider client and storage; validate before startup |
| Init installs the callback handler before starting workers | Start reconciliation only after Nakama services and the callback handler are usable |
| Create returns acceptance metadata; SQS later resolves the allocation | Separate synchronous rejection from accepted asynchronous work; preserve i3D's native completion mechanism and use nil return metadata unless there is a real synchronous ID |
| Placement fulfillment calls Get, which persists ACTIVE sessions, before invoking success | Require a valid ready instance and successful persistence before i3D success notification |
| Join invokes AWS player-session APIs | Explain i3D's supported capacity/reservation behavior explicitly; AWS-issued session IDs are provider-specific |
| Empty-query List uses provider search; filtered List uses Nakama's index | Preserve and document both modes; test limits and cursors independently in each mode |
| Update reports player counts/metadata; Delete removes a terminated session record | Keep server-only RPCs; clearly distinguish i3D's restart-on-Delete behavior and storage-only reconciliation |

[Constructor and Init](https://github.com/heroiclabs/nakama-gamelift/blob/da1ffe91d8e13ab0160edc7d5d22c6269fc923ec/fleetmanager/gamelift_fleet_manager.go#L125), [Create](https://github.com/heroiclabs/nakama-gamelift/blob/da1ffe91d8e13ab0160edc7d5d22c6269fc923ec/fleetmanager/gamelift_fleet_manager.go#L185), [Get/List](https://github.com/heroiclabs/nakama-gamelift/blob/da1ffe91d8e13ab0160edc7d5d22c6269fc923ec/fleetmanager/gamelift_fleet_manager.go#L296), [completion worker](https://github.com/heroiclabs/nakama-gamelift/blob/da1ffe91d8e13ab0160edc7d5d22c6269fc923ec/fleetmanager/gamelift_fleet_manager.go#L794).

Two additional documentation checks follow from this closer pass. GameLift's repository README has the updated two-return Create example, although the website guide still has the older form. Its README also shows playerCount in an update payload, while the Go request field is tagged player_count, and includes typed parameter declarations inside a Join call example. Compile Go examples and execute JSON examples through the actual decoder/handler; compilation alone cannot catch wire-format drift. [README](https://github.com/heroiclabs/nakama-gamelift/blob/da1ffe91d8e13ab0160edc7d5d22c6269fc923ec/README.md), [request type](https://github.com/heroiclabs/nakama-gamelift/blob/da1ffe91d8e13ab0160edc7d5d22c6269fc923ec/fleetmanager/gamelift_fleet_manager.go#L578).

## Comparison

| Concern | GameLift source | Edgegap source | Current i3D / implication |
| --- | --- | --- | --- |
| Packaging | Root public module; importable /fleetmanager package | Root public module; importable /pkg/fleetmanager package | Nested private GitLab module identity blocks the expected GitHub install. A root public module is the right fix; package layout is our choice. |
| Version pins | Nakama 3.39.0, common 1.46.0, Go 1.26.3 | Same tuple | i3D pins 3.26.0/common 1.36.0. Keep our proposed 3.41.0 target, tested independently. |
| Create return | Placement ID immediately; final outcome via callback | Deployment ID immediately; final outcome via callback | i3D allocates in a goroutine and has no provider ID at method return. A nil metadata return is reasonable; an extra synchronous provider call is unnecessary. |
| Readiness | Placement fulfillment arrives through SNS/SQS | Provider-ready sets connection data; game-server READY triggers success | i3D's allocation request finishes the provider allocation transition. Validate ALLOCATED and connection data before success. |
| Placement input | Maps supplied per-user regional latencies to AWS | Retrieves player IPs and passes them to Edgegap | i3D uses allocation filters. Do not claim automatic latency selection or collect player IPs without a separate need. |
| Reservations | AWS creates player sessions and returns real session IDs | Nakama stores reservations, active users, available seats and expiry metadata; Join returns no SessionInfo | i3D currently increments local counts. Define a supported reservation contract before promising parity. |
| State recovery | Periodic provider/storage reconciliation | Provider reconciliation plus reservation cleanup | i3D needs an explicit recovery policy for stale records and lost lifecycle updates. |
| Delete | Removes Nakama storage only | Stops the provider deployment and removes storage | i3D restarts an application instance. Document this consequential semantic difference. |
| Server companion | GameLift SDK accepts/removes sessions; server reports lifecycle updates | Unity helper reports READY/STOP/ERROR and connection snapshots | Document i3D's Arcus and Nakama RPC responsibilities separately. |

Packaging/version sources: [GameLift go.mod](https://github.com/heroiclabs/nakama-gamelift/blob/da1ffe91d8e13ab0160edc7d5d22c6269fc923ec/go.mod), [Dockerfile](https://github.com/heroiclabs/nakama-gamelift/blob/da1ffe91d8e13ab0160edc7d5d22c6269fc923ec/Dockerfile), [Edgegap go.mod](https://github.com/edgegap/nakama-edgegap/blob/20d849e39d46f6fb1be849b6a15e3fb3375c0eb1/go.mod), [Dockerfile](https://github.com/edgegap/nakama-edgegap/blob/20d849e39d46f6fb1be849b6a15e3fb3375c0eb1/Dockerfile).

Behavior sources: [GameLift Create/Join/Delete/event workers](https://github.com/heroiclabs/nakama-gamelift/blob/da1ffe91d8e13ab0160edc7d5d22c6269fc923ec/fleetmanager/gamelift_fleet_manager.go), [Edgegap fleet manager](https://github.com/edgegap/nakama-edgegap/blob/20d849e39d46f6fb1be849b6a15e3fb3375c0eb1/pkg/fleetmanager/fleet_manager.go), [event handlers](https://github.com/edgegap/nakama-edgegap/blob/20d849e39d46f6fb1be849b6a15e3fb3375c0eb1/pkg/fleetmanager/event_manager.go), [deployment requests](https://github.com/edgegap/nakama-edgegap/blob/20d849e39d46f6fb1be849b6a15e3fb3375c0eb1/pkg/fleetmanager/edgegap_manager.go).

## Patterns to adopt or adapt

### Persist success before notifying players

Edgegap's game-server READY handler writes the instance first, then invokes the callback. Its source explicitly explains the stale-read failure this ordering prevents. This directly supports changing i3D's current callback-before-storage sequence. Success should mean that subsequent Get/List/Join can observe the agreed state. Provider success followed by storage failure needs an explicit recovery outcome.

[Edgegap event handler, lines 248–258](https://github.com/edgegap/nakama-edgegap/blob/20d849e39d46f6fb1be849b6a15e3fb3375c0eb1/pkg/fleetmanager/event_manager.go#L248)

### Keep i3D's native readiness mechanism

The i3D allocation guide describes the platform selecting an ONLINE instance, moving it to ALLOCATING, receiving the host-agent allocation result, moving it to ALLOCATED, and returning it. Arcus carries allocation metadata to the game server. This supports response validation and an Arcus staging test, rather than introducing a second READY webhook by analogy with Edgegap.

The prose allocation guide still names allocateWithErrors; the current One API reference documents /empty/allocate, which matches our generated client. Resolve such documentation differences against the actual endpoint contract before changing URLs.

[i3D allocation flow](https://docs.i3d.net/game-hosting/game-integration/matchmaker-allocation), [One API endpoint](https://docs.i3d.net/api/api_one)

### Treat reservations as state with an owner

GameLift delegates reservation enforcement and tokens to the provider. Edgegap models reservation and connection lists inside typed, namespaced metadata and derives available seats. Its decoder round-trips generic metadata into a typed struct, avoiding assumptions about Go integer types after JSON storage.

For i3D, retain compatibility with old records, reserve plugin-owned fields, and use versioned writes. If expiring reservations are required, store each reservation's identity and expiry, deduplicate users, define confirmation/disconnection events, and test conflicts. A single player-count increment is insufficient. Do not fabricate provider session credentials.

[GameLift Join](https://github.com/heroiclabs/nakama-gamelift/blob/da1ffe91d8e13ab0160edc7d5d22c6269fc923ec/fleetmanager/gamelift_fleet_manager.go#L492), [Edgegap metadata](https://github.com/edgegap/nakama-edgegap/blob/20d849e39d46f6fb1be849b6a15e3fb3375c0eb1/pkg/fleetmanager/models.go#L5), [decoding and seat calculation](https://github.com/edgegap/nakama-edgegap/blob/20d849e39d46f6fb1be849b6a15e3fb3375c0eb1/pkg/fleetmanager/storage.go#L49)

### Reconcile missed lifecycle updates safely

Both implementations have provider reconciliation workers. The corresponding i3D worker should operate on the configured application/fleet scope, finish all pages before treating absence as evidence, preserve local reservations/capacity, and avoid deleting sessions created during the scan. Failed or partial scans must not purge records. Reconciliation removes or repairs cached records; it must not invoke i3D Delete, which restarts a server.

Make startup failure visible and provide a bounded recovery path. Polling frequency should be configurable and justified by provider load and recovery expectations.

[GameLift worker](https://github.com/heroiclabs/nakama-gamelift/blob/da1ffe91d8e13ab0160edc7d5d22c6269fc923ec/fleetmanager/gamelift_fleet_manager.go#L843), [Edgegap worker](https://github.com/edgegap/nakama-edgegap/blob/20d849e39d46f6fb1be849b6a15e3fb3375c0eb1/pkg/fleetmanager/fleet_manager.go#L235)

### Keep callback guarantees within what Nakama provides

Both peers use FmCallbackHandler, as i3D already does. Released open-source Nakama stores callbacks in a local sync.Map and consumes them with LoadAndDelete. This suppresses repeated delivery to the same registered callback, but provides neither automatic timeout eviction nor persistence across process restart.

Ensure every accepted i3D allocation reaches a terminal callback under normal process operation, including timeout and shutdown. Do not promise durable exactly-once delivery. Any future webhook/queue design must address callback ownership across nodes; a callback ID alone does not prove routing or durability.

[Nakama 3.41 local callback handler](https://github.com/heroiclabs/nakama/blob/v3.41.0/server/fleet_manager_callback_handler.go)

### Document server authentication and observable outcomes

Heroic Labs' headless guide explicitly rejects RPC calls carrying a player user ID and uses the runtime HTTP key for server calls. Apply that guard to i3D's lifecycle RPCs and verify it through Nakama, not only in direct Go tests.

GameLift also records provider errors, placement outcomes and duration through Nakama metrics. Add a small i3D equivalent for allocation result/duration and reconciliation failure. Keep request/instance IDs in structured logs, rather than high-cardinality metric labels, and exclude credentials.

[Headless authentication guide](https://heroiclabs.com/docs/nakama/guides/concepts/headless-server-auth/), [GameLift event metrics](https://github.com/heroiclabs/nakama-gamelift/blob/da1ffe91d8e13ab0160edc7d5d22c6269fc923ec/fleetmanager/gamelift_fleet_manager.go#L813)

## Source patterns to avoid copying

These findings limit how confidently the peer repositories can be treated as reference implementations:

- **Guide/version drift:** Both guides retain older compatibility claims and one-return Create examples, while both inspected repositories implement two returns and pin 3.39.0. GameLift also shows older connection-field access. Compile our guide examples against the supported tuple.
- **RPC authorization gaps:** GameLift UpdateInstanceInfo/DeleteInstanceInfo have no player-context guard. Edgegap's inspected event handlers unpack headers/query parameters but do not enforce the headless guide's user-ID rejection. Endpoint configuration elsewhere may affect exposure; these functions are not authorization templates. [GameLift handlers](https://github.com/heroiclabs/nakama-gamelift/blob/da1ffe91d8e13ab0160edc7d5d22c6269fc923ec/fleetmanager/gamelift_fleet_manager.go#L584), [Edgegap unpack/handlers](https://github.com/edgegap/nakama-edgegap/blob/20d849e39d46f6fb1be849b6a15e3fb3375c0eb1/pkg/fleetmanager/event_manager.go#L41).
- **Pagination and reconciliation hazards:** GameLift's indexed List discards the cursor. Its worker compares all stored sessions against each individual provider page, potentially removing sessions found on other pages. Gather the full provider set before considering deletion. [List](https://github.com/heroiclabs/nakama-gamelift/blob/da1ffe91d8e13ab0160edc7d5d22c6269fc923ec/fleetmanager/gamelift_fleet_manager.go#L330), [worker](https://github.com/heroiclabs/nakama-gamelift/blob/da1ffe91d8e13ab0160edc7d5d22c6269fc923ec/fleetmanager/gamelift_fleet_manager.go#L843).
- **Reservation concurrency/expiry gaps:** Edgegap's update writes omit storage versions. Its cleanup uses one instance timestamp, clears the entire reservation list, and reads only one index page; Join does not refresh that timestamp. The model is useful, but it does not establish correct per-reservation expiry or concurrent admission. [Writes](https://github.com/edgegap/nakama-edgegap/blob/20d849e39d46f6fb1be849b6a15e3fb3375c0eb1/pkg/fleetmanager/storage.go#L282), [Join and cleanup](https://github.com/edgegap/nakama-edgegap/blob/20d849e39d46f6fb1be849b6a15e3fb3375c0eb1/pkg/fleetmanager/fleet_manager.go#L163).
- **Companion delivery gaps:** Edgegap's Unity helper has useful READY/connection/STOP hooks, but its transport sends once, and a failed connection update clears the pending flag without scheduling retry. Adopt the responsibilities, not an assumption of reliable delivery. [ServerAgent](https://github.com/edgegap/edgegap-server-nakama-plugin-unity/blob/8c4c0bce2d7a0defa3ec55b678d7a8b4e42108fa/Runtime/ServerAgent.cs#L169), [API transport](https://github.com/edgegap/edgegap-server-nakama-plugin-unity/blob/8c4c0bce2d7a0defa3ec55b678d7a8b4e42108fa/Runtime/Api.cs#L105).

## Changes carried into the implementation plan

1. Keep the root-module and exact-version decisions; peer layouts validate the distribution model.
2. Make callback completion, timeout cleanup, storage visibility and restart limitations explicit.
3. Validate i3D allocation readiness using its existing provider contract.
4. Add safe periodic reconciliation and tests for multiple pages, failed scans and allocation reuse.
5. Tighten reservation decisions and lifecycle docs without silently committing to a new session service or Unity SDK.
6. Add focused operational metrics, server-authentication smoke checks, and executable documentation examples.

Remaining product questions stay in the plan: actual deployment identity, whether backfill/Join is used, required reservation guarantees, and responsibility for updating the partner guide.
