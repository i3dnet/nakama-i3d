# Release candidate checklist — 22 September 2026

Implementation candidate: f9673c4415288fd3c740da609178fcdf365105ce, publicly resolved as v0.0.0-20260922154340-f9673c441528. Main was reviewed at f26ac9d83ebb1a70f492db70b69eacf842403cb6. The deployed artifact is unknown.

## Review stack

Merge in order after review; each PR is based on the preceding branch. Do not release intermediate packaging/security commits as the supported complete integration.

| PR | Change |
| --- | --- |
| [#8](https://github.com/i3dnet/nakama-i3d/pull/8) | Repair baseline fixtures and establish CI |
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

## Completed candidate evidence

- [x] All four modules pass readonly tests with race detection and vet.
- [x] Actual plugin build/load on Nakama 3.41.0/common 1.48.0/Go 1.27.1/protobuf 1.36.12: Linux ARM64 locally and Linux AMD64 in CI.
- [x] Two-client lifecycle smoke on both architectures: one matchmaking allocation, metadata, persisted storage before notifications, correct IP/port, trusted updates/restarts and rejection of player lifecycle RPCs.
- [x] Native PostgreSQL-backed StorageWriteRetry conflict, concurrent Join admission and rejected stale delete.
- [x] Actual storage-index sorting and continuation pages. Sort fields use value.player_count and value.create_time.
- [x] Provider multi-page reconciliation, missed termination without restart, invalid readiness and bounded timeout without allocation retry.
- [x] Unit regression coverage for callback lifetime/one terminal callback, failed provider pages, Create/Join/reallocation races, storage pagination boundaries, OAuth refresh and cancellation.
- [x] Fresh public Go module cache, no clone/replace/private GitLab credentials: scripts/check-install.sh at f9673c4 resolves the pseudo-version above and compiles the real example.
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
