// +build windows
package wss

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func ss2GetOptions(ss2 *ISpeechSynthesizer2, options *ISpeechSynthesizerOptions) (err error) {
	hr, _, _ := syscall.Syscall(
		ss2.VTable().GetOptions,
		2,
		uintptr(unsafe.Pointer(ss2)),
		uintptr(unsafe.Pointer(options)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
