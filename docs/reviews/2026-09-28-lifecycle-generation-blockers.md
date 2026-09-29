# Lifecycle review: documented Arcus contract and local verification

Updated 29 September 2026 after reading the public provider documentation and the OpenAPI schema linked from it. This supersedes the preliminary assessment that a new conditional Arcus API was an established release prerequisite.

The two review threads remain open for an explicit maintainer disposition:

- [Update and replacement allocations](https://github.com/i3dnet/nakama-i3d/pull/42#discussion_r4095589354).
- [Delete and replacement allocations](https://github.com/i3dnet/nakama-i3d/pull/42#discussion_r4095589435).

## Verified provider contract

- [Matchmaker allocation](https://docs.i3d.net/game-hosting/game-integration/matchmaker-allocation) selects an empty ONLINE instance, moves it through ALLOCATING to ALLOCATED, and documents host-agent rejection of already allocating/allocated instances (76005/76006). Reuse follows an explicit release back to ONLINE or process shutdown/restart. Being empty alone does not release an allocated instance.
- [Automatic scaling](https://docs.i3d.net/game-hosting/processes/automatic-scaling) removes unused instances. Patrick's explanation, relayed by the project owner, confirms occupied instances are excluded from automatic allocation/scaling, while explicit restarts still execute.
- [ApplicationInstance fields](https://docs.i3d.net/game-hosting/elements/application/applicationinstance) and the published OpenAPI schema mark numPlayers, numPlayersMax and ordinary instance status fields read-only. [Arcus V2 live-state messages](https://docs.i3d.net/game-hosting/game-integration/index/index-1/request-response#live-state) carry player counts from the game server to the host agent. The separate status endpoint changes operational status.
- [Metadata updates](https://docs.i3d.net/game-hosting/elements/application/metadata) merge keys. Omitted keys survive, an empty string is a value, and an explicit null deletes a key. Allocation metadata uses the same merge behavior.
- The [restart endpoint](https://docs.i3d.net/api/api_one) stops and starts an instance using the supplied/configured stop method, defaulting to hard kill. [Arcus soft stop](https://docs.i3d.net/game-hosting/game-integration/index/index-1/request-response#soft-stop) requires explicit configuration; a positive timeout eventually permits a hard stop.

## Corrections to this integration

The adapter now sends a metadata-only PUT without reading and replaying the full provider object. Player-count reports affect Nakama's local admission state; they do not write ONE telemetry. Explicit null metadata values survive serialization on both allocation and update.

The local provider service accepts metadata patches, merges them, rejects unsupported read-only fields, and excludes occupied ONLINE instances from allocation. Its strict rejection of read-only writes is a local test guard, not a claim that the real API rejects rather than ignores such fields. The previous mock incorrectly accepted numPlayers writes and replaced metadata, so the earlier player-count propagation smoke assertion was not valid provider-contract evidence.

## Local verification and its limits

This project has no staging environment. The verification path is local unit/race tests, captured HTTP requests, the contract mock, and real Nakama/PostgreSQL containers loading the plugin. CI repeats the automated suites; it does not connect to a live i3D fleet.

Coverage includes metadata-only request bodies; null versus empty-string semantics; preservation of Nakama counts when provider telemetry differs; exclusion of occupied, ALLOCATING and ALLOCATED instances; and allocation, metadata update, restart and reuse of the same instance. The smoke checks metadata changes/deletion and proves that a Nakama count report does not overwrite provider telemetry.

The mock completes restarts synchronously. These checks do not emulate the Arcus transport, host-agent scheduling or real propagation delay, and do not prove every ordering of an in-flight request across release/reallocation.

## Supported lifecycle and remaining review decision

Use one owner for end-of-session release. Stop producing reports and finish outstanding updates before releasing the server. Choose either the adapter restart, process exit, or returning to ONLINE; do not independently release the instance and subsequently send cleanup for its previous match. Do not blindly retry a restart or allocation after an ambiguous network failure. A restart HTTP response acknowledges the provider operation; it is not a game-server readiness signal. New allocations still use the normal provider allocation handshake.

Requests identify the application instance, not an individual match. The documented allocation guards prevent ordinary double allocation, but they are not an atomic ownership check on explicit metadata updates or restarts. A delayed/duplicate request after an independent release could still target a reused instance; local storage version checks cannot undo a remote action. No real Arcus reproduction of that ordering has been performed, and local tests do not establish that it is impossible.

The maintainer should disposition the two threads against this documented lifecycle boundary. They have not been silently marked fixed, and no allocation-token protocol or provider API change is claimed or required by this assessment. If stronger protection against stale callers is required, that is a separate design decision rather than an inferred property of the current API.
