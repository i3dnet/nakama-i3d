# Superseded PR stack — 24 September 2026

This is the historical review/merge map from before the common-base regrouping. It is retained for discussion provenance, not as current merge instructions. See the [current release checklist](../release-checklist.md) for the active PRs. Original branches and review discussions remain available.

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
