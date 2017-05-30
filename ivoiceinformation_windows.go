// +build windows
package wss

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func infoGetDisplayName(info *IVoiceInformation, name *ole.HString) (err error) {
	hr, _, _ := syscall.Syscall(
		info.VTable().GetDisplayName,
		2,
		uintptr(unsafe.Pointer(info)),
		uintptr(unsafe.Pointer(name)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}

	return
}

func infoGetId(info *IVoiceInformation, id *ole.HString) (err error) {
	hr, _, _ := syscall.Syscall(
		info.VTable().GetId,
		2,
		uintptr(unsafe.Pointer(info)),
		uintptr(unsafe.Pointer(id)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}

	return
}

func infoGetLanguage(info *IVoiceInformation, language *ole.HString) (err error) {
	hr, _, _ := syscall.Syscall(
		info.VTable().GetLanguage,
		2,
		uintptr(unsafe.Pointer(info)),
		uintptr(unsafe.Pointer(language)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}

	return
}

func infoGetDescription(info *IVoiceInformation, description *ole.HString) (err error) {
	hr, _, _ := syscall.Syscall(
		info.VTable().GetDescription,
		2,
		uintptr(unsafe.Pointer(info)),
		uintptr(unsafe.Pointer(description)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}

	return
}

func infoGetGender(info *IVoiceInformation, gender *ole.HString) (err error) {
	hr, _, _ := syscall.Syscall(
		info.VTable().GetGender,
		2,
		uintptr(unsafe.Pointer(info)),
		uintptr(unsafe.Pointer(gender)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}

	return
}
