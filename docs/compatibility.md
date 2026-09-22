# Compatibility and migration

The candidate targets exactly Nakama 3.41.0, nakama-common 1.48.0, Go 1.27.1 and protobuf 1.36.12. The Dockerfile pins the matching runtime and pluginbuilder image index digests. Build with that pluginbuilder; a plugin compiled with another Go toolchain or shared dependency version may fail to load.

The legacy source at f26ac9d83ebb1a70f492db70b69eacf842403cb6 pins Nakama 3.26.0/common 1.36.0. Heroic Labs reported successful use with 3.26.x. That is separate from the new candidate; no broader compatibility range is promised.

Use the root import github.com/i3dnet/nakama-i3d (package fleetmanager) and github.com/i3dnet/nakama-i3d/config. Create now returns (map[string]string, error). The first value is currently nil; connection information arrives asynchronously in the callback.

Create rejects canceled calls and invalid capacity/user lists before registering a callback. Accepted allocations retain context values but survive normal matchmaker-hook cancellation. They have a 120-second default deadline and are canceled on graceful shutdown. Set shutdown_grace_sec to a positive value (the example uses 15 seconds) so Nakama invokes shutdown hooks. Success is reported only after storage succeeds. If provider allocation succeeds but storage fails, the callback reports an error; the adapter does not restart the allocated instance as an automatic rollback.

Each accepted callback receives one terminal outcome while this process remains running. Error outcomes carry nil instance, sessions and metadata. Calls without users return nil sessions. Callbacks are local and are not durable across process restart; they do not route between nodes. The example sends notifications using a separate bounded context.

Validation on 2026-09-22: root and example race tests and vet pass; scripts/check-plugin.sh built and loaded the Linux ARM64 plugin in the exact runtime with disposable PostgreSQL. The CI plugin job repeats build/load on Linux AMD64. The full lifecycle smoke test and live i3D staging are separate checks.

Lifecycle RPCs update_instance_info and delete_instance_info require trusted server HTTP-key authentication. Player-session calls return PERMISSION_DENIED before payload processing. Nakama validates the HTTP key; the adapter's user-context check is not a standalone authentication mechanism. Configure a private runtime.http_key for the headless servers and verify existing callers before deploying this migration. The RPC names and player_count field are unchanged. Malformed payloads, unknown fields, empty IDs, negative counts and plugin-owned metadata return INVALID_ARGUMENT. Direct internal Go Update/Delete calls remain supported.
