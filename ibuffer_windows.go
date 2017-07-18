// +build windows

package wss

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func bGetCapacity(b *iBuffer, capacity *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		b.VTable().GetCapacity,
		2,
		uintptr(unsafe.Pointer(b)),
		uintptr(unsafe.Pointer(capacity)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func bGetLength(b *iBuffer, length *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		b.VTable().GetLength,
		2,
		uintptr(unsafe.Pointer(b)),
		uintptr(unsafe.Pointer(length)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func bPutLength(b *iBuffer, length uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		b.VTable().PutLength,
		2,
		uintptr(unsafe.Pointer(b)),
		uintptr(length),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
