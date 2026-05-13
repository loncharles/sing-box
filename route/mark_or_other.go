//go:build !linux

package route

import "github.com/sagernet/sing/common/control"

// orAutoRedirectMark on non-Linux platforms is a no-op — auto_redirect is
// only meaningful on Linux via SO_MARK / nftables.
func orAutoRedirectMark(mark uint32) control.Func {
	return nil
}
