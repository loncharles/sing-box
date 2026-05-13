//go:build !linux

package dialer

import "github.com/sagernet/sing/common/control"

// orRoutingMark on non-Linux platforms is a no-op — routing_mark is only
// meaningful on Linux via SO_MARK / nftables.
func orRoutingMark(mark uint32) control.Func {
	return nil
}
