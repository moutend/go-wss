// +build windows
package wss

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func bbaBuffer(bba *IBufferByteAccess, bufPtr **byte) (err error) {
	hr, _, _ := syscall.Syscall(
		bba.VTable().Buffer,
		2,
		uintptr(unsafe.Pointer(bba)),
		uintptr(unsafe.Pointer(bufPtr)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
