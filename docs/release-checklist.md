# Release candidate checklist — 24 September 2026

Implementation candidate: 60cea4930eae7c464d6f4342663a6985ca9f0615, publicly resolved as v0.0.0-20260924154401-60cea4930eae. The initial main review used f26ac9d83ebb1a70f492db70b69eacf842403cb6. Upstream main is now 9a20297 after #8 was merged; this work has not merged PRs. The deployed artifact is unknown.

## Review layout

Both replacement PRs start at the same main commit, 9a20297372909469ffdc88e203f64ff761a40092. Their changed files do not overlap. Review and merge the runtime and documentation scopes; do not merge the superseded stack in addition.

| PR | Scope |
| --- | --- |
| [Runtime #42](https://github.com/i3dnet/nakama-i3d/pull/42) | Public module, Nakama 3.41 migration, provider/storage/lifecycle fixes, build tooling and executable verification |
| [Documentation #43](https://github.com/i3dnet/nakama-i3d/pull/43) | Partner guide, user/reference documentation, historical evidence and standalone guide compilation |

The runtime PR includes the prior review fixes and the dependency updates from #36/#37. The documentation workflow compiles the pinned public runtime candidate while the root module is absent from main, then compiles the checked-out runtime once the root go.mod is present. No intermediate packaging or security-only revision is a supported release.

The [superseded stack map](reviews/2026-09-24-superseded-pr-stack.md) retains the original PR links. Finding-by-finding evidence remains available for [September 23](reviews/2026-09-23-copilot-resolution.md) and [September 24](reviews/2026-09-24-review-resolution.md). The [regrouping plan](superpowers/plans/2026-09-24-common-base-prs.md) records the file ownership and verification requirements.

## Completed candidate evidence

- [x] All four modules pass readonly tests with race detection and vet.
- [x] Actual plugin build/load on Nakama 3.41.0/common 1.48.0/Go 1.27.1/protobuf 1.36.12: Linux ARM64 locally and Linux AMD64 in CI.
- [x] Two-client lifecycle smoke on both architectures: one matchmaking allocation, metadata, persisted storage before notifications, correct IP/port, trusted updates/restarts and rejection of player lifecycle RPCs.
- [x] Native PostgreSQL-backed StorageWriteRetry conflict, concurrent Join admission and rejected stale delete.
- [x] Actual storage-index sorting, tied player counts with descending creation times, and continuation pages. Sort fields use value.player_count and value.create_time.
- [x] Provider multi-page reconciliation, missed termination without restart, invalid readiness, exactly one restart for a confirmed failed allocation, and bounded timeout without allocation retry.
- [x] Unit regression coverage for completed-result precedence, fresh persistence and cleanup contexts, cleanup error reporting, callback lifetime/one terminal callback, failed provider pages, Create/Join/reallocation races, storage pagination boundaries and bounded clock skew, OAuth refresh and cancellation.
- [x] Fresh public Go module cache, no clone/replace/private GitLab credentials: scripts/check-install.sh at 60cea49 resolves the pseudo-version above and compiles the real example.
- [x] Partner guide Go files and filter block compiled/tested in an isolated consumer through scripts/check-docs.py. CI recompiles the actual Markdown blocks.
- [x] Documented unwrapped JSON lifecycle payloads executed against Nakama's HTTP-key endpoint.
- [x] Ordinary plugin builds exclude the i3d_smoke test RPCs.
- [x] The exact prior online-docs draft is preserved at docs/drafts/2026-09-22-online-docs.previous.md (SHA-256 adc77c712904519e38efa7ead11df0f6913102e27bf2d0dc09f2b1ee4853cafa).

The active organization ruleset requires a PR, one approval, code-owner review, last-push approval and resolved conversations, with linear history and no force pushes on the default branch. Its rules do not require named CI checks. The legacy branch-protection endpoint returns “Branch not protected”; the organization ruleset still applies. No repository/ruleset settings were changed. Require the new CI checks before release through the normal owner review process.

## Balanced review follow-up

The runtime follow-up at 987fc15 addresses the three findings in [Copilot's review of #42](https://github.com/i3dnet/nakama-i3d/pull/42#pullrequestreview-5306278771):

| Finding | Fix and verification |
| --- | --- |
| [Provider-only allocations lack recoverable capacity](https://github.com/i3dnet/nakama-i3d/pull/42#discussion_r4095147495) | Get/List/reconciliation return provider-only discovery without importing it. Update requires an existing session. Tests verify no storage writes/imports and refusal of Join. |
| [Canceled OAuth refresh leader fails live waiters](https://github.com/i3dnet/nakama-i3d/pull/42#discussion_r4095147568) | Live waiters share a replacement refresh; each waiter keeps its own cancellation. Deterministic tests cover leader cancellation/deadline, provider failures and a canceled waiter racing success. |
| [List cache records lose cleanup ownership](https://github.com/i3dnet/nakama-i3d/pull/42#discussion_r4095147617) | Unknown records are not cached. A replaced generation retains ownership while invalidated or is conditionally retired. Tests cover mixed pages, cleanup, repeated reads, older provider generations and a delayed Update that must not revive admission. |

The full root race suite, vet, real Nakama/PostgreSQL lifecycle smoke and fresh public consumer installation passed after these fixes. A follow-up at bb9a621 corrects the new timestamp regression to compare instants rather than time-zone representation; the full root race suite also passes with TZ=UTC. The guide's dependency pin and behavior documentation track the current candidate listed above. Fresh current-head Copilot Balanced review and CI remain separate review gates.

The [next runtime Balanced review](https://github.com/i3dnet/nakama-i3d/pull/42#pullrequestreview-5306571688) reported additional issues in expanded details despite showing “Findings: None.” The [documentation review](https://github.com/i3dnet/nakama-i3d/pull/43#pullrequestreview-5306598408) used the same presentation. These findings have no inline threads to resolve.

| Additional finding | Resolution |
| --- | --- |
| Reported player count exceeds local capacity | Validate against stored capacity before the provider call and again in the versioned mutation; reject over-capacity reports as INVALID_ARGUMENT. |
| One API retains the previous player count | Pass the reported count through the provider client and write numPlayers alongside metadata. |
| Shutdown grace is shorter than cleanup budget | Example and smoke Nakama configs allow 45 seconds; Compose allows 60 seconds before forced termination, exceeding the default 30-second cleanup budget. |
| Generated build-provisioning endpoint leaves an ID placeholder in the URL | Align the generated route and source schema placeholder; an HTTP-capture regression verifies the exact path. |
| Guide verifier omits the import-only block | Compile that block in its own Go file and reject unrecognized blocks. A temporary malformed import was accepted before the fix and rejected afterward. |
| Historical review record claims an obsolete active pin | Describe the prior candidate in the past tense and link to this current checklist. |

The additional runtime fixes are included in 60cea49. Root and mock-provider race suites and vet pass with TZ=UTC; the generated endpoint HTTP capture and real Nakama/PostgreSQL lifecycle smoke pass, including provider player-count propagation. Fresh public consumer installation resolves the candidate above and builds the real example without replace.

## Remaining release gates

- [ ] Complete human review and merge both replacement PRs. No PR has been merged by this work.
- [ ] Choose the first public root-module version. The repository has no tags; v0.1.0 is proposed, not created.
- [ ] Identify the live commit/image digest, Nakama version, runtime/process configuration source, application/fleet and headless-server authentication.
- [ ] Confirm whether a security backport is needed for an existing 3.26 deployment. A 3.41 plugin cannot be substituted into the old runtime.
- [ ] Validate on an explicitly designated staging i3D fleet: real One API/Arcus metadata delivery, readiness, endpoint assignment, authentication, transient failure and propagation delay.
- [ ] Tune timeouts and reconciliation grace to staging measurements. Two successful but incomplete provider scans can still falsely imply absence.
- [ ] Retain the previous runtime/plugin image and configuration, take the normal database backup, document the rollout owner and rollback steps.
- [ ] Review stored-session compatibility before rolling back: old records remain readable, but older code ignores new admission/version semantics and is not safe for concurrent mixed-version writers.
- [ ] Publish an approved tag, rerun scripts/check-install.sh against it, replace the candidate pin in online-docs.md/readme.md and complete the normal deployment approval.
- [ ] User sends the final partner draft. No external partner message has been sent.

Live deployment identity and staging access are external gates. Local mocks and a successfully loaded plugin do not establish live provider readiness or production safety.
