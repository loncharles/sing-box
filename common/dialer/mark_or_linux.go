//go:build linux

package dialer

import (
	"syscall"

	"github.com/sagernet/sing/common/control"
)

// orRoutingMark sets SO_MARK on the socket by OR'ing the given mark into
// any existing mark bits, so mark writes from different layers (auto_redirect
// + routing_mark) can coexist. Linux only; the non-Linux stub is a no-op.
func orRoutingMark(mark uint32) control.Func {
	return func(network, address string, conn syscall.RawConn) error {
		return control.Raw(conn, func(fd uintptr) error {
			current, _ := syscall.GetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_MARK)
			return syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_MARK, current|int(mark))
		})
	}
}
