// +build windows
package wss

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func vvGetAt(vv *IVectorView, index uint16, voiceInformation **IVoiceInformation) (err error) {
	hr, _, _ := syscall.Syscall(
		vv.VTable().GetAt,
		3,
		uintptr(unsafe.Pointer(vv)),
		uintptr(uint32(index)),
		uintptr(unsafe.Pointer(voiceInformation)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func vvGetSize(vv *IVectorView, size *uint16) (err error) {
	hr, _, _ := syscall.Syscall(
		vv.VTable().GetSize,
		2,
		uintptr(unsafe.Pointer(vv)),
		uintptr(unsafe.Pointer(size)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func vvIndexOf(vv *IVectorView, voiceInformation *IVoiceInformation, index *uint16, found *bool) (err error) {
	hr, _, _ := syscall.Syscall6(
		vv.VTable().IndexOf,
		4,
		uintptr(unsafe.Pointer(vv)),
		uintptr(unsafe.Pointer(voiceInformation)),
		uintptr(unsafe.Pointer(index)),
		uintptr(unsafe.Pointer(found)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func vvGetMany(vv *IVectorView) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
