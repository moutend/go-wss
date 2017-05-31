// +build windows
package wss

import (
	"reflect"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func sssGetMarkers(sss *ISpeechSynthesisStream, markers interface{}) (err error) {
	markersValue := reflect.ValueOf(markers).Elem()
	hr, _, _ := syscall.Syscall(
		sss.VTable().GetMarkers,
		2,
		uintptr(unsafe.Pointer(sss)),
		markersValue.Addr().Pointer(),
		0)

	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
