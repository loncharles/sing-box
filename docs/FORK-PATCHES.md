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
| 3 | lx-port SPEC-019 | urltest: options + constants for balancer | `URLTestOutboundOptions.Mode` (least_test | round_robin), `Balancer` sub-struct (Pool + PoolTolerance + StickyHash), `PassiveCheck` bool. Constants for the two modes, the sticky-hash component names, and the default pool size. Ported from Leadaxe/sing-box-lx. |
| 4 | lx-port SPEC-019 | urltest: balancer core + Dial/Listen integration | `protocol/group/urltest_balance_lx.go` (426 lines, verbatim from lx) — balancer struct, slot pool, FNV-64a stickiness, planFirstLivePool / planTolerantPool. `protocol/group/reachability_lx.go` — no-op stub for `invalidateReachability`. `protocol/group/urltest.go` — URLTest and URLTestGroup gain balancer + passiveCheck fields; DialContext and ListenPacket branch on balancer; balancePool / rebuildPool / seedPool / passive helpers added. SPEC 020 reachability, SPEC 050 zombie fix, SPEC 054 penalty failover were stripped during the port — they are unrelated features from the same lx source file. Ported from Leadaxe/sing-box-lx. |
| 5 | lx-port SPEC-019 | urltest: balancer tests (26 balancer + 10 pool health) | `urltest_balance_lx_test.go` (551 lines) + `urltest_pool_health_lx_test.go` (243 lines). One test omitted from the source: `TestPerformUpdateCheck_firstSelectionInvalidates` (exercises SPEC 020's `adapter.ReachabilityInvalidator`, not ported). Ported from Leadaxe/sing-box-lx. |
