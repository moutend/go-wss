// +build windows
package wss

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func viGetDisplayName(vi *IVoiceInformation, name *ole.HString) (err error) {
	hr, _, _ := syscall.Syscall(
		vi.VTable().GetDisplayName,
		2,
		uintptr(unsafe.Pointer(vi)),
		uintptr(unsafe.Pointer(name)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}

	return
}

func viGetId(vi *IVoiceInformation, id *ole.HString) (err error) {
	hr, _, _ := syscall.Syscall(
		vi.VTable().GetId,
		2,
		uintptr(unsafe.Pointer(vi)),
		uintptr(unsafe.Pointer(id)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}

	return
}

func viGetLanguage(vi *IVoiceInformation, language *ole.HString) (err error) {
	hr, _, _ := syscall.Syscall(
		vi.VTable().GetLanguage,
		2,
		uintptr(unsafe.Pointer(vi)),
		uintptr(unsafe.Pointer(language)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}

	return
}

func viGetDescription(vi *IVoiceInformation, description *ole.HString) (err error) {
	hr, _, _ := syscall.Syscall(
		vi.VTable().GetDescription,
		2,
		uintptr(unsafe.Pointer(vi)),
		uintptr(unsafe.Pointer(description)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}

	return
}

func viGetGender(vi *IVoiceInformation, gender *VoiceGender) (err error) {
	hr, _, _ := syscall.Syscall(
		vi.VTable().GetGender,
		2,
		uintptr(unsafe.Pointer(vi)),
		uintptr(unsafe.Pointer(gender)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}

	return
}
