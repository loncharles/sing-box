package group

import (
	"context"
)

// invalidateReachability is a stub for lx-port SPEC-020 (reachability tracking),
// which we do NOT port in this codebase. In the source project (Leadaxe/sing-box-lx),
// this call notifies the router that the active routing tree changed so its cached
// reachable set can be recomputed. Here it is a no-op, so the balancer core file
// (which is copied verbatim from lx) can call it without dragging SPEC-020's
// adapter interfaces into our tree.
func invalidateReachability(ctx context.Context) {
	_ = ctx
}
