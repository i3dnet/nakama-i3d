# Lifecycle allocation identity: open review blockers

Two findings remain open in [runtime review 5306795543](https://github.com/i3dnet/nakama-i3d/pull/42#pullrequestreview-5306795543):

- [Update can affect a replacement allocation](https://github.com/i3dnet/nakama-i3d/pull/42#discussion_r4095589354).
- [Delete can restart a replacement allocation](https://github.com/i3dnet/nakama-i3d/pull/42#discussion_r4095589435).

The current RPCs identify an application instance, not the allocation that produced the request. Instance IDs can be reused. Server authentication does not establish allocation ownership, and storage version checks after the provider request cannot undo a remote mutation against a replacement.

The bundled schema exposes an instance-ID PUT and an instance-ID restart POST, with no allocation-generation, ETag or If-Match precondition. The provider createdAt field describes application-instance creation, not allocation identity. Sources are requests.go, internal/clients/application_instance.go, internal/storage/fleet_manager_storage.go and internal/openapi/api/openapi.yaml in the [reviewed runtime tree](https://github.com/i3dnet/nakama-i3d/tree/60cea4930eae7c464d6f4342663a6985ca9f0615).

Reading a token before writing would reject already-stale callers, but reallocation between that read and the unconditional write would remain possible. A process mutex cannot protect another Nakama node or external provider changes. No atomic provider mutation guarantee has been confirmed beyond this bundled contract.

## Required contract and implementation

Confirm a provider-enforced allocation/version condition for both update and restart, or establish a deployment policy that prevents reuse throughout every in-flight lifecycle mutation. The user has been asked whether One API offers such a condition.

Once that guarantee is available:

1. Generate an allocation UUID before allocation and deliver it to Arcus through a trusted reserved metadata field that callers cannot override.
2. Persist that same UUID as local identity. Currently the local UUID is generated after allocation and never reaches the headless server.
3. Require the server to echo allocation_id in lifecycle RPCs; reject missing/stale identity and absent/invalidated sessions before provider calls. This requires a headless-server protocol migration.
4. Pass the expected identity through provider helpers and enforce the confirmed atomic precondition. Preserve the reserved identity when ordinary metadata is replaced.
5. Bind direct Go lifecycle calls and late-allocation cleanup to the expected allocation as well. Looking up the newest identity on behalf of an old caller does not prove that caller owns it.
6. Recheck identity during versioned local mutations and retain conditional cache cleanup.

Regressions must cover stale RPC identities, a replaced provider allocation, zero remote mutations after rejection, metadata-token preservation, and reallocation deliberately paused between provider read and mutation. The last case must be rejected by the provider precondition itself.

Both review threads remain unresolved and this is a release blocker. No partial read-then-write guard or new RPC payload has been shipped under a claim of closing the race.
