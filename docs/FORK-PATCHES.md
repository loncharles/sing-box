# Fork patches — sing-box

Ledger of every non-upstream commit on this fork. Each row corresponds to a
commit whose message ends with an `Origin:` trailer. Grep-able (see
`git log --grep 'Origin:'`) and machine-readable.

Conventions:
- `Origin: local` — commit authored here, not derived from another repo.
- `Origin: <name>-port <spec> <sha>` — commit ported from another fork; sha
  refers to the source commit.
- Commit subject prefix `[local]` or `[<name>-port <spec>]` matches the
  Origin: trailer for `git log --oneline` clarity.

## Current patch stack (on `rebase/v1.14.0-beta.8`)

| Order | Origin | Purpose | Notes |
|-------|--------|---------|-------|
| 1 | local | Allow routing_mark coexistence with auto_redirect | The core patch changes `AutoRedirectOutputMarkFunc` (route/network.go) and `setMarkWrapper` (common/dialer/default.go) from write-once to read-OR-write on SO_MARK, so a routing_mark bit-mask can coexist with auto_redirect on the same socket. Adds `AutoRedirectMarkMask` to the `NetworkManager` interface. The raw SO_MARK usage was extracted into build-tagged helpers (`mark_or_linux.go` / `mark_or_other.go` in both `common/dialer` and `route`) so the code builds on non-Linux platforms (Darwin, Windows) with the mark logic no-op'd — sing/common/control's `RoutingMark` handles the platform partition upstream, so we mirror the same convention. |
| 2 | local | Make tor outbound startup non-blocking | `tor.Outbound.Start()` was calling `EnableNetwork(ctx, true)` (blocks until bootstrap PROGRESS=100). If bootstrap fails (dead detour), that call never returns and the whole outbound-startup loop deadlocks — no inbounds initialize either. Change to `EnableNetwork(ctx, false)` and rely on tor's SOCKS listener being ready at process spawn. Individual pre-bootstrap connections fail; the gateway does not deadlock. |
