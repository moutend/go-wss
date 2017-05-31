// +build windows
package wss

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func ssoGetIncludeWordBoundaryMetadata(sso *ISpeechSynthesizerOptions, value *bool) (err error) {
	hr, _, _ := syscall.Syscall(
		sso.VTable().GetIncludeWordBoundaryMetadata,
		2,
		uintptr(unsafe.Pointer(sso)),
		uintptr(unsafe.Pointer(value)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func ssoPutIncludeWordBoundaryMetadata(sso *ISpeechSynthesizerOptions, value bool) (err error) {
	var boolValue uint32

	if value {
		boolValue = 1
	}
	hr, _, _ := syscall.Syscall(
		sso.VTable().PutIncludeWordBoundaryMetadata,
		2,
		uintptr(unsafe.Pointer(sso)),
		uintptr(boolValue),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func ssoGetIncludeSentenceBoundaryMetadata(sso *ISpeechSynthesizerOptions, value *bool) (err error) {
	hr, _, _ := syscall.Syscall(
		sso.VTable().GetIncludeSentenceBoundaryMetadata,
		2,
		uintptr(unsafe.Pointer(sso)),
		uintptr(unsafe.Pointer(value)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func ssoPutIncludeSentenceBoundaryMetadata(sso *ISpeechSynthesizerOptions, value bool) (err error) {
	var boolValue uint32

	if value {
		boolValue = 1
	}
	hr, _, _ := syscall.Syscall(
		sso.VTable().PutIncludeSentenceBoundaryMetadata,
		2,
		uintptr(unsafe.Pointer(sso)),
		uintptr(boolValue),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
