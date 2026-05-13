//go:build linux

package route

import (
	"syscall"

	"github.com/sagernet/sing/common/control"
)

// orAutoRedirectMark builds a control.Func that OR's the auto_redirect output
// mark onto any existing SO_MARK bits, so routing_mark and auto_redirect can
// coexist on the same socket. Linux only; other platforms get a no-op stub.
func orAutoRedirectMark(mark uint32) control.Func {
	return func(network, address string, conn syscall.RawConn) error {
		return control.Raw(conn, func(fd uintptr) error {
			current, _ := syscall.GetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_MARK)
			return syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_MARK, current|int(mark))
		})
	}
}
