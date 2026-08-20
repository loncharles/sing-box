# Fork rebase log — sing-box

Running record of upstream rebase events. Each entry documents what upstream
tag/commit the fork was re-grafted onto, what conflicts required manual
resolution, and how the ported patches behaved.

## 2026-08-20 — Rebase feat/routing-mark-auto-redirect-coexistence → rebase/v1.14.0-beta.8

**Old base**: `553cfa1` (upstream "Bump version" root — ~1.13.x era squashed initial import)

**New base**: `v1.14.0-beta.8` (`0c23cdb`). Chosen because Leadaxe/sing-box-lx
targets that base for its SPEC 019 balancer patch we intend to port; taking
the same base minimizes port surface. No stable `v1.14.0` exists on upstream
at rebase time — beta.8 is the latest 1.14 tag.

**Old branch**: `feat/routing-mark-auto-redirect-coexistence` — kept in place,
not touched. This is the branch the fabric-gateway Dockerfile currently
clones from (`SINGBOX_BRANCH=feat/routing-mark-auto-redirect-coexistence`).
**New branch**: `rebase/v1.14.0-beta.8` — new work lands here. Dockerfile
will be redirected once this branch is pushed.

**Local commits carried over**:
- `fa983a3` (author Lon Lundgren) — Allow routing_mark coexistence with
  auto_redirect. Re-applied as `ba6ab6c` on the rebase branch.
- `8375c28` (same author) — Make tor outbound startup non-blocking.
  Re-applied as `4550791`.

**Cherry-pick conflicts**:
- `go.sum` — resolved with `--theirs` + `go mod tidy` (auto-regenerated,
  no code intent to preserve).
- Code files (`adapter/network.go`, `common/dialer/default.go`,
  `route/network.go`, `protocol/tor/outbound.go`) auto-merged cleanly.

**Platform build fix** (squashed into commit 1 as part of this rebase):
The original patch called `syscall.SetsockoptInt(..., syscall.SO_MARK, ...)`
directly in `setMarkWrapper` and `AutoRedirectOutputMarkFunc`. `syscall.SO_MARK`
is undefined on Darwin/Windows, breaking cross-platform builds. Fix: extract
the raw-conn OR-with-existing logic into build-tagged helpers
(`mark_or_linux.go` — real impl; `mark_or_other.go` — nil stub) in both
`common/dialer/` and `route/`. Call sites check the helper for nil and fall
back to `control.RoutingMark(mark)` (the sing common wrapper, which itself
uses the same platform partition upstream). Squashed via `git commit --fixup`
+ `git rebase -i --autosquash`.

**go.mod**: `replace github.com/sagernet/sing-tun => github.com/loncharles/sing-tun v0.8.10-0.20260820184418-82f3217956aa` — the pseudo-version references the pushed `rebase/v0.8.12-dev` HEAD (`82f3217`). Historical note: during the rebase work the replace directive briefly pointed at a local path `/Users/lon/workspace/sing-tun-fork` while sing-tun's rebase branch was unpushed; the flip to the pseudo-version landed in commit `677971b`.

**Build verification**: `go build ./cmd/sing-box/` clean on Darwin and
cross-compiled Linux amd64, with the fabric-gateway build tag set
(`with_gvisor,with_quic,with_dhcp,with_wireguard,with_utls,with_acme,with_clash_api,with_tailscale,with_ccm,with_ocm`).

**Test suite**: `go test ./...` — 3 tests failed in `common/tlsfragment`,
all `dial tcp 1.1.1.1:443: connect: operation timed out`. These are
network-dependent integration tests that require reach to Cloudflare
TLS. The Mac executing the rebase couldn't reach `1.1.1.1:443` due to
its home-network path (also confirmed unreachable via bare curl/nc).
The failing tests do not touch any file our patches modify (`conn_test.go`
in `common/tlsfragment/`). Left as an environment-conditional failure to
resolve in CI or on a machine with unfiltered 1.1.1.1 access.

**Result commits (top to bottom on `rebase/v1.14.0-beta.8`)**:
- `677971b` — go.mod: point sing-tun replace at pushed rebase/v0.8.12-dev
- `e07398d` — docs: record SPEC-019 balancer commits in fork ledger
- `1c19b87` — [lx-port SPEC-019] urltest: balancer tests (26 balancer + 10 pool health)
- `3579f86` — [lx-port SPEC-019] urltest: balancer core + Dial/Listen integration
- `b4ab353` — [lx-port SPEC-019] urltest: options + constants for balancer
- `ba01096` — docs: initial fork ledger (this file's earliest revision)
- `4550791` — Make tor outbound startup non-blocking
- `ba6ab6c` — Allow routing_mark coexistence with auto_redirect
- (base) `0c23cdb` — upstream v1.14.0-beta.8 Bump version
