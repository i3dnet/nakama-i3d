# Release candidate checklist — 24 September 2026

Implementation candidate: 983f46c78e19656d86a3d7b6b069b3489a711cd7, publicly resolved as v0.0.0-20260924105049-983f46c78e19. The initial main review used f26ac9d83ebb1a70f492db70b69eacf842403cb6. Upstream main is now 9a20297 after #8 was merged; this work has not merged PRs. The deployed artifact is unknown.

## Review stack

Merge in order after review; each PR is based on the preceding branch. Do not release intermediate packaging/security commits as the supported complete integration.

| PR | Change |
| --- | --- |
| [#8](https://github.com/i3dnet/nakama-i3d/pull/8) | Repair baseline fixtures and establish CI; merged upstream |
| [#9](https://github.com/i3dnet/nakama-i3d/pull/9) | Publish the library at the repository root |
| [#10](https://github.com/i3dnet/nakama-i3d/pull/10) | Nakama 3.41 compatibility and bounded async allocation |
| [#11](https://github.com/i3dnet/nakama-i3d/pull/11) | Trusted-server lifecycle RPC authorization |
| [#12](https://github.com/i3dnet/nakama-i3d/pull/12) | Provider request bodies, routes, response validation and filters |
| [#13](https://github.com/i3dnet/nakama-i3d/pull/13) | Storage decoding, numeric capacity and cursors |
| [#14](https://github.com/i3dnet/nakama-i3d/pull/14) | Versioned admission and metadata ownership |
| [#15](https://github.com/i3dnet/nakama-i3d/pull/15) | One validated configuration contract |
| [#16](https://github.com/i3dnet/nakama-i3d/pull/16) | Synchronized OAuth and safe, bounded retries |
| [#17](https://github.com/i3dnet/nakama-i3d/pull/17) | Complete-scan reconciliation and version guards |
| [#18](https://github.com/i3dnet/nakama-i3d/pull/18) | Fleet metrics and redacted HTTP telemetry |
| [#19](https://github.com/i3dnet/nakama-i3d/pull/19) | Stateful provider mock and request validation |
| [#20](https://github.com/i3dnet/nakama-i3d/pull/20) | Real two-client lifecycle/native storage smoke |
| [#23](https://github.com/i3dnet/nakama-i3d/pull/23) | This guide, compiled snippets and release evidence |
| [#24](https://github.com/i3dnet/nakama-i3d/pull/24) | Preserve allocations and storage updates made during reconciliation pagination |
| [#25](https://github.com/i3dnet/nakama-i3d/pull/25) | Verify the public Go resolver and checksums |
| [#26](https://github.com/i3dnet/nakama-i3d/pull/26) | Copy metadata without numeric loss |
| [#27](https://github.com/i3dnet/nakama-i3d/pull/27) | Validate provider identity, endpoints and truncated reads |
| [#28](https://github.com/i3dnet/nakama-i3d/pull/28) | Atomic provider pages and generation-safe cleanup |
| [#29](https://github.com/i3dnet/nakama-i3d/pull/29) | Metadata map keys, opaque state and unsupported references |
| [#30](https://github.com/i3dnet/nakama-i3d/pull/30) | Bounded clock-skew protection and accurate cleanup wording |
| [#31](https://github.com/i3dnet/nakama-i3d/pull/31) | Mock filter, allocation-metadata and cursor contracts |
| [#32](https://github.com/i3dnet/nakama-i3d/pull/32) | Defensive skew minimum for direct configuration |
| [#33](https://github.com/i3dnet/nakama-i3d/pull/33) | Path-prefix telemetry and terminal metric coverage |
| [#34](https://github.com/i3dnet/nakama-i3d/pull/34) | Secondary index ordering and resilient smoke cleanup |
| [#35](https://github.com/i3dnet/nakama-i3d/pull/35) | Archive newline guard, dotenv redaction and September 23 review evidence |
| [#38](https://github.com/i3dnet/nakama-i3d/pull/38) | All four modules and build/test tooling on Go 1.27.1; support dependency updates |
| [#39](https://github.com/i3dnet/nakama-i3d/pull/39) | Independent persistence deadline, completed-result precedence, failed-allocation cleanup and fallback warning |
| [#40](https://github.com/i3dnet/nakama-i3d/pull/40) | Current candidate pin, direct client dependency classification and September 24 review evidence |
| [#41](https://github.com/i3dnet/nakama-i3d/pull/41) | Preserve escaped telemetry path segments and verify full review-summary follow-ups |

PRs [#36](https://github.com/i3dnet/nakama-i3d/pull/36) and [#37](https://github.com/i3dnet/nakama-i3d/pull/37) are independent Dependabot changes against legacy main. Their own module, CI and Docker toolchains are aligned; their dependency updates are also incorporated in #38. They are not additional steps in this stack, and do not establish support for loading the new plugin into the legacy Nakama runtime.

Finding-by-finding evidence is recorded for [September 23](reviews/2026-09-23-copilot-resolution.md) and [September 24](reviews/2026-09-24-review-resolution.md).

## Completed candidate evidence

- [x] All four modules pass readonly tests with race detection and vet.
- [x] Actual plugin build/load on Nakama 3.41.0/common 1.48.0/Go 1.27.1/protobuf 1.36.12: Linux ARM64 locally and Linux AMD64 in CI.
- [x] Two-client lifecycle smoke on both architectures: one matchmaking allocation, metadata, persisted storage before notifications, correct IP/port, trusted updates/restarts and rejection of player lifecycle RPCs.
- [x] Native PostgreSQL-backed StorageWriteRetry conflict, concurrent Join admission and rejected stale delete.
- [x] Actual storage-index sorting, tied player counts with descending creation times, and continuation pages. Sort fields use value.player_count and value.create_time.
- [x] Provider multi-page reconciliation, missed termination without restart, invalid readiness, exactly one restart for a confirmed failed allocation, and bounded timeout without allocation retry.
- [x] Unit regression coverage for completed-result precedence, fresh persistence and cleanup contexts, cleanup error reporting, callback lifetime/one terminal callback, failed provider pages, Create/Join/reallocation races, storage pagination boundaries and bounded clock skew, OAuth refresh and cancellation.
- [x] Fresh public Go module cache, no clone/replace/private GitLab credentials: scripts/check-install.sh at 983f46c resolves the pseudo-version above and compiles the real example.
- [x] Partner guide Go files and filter block compiled/tested in an isolated consumer through scripts/check-docs.py. CI recompiles the actual Markdown blocks.
- [x] Documented unwrapped JSON lifecycle payloads executed against Nakama's HTTP-key endpoint.
- [x] Ordinary plugin builds exclude the i3d_smoke test RPCs.
- [x] The exact prior online-docs draft is preserved at docs/drafts/2026-09-22-online-docs.previous.md (SHA-256 adc77c712904519e38efa7ead11df0f6913102e27bf2d0dc09f2b1ee4853cafa).

The active organization ruleset requires a PR, one approval, code-owner review, last-push approval and resolved conversations, with linear history and no force pushes. Its rules do not require named CI checks. The legacy branch-protection endpoint returns “Branch not protected”; the organization ruleset still applies. No repository/ruleset settings were changed. Require the new CI checks before release through the normal owner review process.

## Remaining release gates

- [ ] Complete human review and merge the stack. No PR has been merged by this work.
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
