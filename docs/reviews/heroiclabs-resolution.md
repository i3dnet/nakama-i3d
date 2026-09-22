# Heroic Labs feedback resolution

Candidate: 20a8631305f5f9adfa66d2869bd593230a760477. Partner draft: [online-docs.md](../../online-docs.md). This is a review candidate, not a tagged or deployed release.

| Feedback | Resolution | Evidence |
| --- | --- | --- |
| HL-01: old repository/import path in the online guide | Root module and imports now use github.com/i3dnet/nakama-i3d. The replacement guide has complete registration/matchmaking files, current config and lifecycle examples. | PR #9; scripts/check-docs.py compiles the actual draft. |
| HL-02: common 1.36 pins limited runtime compatibility | Upgrade all shared dependencies, FleetManager Create signature and images together to Nakama 3.41.0/common 1.48.0/Go 1.27.1/protobuf 1.36.12. Keep the 3.26 legacy revision identifiable. | PR #10; real Linux ARM64 and AMD64 plugin loading; PR #20 lifecycle smoke. |
| HL-03: go get does not provide an installable module | Reusable code lives under a public root go.mod. Local replace remains only for development of the separate example. | Fresh public resolution to v0.0.0-20260922152017-20a8631305f5; actual example compiles without cloning or replace. |

Additional corrections from the provider comparison:
- Preserve i3D's native allocation completion instead of copying a provider-specific READY webhook.
- Persist success before callback; manage accepted allocation and notification contexts independently of the matchmaking hook.
- Use server HTTP-key authentication for lifecycle RPCs and execute the exact player_count JSON payloads.
- Compile guide examples continuously to avoid one-return Create and runtime-version drift.
- Reconcile every provider page before absence cleanup, with versions and a grace/confirmation window.
- State Join's local admission limitations explicitly: no provider-issued token or expiring reservation.

The comparison with Heroic Labs' GameLift adapter and Edgegap, including pinned revisions and the reportable Edgegap documentation differences, remains in [the source review](2026-09-22-provider-integration-comparison.md). It is source analysis, not a claim that those providers were tested live.

Suggested accompanying message for the user to send:

> We've prepared a replacement i3D integration guide and implementation candidate. The library is now installable from the public repository root; the candidate is pinned to Nakama 3.41.0/common 1.48.0 and has passed real plugin loading and a two-client lifecycle smoke test. The draft also updates configuration, server authentication, callback lifetime and the documented Join limitations. Please review the attached online-docs.md. We'll replace the candidate commit with the approved release tag after review and i3D staging validation.

No message was sent automatically.
