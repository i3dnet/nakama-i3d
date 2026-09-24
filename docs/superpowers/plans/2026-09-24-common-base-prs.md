# Common-base PR regrouping Implementation Plan

> **For agentic workers:** Use superpowers:executing-plans to implement this approved reorganization inline.

**Goal:** Replace the dependent PR stack with two non-overlapping PRs against the exact same main commit.
**Architecture:** Runtime, build configuration and executable verification remain together because they depend on the new public module and FleetManager interface. Partner documentation and its snippet compiler use a separate workflow, compiling the published runtime candidate until the runtime reaches main and compiling the checkout afterwards.
**Tech Stack:** Git/GitHub, Go 1.27.1, Nakama 3.41.0, Docker Compose, Python.
**Spec:** User's September 24 request: fewer PRs, distinct scopes, no overlaps, same base tree.

## Global Constraints

- Both branches start directly at main commit 9a20297372909469ffdc88e203f64ff761a40092.
- Preserve the reviewed source from candidate d5f5a8147f4aa4644808c688d722cb2227dd7d51.
- No shared changed paths between the two PRs; runtime owns CI, documentation owns a new guide workflow.
- Keep original PR branches and review discussions available. Close superseded PRs only after replacements and reviews are established.
- Request PatrickB1977 and Copilot Balanced on both replacements.
- Do not merge, tag, deploy, or send partner messages.
- Verify .env remains ignored and scan all staged additions for secrets before committing.

## Review Focus

- Losing fixes while collapsing history: compare every source blob against the reviewed candidate.
- Hidden dependencies between PRs: test the runtime alone and compile guide examples on the docs-only branch.
- Overlapping diffs: compare changed path sets and verify the branches merge cleanly.
- Stale install pins: resolve the new runtime commit publicly and compile the guide against that exact version.
- Skipped final integration checks: run both the lifecycle smoke and local guide compilation on the combined tree.

## Task 1: Runtime and verification

Files: all changed paths except .gitattributes, readme.md, online-docs.md, docs/**, scripts/check-docs.py.
Interface: publishes the root Go module; ordinary CI checks runtime/plugin/smoke/public consumer without requiring the documentation PR.
- [x] Apply the saved candidate diff to a new branch from the frozen main commit.
- [x] Remove only the guide-compiler invocation from .github/workflows/ci.yml.
- [x] Compare source blobs with the saved candidate; run all four module race suites and vet, plugin load, smoke cleanup and lifecycle smoke.
- [x] Scan the staged patch, commit, push, and create the runtime PR.

## Task 2: Documentation and guide verification

Files: .gitattributes, readme.md, online-docs.md, docs/**, scripts/check-docs.py, new .github/workflows/guide.yml.
Interface: uses the public runtime candidate without clone/replace on the docs-only tree; verifies local source once go.mod exists.
- [x] Apply the documentation paths to a separate branch from the same main commit.
- [x] Update active candidate pins and the release checklist; keep historical review evidence and the archived draft intact.
- [x] Add a guide workflow: compile local source when go.mod exists, otherwise compile the published candidate.
- [x] Compile the actual guide blocks on the documentation-only tree and verify archived draft hash.
- [ ] Scan staged additions, commit, push, and create the documentation PR.

## Task 3: Verify and replace the old stack

- [ ] Verify both branch fork points equal the frozen main commit and changed-path sets have an empty intersection.
- [ ] Combine the branches locally; verify source blobs are unchanged and run guide compilation against local runtime.
- [ ] Request Copilot Balanced and PatrickB1977 on both PRs; read complete review bodies and inline findings.
- [ ] Address actionable findings in the owning PR, retest, and resolve fixed threads.
- [ ] Confirm CI and replacement links, then close all superseded PRs without deleting their branches.
- [ ] Update release/review mapping and report the two review links.
