// +build windows
package wss

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func ivsGetAllVoices(ivs *IInstalledVoicesStatic, voices *IVoiceInformation) (err error) {
	hr, _, _ := syscall.Syscall(
		ivs.VTable().GetAllVoices,
		2,
		uintptr(unsafe.Pointer(ivs)),
		uintptr(unsafe.Pointer(voices)),
		0)

	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func ivsGetDefaultVoice(ivs *IInstalledVoicesStatic, voice *IVoiceInformation) (err error) {
	hr, _, _ := syscall.Syscall(
		ivs.VTable().GetDefaultVoice,
		2,
		uintptr(unsafe.Pointer(ivs)),
		uintptr(unsafe.Pointer(voice)),
		0)

	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
