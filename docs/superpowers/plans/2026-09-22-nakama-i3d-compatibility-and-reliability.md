# Nakama–i3D Compatibility and Reliability Implementation Plan

> **For agentic workers:** Use superpowers:executing-plans to execute this plan task by task. Use superpowers:subagent-driven-development only if the user chooses delegation. Steps use checkboxes for tracking. The user approved implementation and small PRs on 2026-09-22. Implementation is complete through the candidate PR stack; publication/deployment gates remain open. See ../../release-checklist.md for current evidence.

**Goal:** Make the integration installable from its actual GitHub repository, compatible with a precisely supported Nakama release, and reliable across the allocation/session lifecycle, addressing both the main-branch review and Heroic Labs' comments.

**Architecture:** Publish the reusable Go library at the repository root, with configuration and internal adapters beneath it. Preserve separate example, mock API, and test-client modules. Upgrade the library and example together against one exact Nakama/common/pluginbuilder combination, then repair the confirmed HTTP, authorization, storage, concurrency, and test gaps.

**Provider comparison:** See [the review of GameLift, Edgegap and the linked server companion](../../reviews/2026-09-22-provider-integration-comparison.md). Use heroiclabs/nakama-gamelift as the primary architectural reference and Edgegap as a secondary comparison; the target release's nakama-common interface governs compatibility. Source revisions and limitations are recorded separately. Readiness, reconciliation and callback requirements below incorporate that comparison.

**Tech stack:** Go; Nakama Go runtime; i3D One API; Nakama storage/PostgreSQL; Docker Compose; GitHub Actions.

**Spec:** The requirements, evidence, proposed decisions, and acceptance criteria in this document, including the partner feedback register below. Approved for implementation. Checked items below are supported by the candidate evidence; unchecked live/publication gates are deliberately outstanding.

**Baseline:** Reviewed main and remote main at f26ac9d83ebb1a70f492db70b69eacf842403cb6 on 2026-09-22. The deployed artifact/configuration has not been identified.

## 1. Partner feedback and verified evidence

| ID | Feedback supplied by the user | Verification | Resolution |
| --- | --- | --- | --- |
| HL-01 | Heroic Labs' i3D guide references the outdated github.com/i3d/nakama-fleetmanager location. | The current guide still uses that install/import path and claims Nakama 3.36+ support. | Tasks 2 and 9: establish one public module path, compile the examples, and prepare the partner guide correction. |
| HL-02 | The plugin uses nakama-common 1.36.0; the FleetManager interface has changed; the partner could use Nakama 3.26.x. | Current Docker pins are 3.26.0 and common 1.36.0. The current released runtime requires a different common version and Create signature. | Task 3: upgrade and test an exact version tuple; document legacy and new compatibility separately. Do not infer all possible older compatibility ranges from one partner test. |
| HL-03 | go get github.com/i3dnet/nakama-i3d does not yield a usable library; cloning plus replace works. | There is no root go.mod; the library is nested at plugin/fleetmanager/src and declares a private GitLab module identity. | Task 2: make the root a public module and prove consumer installation/import without a local replace. |

Verified current target on 2026-09-22:
- Nakama v3.41.0, released September 18, 2026.
- nakama-common v1.48.0 and Go 1.27.1, from that release's go.mod.
- Matching nakama-pluginbuilder and runtime image tags/digests must be verified during Task 3.
- Create now returns (map[string]string, error); the existing adapter returns only error.
- Newer Nakama cancels the matchmaker hook context after completion. Reusing it for asynchronous allocation and notification is unsafe.
- Nakama 3.41 native Go writes use runtime.ErrStorageRejectedVersion and StorageWriteRetry; REST codes are not the native error contract. Conditional deletes return a plain error, so confirm a changed version by rereading rather than parsing its message.

Sources:
- [Partner guide](https://heroiclabs.com/docs/nakama/guides/concepts/i3d-integration/)
- [Nakama release notes](https://heroiclabs.com/docs/nakama/getting-started/release-notes/)
- [Released Nakama go.mod](https://github.com/heroiclabs/nakama/blob/v3.41.0/go.mod)
- [Released FleetManager interface](https://github.com/heroiclabs/nakama-common/blob/v1.48.0/runtime/runtime.go)
- [Go module publishing](https://go.dev/doc/modules/publishing)

The partner website and the local untracked online-docs.md are different revisions. For example, the website's lifecycle examples have already corrected some status descriptions that remain wrong locally. Compare each independently.

## 2. Recommended decisions and scope

1. Publish module github.com/i3dnet/nakama-i3d at repository root. Put the existing fleetmanager package files at root, retaining package name fleetmanager. Consumer imports become:

   ~~~go
   import (
       fleetmanager "github.com/i3dnet/nakama-i3d"
       "github.com/i3dnet/nakama-i3d/config"
       "github.com/heroiclabs/nakama-common/runtime"
   )
   ~~~

2. Target Nakama 3.41.0/common 1.48.0/Go 1.27.1 for the new release. Keep the old source revision identifiable for existing 3.26 deployments. Do not promise one plugin binary or one adapter method signature supports both runtimes. Backport urgent security fixes separately if the deployed environment requires it.
3. Use focused PRs: baseline, packaging, compatibility, security, request/storage repairs, integration, documentation/release. Finish the known test-fixture repair before using green tests to assess new changes.
4. Preserve allocation filters, static-token and OAuth support, RPC names, and existing stored session records where possible. Document every required consumer migration.
5. Use i3D's native allocation completion contract: validate the returned ALLOCATED instance and connection data; do not add a READY webhook solely because Edgegap uses one.
6. Add bounded, configurable reconciliation for missed lifecycle updates. Treat incomplete provider scans as inconclusive and never call the public restart/Delete operation to repair cache entries.
7. Treat full reservation expiry and game-server session-token validation as a separate contract decision. Do not fabricate tokens or expand into a new session service merely to populate SessionId.

Alternatives considered:
- Keep the deeply nested module and change only its declared path: fewer moves, but leaves an awkward public install/import path and does not match the partner's root install expectation.
- Fix the old release only: quickest mitigation, but does not resolve current-runtime compatibility.
- Broad rewrite: larger compatibility and review risk than repairing the existing adapter.

## 3. Global constraints

- Planning originally performed no implementation or publication. The user subsequently authorized implementation and PRs; tags, partner messages and deployment still require their separate release gates.
- Use an isolated codex/ branch or worktree for implementation.
- Never commit credentials or .env files. Before every implementation commit, inspect staged content for secrets and verify .env is ignored.
- The user authorized rewriting online-docs.md as the Heroic Labs partner draft. Preserve the exact prior draft at docs/drafts/2026-09-22-online-docs.previous.md. The user will send the replacement to Heroic Labs; do not send it automatically.
- Avoid unrelated dependency upgrades. Match shared plugin dependencies, especially nakama-common and protobuf, to the target runtime and verify actual plugin loading.
- Keep generated API models intact except required import-path rewrites. New behavior belongs in handwritten adapters; identify preservation of existing custom instrumentation before any regeneration.
- Use local mocks and disposable containers for routine verification. A live i3D allocation/restart test requires an identified test fleet and authorization.
- Identify the deployed commit/artifact, Nakama version, configuration source, and game-server RPC authentication before applying a release to the live environment.
- Every behavioral fix gets a regression test that fails before the change and passes afterward. Do not substitute tests of canned responses for assertions on outgoing requests.
- Successful compilation, passing unit tests, loading the plugin, and a real i3D integration are separate evidence levels.

## 4. File map after packaging

| Current path | Proposed path | Responsibility |
| --- | --- | --- |
| plugin/fleetmanager/src/go.mod, go.sum | go.mod, go.sum | Public module identity and library dependencies |
| plugin/fleetmanager/src/fleetmanager/*.go | ./*.go | Public fleet-manager adapter and its tests |
| plugin/fleetmanager/src/config/ | config/ | Configuration loading/validation |
| plugin/fleetmanager/src/internal/ | internal/ | Provider clients, generated API, storage, test doubles |
| plugin/example/src/ | Unchanged | Separate Nakama example module |
| mock_api_server/src/ | Unchanged | Separate mock module |
| nakama_test_client/src/ | Unchanged | Separate test-client module |
| plugin/Dockerfile | Same file; build context becomes repository root | Builds the root library plus example |
| _infra/docker-compose.yml | Same path | Local services and configuration |
| .github/workflows/ci.yml | New | Module tests, race checks, plugin build/load, smoke checks |
| scripts/smoke.sh, scripts/check-install.sh | New | Local lifecycle and external consumer verification |
| docs/compatibility.md, docs/release-checklist.md | New | Supported versions, migration, release evidence |

Task 1 uses the existing paths. Tasks 3 onward use the new root-module paths. Keep support-module imports consistent if their module identities are changed; their names do not block publication of the root library.

## 5. Task order

| Task | Deliverable | Dependencies |
| --- | --- | --- |
| 1 | Passing existing tests and reproducible dependencies | Current main |
| 2 | Public root Go module | 1 |
| 3 | Current Nakama compatibility and async lifetime | 2 |
| 4 | Authorized lifecycle RPCs | 3; verify deployed server authentication before rollout |
| 5 | Correct i3D HTTP requests and response validation | 3 |
| 6 | Correct session storage, capacity, pagination, concurrency | 3, 5 |
| 7 | Consistent config, safe auth and retries | 3, 5 |
| 8 | Reproducible local lifecycle tests and CI | 4–7 |
| 9 | Install verification, docs, partner resolution and release | 2–8 |

Each task ends in a reviewed commit after its checks pass. Do not publish an intermediate packaging/compatibility commit as the supported release before the remaining release criteria pass.

## Task 1: Repair the existing test/build baseline

**Files:** plugin/fleetmanager/src/internal/clients/application_instance_test.go; fleetmanager/fleet_manager_i3d_test.go; affected go.mod/go.sum files; initial .github/workflows/ci.yml.

- [x] Reproduce the two checked-in allocation-test panics using the core's declared Go 1.23.5 toolchain.
- [x] Replace manual test-client construction with the production constructor, then inject the test transport before the first API call:

~~~go
client := NewOneApiClient(cfg, authentication, logger)
client.httpClient = httpClient
return client
~~~

- [x] Give non-retry fixtures explicit attempts and zero delay. Keep backoff assertions in the existing dedicated retry suite.
- [x] Make each mocked HTTP call return a fresh response/body. Existing response reuse can hide multi-call behavior.
- [x] Replace sleep-based async assertions with channels that observe the specific storage call and callback; waiting for a callback does not prove storage has finished.
- [x] Reconcile example dependency files; review the diff and avoid unrelated upgrades.
- [x] Add initial CI per module: go test -mod=readonly, go vet, and core -race. Use Go 1.23.5 for the legacy library/example baseline and 1.24.1 for the current support modules.
- [x] Confirm intentional dependency-file edits are committed; subsequent verification must not rewrite them.

**Validation:**

~~~sh
(cd plugin/fleetmanager/src && go test -mod=readonly -race -count=1 ./... && go vet -mod=readonly ./...)
(cd plugin/example/src && go test -mod=readonly ./...)
(cd mock_api_server/src && go test -mod=readonly ./...)
(cd nakama_test_client/src && go test -mod=readonly ./...)
~~~

**Acceptance:** Existing tests pass without weaker assertions. Distinguish fixture repairs from runtime fixes. Suggested commit: test: repair client fixtures and establish CI baseline.

## Task 2: Make the repository an installable public Go module

**Files:** Move the library according to the file map; update all library/self imports; plugin/example/src/go.mod and imports; plugin/Dockerfile; _infra/docker-compose.yml; .github/workflows/ci.yml; scripts/check-install.sh.

- [x] Set root module identity to github.com/i3dnet/nakama-i3d; move the existing package to root and config/internal directories beneath it.
- [x] Rewrite exact old module imports, handling the old /fleetmanager package suffix separately from /config and /internal. Remove the obsolete nested library go.mod/go.sum rather than leave two competing identities.
- [x] Keep the example a separate consumer module. A local development replace may point three directories upward to the root; published installation instructions and the clean consumer check must use no replace.
- [x] Change plugin image build context to repository root. Update COPY/WORKDIR paths for the example and root library. Require committed dependencies during builds instead of running go mod tidy inside the Dockerfile.
- [x] Update CI's core module path from plugin/fleetmanager/src to repository root.
- [x] Create a minimal consumer project that imports the two public packages and compiles the real InitModule example.
- [x] Before publication, validate the candidate module using a temporary file-based module proxy or module archive. After the candidate commit/tag is available remotely, repeat through the normal Go resolver with GOWORK=off and a fresh module cache. No local replace and no private GitLab credentials may be required.

**Public install contract after publication:**

~~~sh
go mod init example.com/i3d-consumer
go get github.com/i3dnet/nakama-i3d@latest
go list github.com/i3dnet/nakama-i3d github.com/i3dnet/nakama-i3d/config
go build ./...
~~~

During candidate verification, replace @latest with the actual candidate version selected from repository tags; record the exact version in the release evidence. Do not invent a tag or claim public installation was verified before the candidate is available.

**Acceptance:** The reusable package is importable from its GitHub identity. Development replace directives are unnecessary for external consumers. Suggested commit: refactor: publish fleet manager as root Go module.

## Task 3: Upgrade the adapter and example to the current Nakama contract

**Files:** go.mod/go.sum; fleet_manager_i3d.go and tests; plugin/example/src/go.mod/go.sum and cmd/main/main.go; plugin/Dockerfile; internal test doubles; CI; docs/compatibility.md.

- [x] Pin common v1.48.0 and Go 1.27.1 in the library/example. Match the target runtime's shared protobuf version, v1.36.12, and inspect the full shared dependency graph before plugin-load verification.
- [x] Verify and pin matching runtime/pluginbuilder 3.41.0 images. Do not build a Linux plugin using the unrelated host toolchain.
- [x] Add an explicit compile-time compatibility assertion:

~~~go
var _ runtime.FleetManagerInitializer = (*I3dFleetManager)(nil)
~~~

- [x] Adapt Create to return (map[string]string, error), update every call site/mock, and retain the asynchronous callback. For this adapter, propose nil synchronous metadata unless a verified consumer contract requires it; do not invent IDs or credentials for that return value.
- [x] Update the sample call:

~~~go
_, err := fm.Create(ctx, maxPlayers, userIDs, nil, metadata, callback)
if err != nil {
    return "", err
}
return "", nil
~~~

- [x] Compare the compatibility change with GameLift commit 6c64227164759d5532fa0ce535aefa8b011e8ed0, which migrated the same 3.26/common1.36 baseline. Use its coordinated signature/assertion/dependency/image changes as a checklist; retain our independently verified 3.41 target.
- [x] Distinguish immediate validation rejection from accepted asynchronous work. Rejected input must not leave a callback registered; register callbacks before starting accepted work, and initialize Nakama dependencies before starting background workers. Test an allocation that completes immediately.
- [x] Audit the complete released interface, including initializer and callback types, instead of assuming the extra return value is the only migration. Run the newer toolchain's vet checks too; fix newly exposed issues in the existing custom OpenAPI instrumentation with narrow changes that survive regeneration.
- [x] Define accepted async work to survive the hook returning. Reject already-cancelled calls before accepting work; derive a separate bounded context after acceptance, retaining necessary values. Tie shutdown to its lifetime. The example's notification callback must also use a valid bounded context.
- [x] Regression-test: start Create with a hook context, return from the hook and cancel it, then complete the mock allocation; exactly one callback/notification succeeds. Also test timeout and shutdown cancellation.
- [x] Check nil-callback handling, error callback arguments, no-user session behavior, and one terminal callback per accepted operation while the process remains running. Exercise timeout/success races and duplicate completion. Invoke a terminal outcome on timeout so the registered callback is consumed; the released local callback handler has no automatic expiry.
- [x] Document that callbacks are local and not durable across process restart in the reviewed open-source runtime. Do not promise durable exactly-once notification or assume a callback ID routes across nodes. Keep the allocation path within its owning process; test any future distributed completion path separately.
- [x] Build and load the plugin in Nakama 3.41.0; package compilation alone does not validate Go plugin compatibility.

**Acceptance:** Interface assertions, unit/race tests and real plugin loading pass on the exact tuple. Compatibility docs distinguish the legacy 3.26/common1.36 baseline from the new release; they do not claim unrestricted "3.36+" support. Suggested commit: feat: support Nakama 3.41 fleet manager runtime.

## Task 4: Restrict lifecycle RPCs to authorized server callers

**Files:** fleet_manager_i3d.go; requests.go; constants.go; new rpc_test.go; example/config docs.

- [x] Reproduce ordinary-client invocation of update_instance_info and delete_instance_info; configure mocks to reject any provider/storage call.
- [x] Reject contexts containing runtime.RUNTIME_CTX_USER_ID before payload processing. Nakama authenticates server HTTP-key calls; absence of a user ID alone is not an authentication system for an arbitrary new endpoint.
- [x] Add positive tests for legitimate server-context calls.
- [x] Validate JSON, nonempty IDs, nonnegative player counts and reserved metadata ownership. Return INVALID_ARGUMENT for payload problems, PERMISSION_DENIED for unauthorized callers, and controlled operation errors for provider failures.
- [ ] Verify deployed game-server authentication and prepare any migration before enabling the restriction there. Preserve RPC names/payload field names.
- [x] Keep direct internal Go Update/Delete usage working.
- [x] Document the trusted game-server HTTP-key call path using Heroic Labs' headless authentication guide. Prove through Nakama that HTTP-key calls work and ordinary player-session calls are rejected; do not infer authorization from a successful direct handler test.

**Regression shape:**

~~~go
ctx := context.WithValue(context.Background(), runtime.RUNTIME_CTX_USER_ID, "ordinary-player")
_, err := fm.DeleteInstanceInfo(ctx, logger, nil, nil, "{\"id\":\"instance-1\"}")
require.Error(t, err)
var runtimeErr *runtime.Error
require.ErrorAs(t, err, &runtimeErr)
require.Equal(t, PERMISSION_DENIED, runtimeErr.Code)
// No i3D restart or storage delete is expected.
~~~

**Validation:** go test -mod=readonly -race ./

**Acceptance:** Authenticated players cannot mutate or restart instances. Valid game-server callers continue to work. Suggested commit: fix: enforce server authorization for lifecycle RPCs.

## Task 5: Correct provider requests and defensive mapping

**Files:** internal/clients/application_instance.go; application_instance_test.go; new request_contract_test.go; filter_builder.go/tests if needed.

- [x] Capture outgoing HTTP requests with a transport or httptest server. Verify method, URL, headers and body, not only mapped canned responses.
- [x] Reproduce the allocation body null, absent pagination headers, and wrong GET /application/instance-ID path.
- [x] Assign the generated request-builder return values:

~~~go
request = request.MetadataCollection(createMetaData(metaData))
// In ListApplicationInstances:
request = request.RANGEDDATA(createRangedData(limit))
request = request.PAGETOKEN(previousCursor)
~~~

- [x] Use GetApplicationInstance(ctx, instanceID) for the update prefetch, not GetApplicationInstanceApplication.
- [x] Validate nonempty arrays before indexing; validate a usable public address and numeric port in 1–65535 before constructing ConnectionInfo.
- [x] Decide explicit list mapping-error behavior; do not silently present an incomplete page as complete.
- [x] Validate allocation completion against the actual /empty/allocate contract: expected ALLOCATED state, usable endpoint and completed metadata delivery through Arcus. Test unexpected ONLINE/ALLOCATING responses as incomplete or invalid according to that contract; never send CreateSuccess merely because HTTP returned 200. The separate prose allocation guide's allocateWithErrors name is not a reason to change the working endpoint.
- [x] Verify filter encoding exactly once at the wire boundary and names containing spaces/quotes. Resolve the bundled API schema's status-filter restriction before claiming default-list compatibility.
- [x] Ensure internal routing metadata is not forwarded as game metadata.

**Regression matrix:**

| Input | Expected result |
| --- | --- |
| Allocate metadata map=arena | Body contains that metadata |
| List limit 7/cursor page-two | RANGED-DATA results=7 and PAGE-TOKEN page-two |
| Update instance-1 | GET and PUT target that instance |
| Empty successful get/allocate/update response | Error; no panic |
| Missing public IP or invalid port | Mapping error; no panic |
| Two list pages | Correct continuation; no repeated first page |

**Validation:** go test -mod=readonly -race ./internal/clients

**Acceptance:** All captured-request and invalid-response tests pass. Suggested commit: fix: correct i3D allocation and instance requests.

## Task 6: Repair stored sessions and define concurrency behavior

**Files:** internal/storage/fleet_manager_storage.go and new tests; constants.go; fleet_manager_i3d.go/tests; new reconciliation.go/reconciliation_test.go; config/config.go for worker settings; internal/tests/fleet_manager_storage_mock.go.

**Internal listing signature:**

~~~go
ListGameSessionsFromStorage(
    ctx context.Context, query string, limit int, order []string, cursor string,
) ([]*runtime.InstanceInfo, string, error)
~~~

- [x] Return distinguishable errors for empty storage reads, invalid JSON and JSON null. Never index an empty result or swallow decoding errors.
- [x] Test capacity through the actual JSON boundary:

~~~go
original := &runtime.InstanceInfo{Metadata: map[string]any{MaxPlayers: 8}}
data, err := json.Marshal(original)
require.NoError(t, err)
var loaded runtime.InstanceInfo
require.NoError(t, json.Unmarshal(data, &loaded))
capacity, err := getMaxPlayers(&loaded)
require.NoError(t, err)
require.Equal(t, 8, capacity)
~~~

- [x] Accept positive in-range int and integral finite JSON numeric capacities. Reject fractional, negative and overflow values. Add other compatibility representations only if deployed data demonstrates them.
- [x] Document empty-query provider listing versus filtered storage-index listing. Verify limits and continuation separately for both paths, including the final page; define cursors as specific to the selected mode/query.
- [x] Forward storage order/cursor, retain the next cursor, and propagate it through FleetManager.List. Propagate the currently discarded cache-write error in API-backed listing.
- [x] Separate ownership: provider owns address/status and provider metadata; plugin owns configured capacity and any local reservations. Refreshing Get/List/Update must not erase plugin-owned fields.
- [x] Use Nakama versioned writes for concurrent joins and bounded conflict retries. A process-local mutex is insufficient for multiple Nakama nodes. Test target-version conflict codes.
- [x] Test two callers competing for the last slot; no over-admission or lost update. Test old instance IDs reused for a new allocation.
- [x] Document and resolve the current gap between local player counting and Nakama's reservation semantics: duplicate users, expiry, session IDs, reconnection and authoritative player-count updates. Do not claim full reservation support without implementing its acceptance cases.
- [ ] Not selected for this release: if full reservations are selected, use per-reservation expiry and an explicit confirmation event; test repeated users, reconnects, expiry racing with confirmation, and cleanup across every index page. An instance-wide timestamp and unconditional writes are insufficient.
- [x] Add a service-lifetime reconciliation worker with configurable interval and bounded calls. Build the full provider set for the configured application/fleet before evaluating absence; aborted or failed pages must not delete stored sessions. Preserve plugin-owned fields and protect writes/deletes with versions or equivalent generation checks. Exclude sessions created during the scan and guard against reused instance IDs; document eventual-consistency assumptions before treating absence as proof.
- [x] Reconciliation must only repair/remove storage entries, never restart provider instances. Test startup failure/recovery, missed termination updates, two provider pages, a failed later page, simultaneous Create/Join, and reused instance IDs. Stop its ticker/work on shutdown.
- [x] Persist before CreateSuccess so List/Join visibility is part of success. Specify what happens if allocation succeeds but storage fails, with a deterministic failure test; do not blindly restart an allocated instance as rollback.

**Acceptance:** Existing serialized sessions remain readable; capacity survives refresh; cursors work; missing/corrupt data returns errors; concurrent admitted joins respect the agreed limit. Any intentionally limited reservation contract is explicit in docs and partner follow-up. Suggested commit: fix: preserve and synchronize fleet session state.

## Task 7: Unify configuration and make authentication/retries safe

**Files:** config/config.go and new tests; internal/clients/authentication.go, one_api_client.go, retry_executor.go, application_instance.go and tests/mocks.

**Internal signatures, updated together:**

~~~go
GetAccessToken(ctx context.Context) (string, error)
GetClient(ctx context.Context) (*openapi.APIClient, error)
Run(ctx context.Context, attempts int, fn func() error) error
~~~

- [x] Table-test equivalent runtime-env and OS-env input: static tokens, OAuth, all three retry settings, invalid booleans/durations, zero attempts, backoff bounds.
- [x] Keep I3D_BASE_URL canonical; accept the documented I3D_API_URL alias, with BASE_URL winning when both exist. Test and document precedence.
- [x] Parse runtime retry overrides instead of leaving defaults. Do not silently fall back to unrelated credentials when runtime configuration is present but invalid.
- [x] Replace the accidental .envsetting.json fallback with an explicit setting.json path. Resolve paths from PROJECT_ROOT or runtime working directory, not the source location baked into a plugin.
- [x] Honor AuthenticationUrl with a context-aware request. Return OAuth errors through GetClient instead of panicking; validate token/expiry responses.
- [x] Synchronize token state and client initialization/refresh. Avoid recursive mutex acquisition between GetAccessToken and IsExpired.
- [x] Use a barrier-started concurrency test and -race; assert refreshed credentials are actually applied.
- [x] Replace uninterruptible Sleep with a timer/select on ctx.Done(). Test cancellation during backoff.
- [x] Add validated reconciliation interval/timeout settings consistently to both configuration sources. Define whether reconciliation can be disabled and expose recovery failure without reporting provider readiness.
- [x] Give provider calls explicit timeout bounds based on the tested allocation/Arcus timing. Document the selected value and test its expiry.
- [x] Replace the always-retry policy with verified classification. Do not blindly retry allocation after ambiguous transport failure: the first attempt may already have allocated a server. Only retry provider failures whose documented semantics make repetition safe.

**Validation:** go test -mod=readonly -race ./config ./internal/clients ./

**Acceptance:** Equivalent configuration sources agree; auth outages produce errors; concurrency is race-free; cancellation and retry classification have behavioral tests. Suggested commit: fix: synchronize provider authentication and bound retries.

## Task 8: Build a realistic local lifecycle smoke test and CI gate

**Files:** plugin/Dockerfile; _infra/docker-compose.yml; new _infra/docker-compose.test.yml; mock_api_server/src/api/mappings.go, controllers/models and tests; test-client startup/handlers; scripts/smoke.sh; .github/workflows/ci.yml; service .dockerignore files.

- [x] Supply the exact Nakama configuration path:

~~~dockerfile
COPY plugin/example/src/local.development.yml /nakama/data/local.yml
~~~

This path assumes the root build context introduced in Task 2.

- [x] Use health-based dependencies for PostgreSQL, mock API and Nakama. Test-client startup must have bounded readiness/retry logic.
- [x] Make mock state synchronized and persistent. Implement actual get/list/allocate/update/restart routes, metadata persistence, pagination and status transitions.
- [x] Reject malformed requests in the mock; avoid masking missing fields by returning an unconditional success fixture.
- [x] Submit one matchmaking ticket per client. Current startup calls both AddToMatch and Join, each submitting a ticket.
- [x] Create a bounded smoke scenario with two clients: one allocation, correct metadata, callback notifications with IP/port, stored session visibility, lifecycle update, restart and unauthorized-call rejection.
- [x] Make test Compose resources disposable: separate project/ports/volumes, no fixed container names, cleanup only of that test project's resources. Never delete the user's development volume.
- [x] Exclude credentials, .env files and unrelated artifacts from build contexts.
- [x] Extend the mock smoke checks to verify storage visibility before the callback, invalid readiness responses, callback timeout/duplicate completion, and recovery after a missed lifecycle update. Use deterministic fixtures for multi-page reconciliation and races; do not make normal CI wait for production polling intervals.
- [x] Add focused Nakama metrics for allocation outcomes/duration and reconciliation failures; test emissions against success/error paths. Use bounded operation/status labels. Keep correlation IDs in redacted structured logs and avoid user, session or instance IDs as metric labels.
- [x] Final CI runs root tests with race detection, all support-module checks, dependency cleanliness, Linux plugin build/load, clean-consumer module verification, and the smoke script.
- [x] Inspect required-check settings separately; a workflow file alone does not enforce branch protection.

**Validation:**

~~~sh
go test -mod=readonly -race -count=1 ./...
go vet -mod=readonly ./...
(cd plugin/example/src && go test -mod=readonly ./...)
(cd mock_api_server/src && go test -mod=readonly ./...)
(cd nakama_test_client/src && go test -mod=readonly ./...)
docker compose -f _infra/docker-compose.yml config --quiet
sh scripts/smoke.sh
~~~

**Acceptance:** Clean checkout builds, plugin loads, two-client lifecycle succeeds, negative paths fail correctly, and CI reproduces the checks without dependency rewrites. Suggested commit: test: verify complete local fleet lifecycle.

## Task 9: Verify installation, correct docs, and prepare the release

**Files:** readme.md; nakama_test_client/src/readme.md; example source/config; docs/compatibility.md; docs/release-checklist.md; new docs/reviews/heroiclabs-resolution.md; online-docs.md (authorized partner draft); docs/drafts/2026-09-22-online-docs.previous.md (preserved original).

- [x] Compile the real guide example in an external consumer project: correct module/config/runtime imports, InitModule signature, matchmaker accessors, two-return Create call, connection fields and notification payload type.
- [x] Document exact tested library/Nakama/common/Go/image versions. Retain a legacy migration note; do not claim all newer versions work.
- [x] Correct config keys, filesystem paths, container log commands, query field names and build directories.
- [x] Fix local status descriptions: 2 offline, 3 starting, 4 online, 5 allocated. Keep website changes limited to differences actually still present there.
- [x] Describe healthcheck as handler liveness unless real provider-readiness checks are added and tested. Do not claim it measures allocation performance.
- [x] Explain supported Join/reservation behavior, metadata ownership, async completion guarantees and game-server authentication.
- [x] Document the responsibilities of One API/Arcus, the headless server, the adapter and Nakama separately. Clarify that i3D Create allocates an existing online server and Delete restarts it. Describe connection/disconnection reporting, missed-update recovery and shutdown; include a minimal server-call example without adding a new Unity SDK to this release.
- [x] Exercise documented JSON RPC payloads through the real decoder/handler and assert resulting player counts and metadata. Check exact field names such as player_count; compiling Go snippets cannot detect JSON wire-format drift.
- [x] Compile both public installation and matchmaker snippets in CI. The peer guides contain outdated Create examples despite newer repository pins; avoid publishing the same drift.
- [x] Prepare a partner resolution table: HL-01 public paths and updated guide; HL-02 exact compatibility plus tested plugin; HL-03 clean installation evidence. Draft the corresponding guide patch/message for review, without sending it automatically.
- [ ] Select a release version consistent with repository tags and the public API migration. Validate the candidate before publishing, then run scripts/check-install.sh against the actual published version using GOWORK=off and a fresh module cache.
- [ ] Identify the live artifact/configuration and prepare its migration and rollback. Retain the prior plugin image and config.
- [ ] Run an authorized staging test against a designated i3D fleet: metadata, readiness, notifications, cleanup, transient failures and state reconciliation.
- [ ] Publish/deploy only after the concrete candidate and release evidence are reviewed and authorized.

## Definition of done

- [x] Every supplied Heroic Labs comment has a mapped change and evidence.
- [x] go get from the actual public GitHub module works in a clean consumer without cloning or replace.
- [x] The example and library compile and the plugin loads on the declared Nakama/common/Go tuple.
- [x] Existing tests plus permanent regression tests pass; concurrency cases pass -race.
- [x] Authenticated players cannot mutate fleet lifecycle through server RPCs.
- [x] Allocation metadata, update paths, response validation, pagination and persisted capacity work.
- [x] Async allocation/notifications survive normal hook completion and remain bounded by timeout/shutdown.
- [x] Configuration sources agree; OAuth uses its configured URL; failure/cancellation/retries are controlled.
- [x] Session-state ownership, concurrency and reservation limitations are explicit and tested.
- [x] Readiness follows i3D's allocation contract; callback timeout/duplicate handling and storage visibility are tested, with process-restart limitations documented.
- [x] Reconciliation handles all pages, protects concurrent/new allocations, and never removes records after a partial scan or restarts servers as cache cleanup.
- [x] Local end-to-end smoke verification and CI checks pass.
- [x] Docs are derived from compiling examples and actual supported versions.
- [ ] Candidate installation and live deployment identity are separately verified.
- [x] No credentials or .env files enter commits, artifacts or build contexts.

## Questions to resolve during execution

These do not prevent writing this plan or fixing the reproduced defects:
- What commit/image is deployed, on what Nakama version, and which RPC authentication path do its game servers use?
- Are Join/backfill and OAuth exercised in production, or mainly Create with static tokens?
- Does Heroic Labs require full expiring reservations from this adapter, or accept a documented provider limitation?
- Is a maintained 3.26 backport needed while consumers migrate, or is an identifiable legacy revision sufficient?

Partner handoff is settled: the user will send the new online-docs.md draft to Heroic Labs. Complete its publication checks against the actual release before describing the target as supported.

## Execution evidence and justified refinements

The ordered PRs and exact verification commands/results are recorded in ../../release-checklist.md. Native runtime tests refined two assumptions from planning: conditional-delete errors differ from write conflicts, and sort expressions must use value.player_count/value.create_time. The smoke harness forced and verified both against Nakama 3.41.0/PostgreSQL. Two clients also verified storage before notifications, one allocation, trusted unwrapped lifecycle payloads, missed-update recovery, invalid readiness and timeout.

Join remains the explicitly limited local-admission contract rather than a new expiring reservation service. Task 9 is complete for candidate installation, compiled documentation and partner preparation; live deployment identity, staging, tagging, merging and publication remain the external release gates listed above. No tags, merges, partner messages or deployments were performed.
