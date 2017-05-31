package wss

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type ISpeechSynthesisStream struct {
	ole.IInspectable
}

type ISpeechSynthesisStreamVtbl struct {
	ole.IInspectableVtbl
	GetMarkers uintptr
}

func (v *ISpeechSynthesisStream) VTable() *ISpeechSynthesisStreamVtbl {
	return (*ISpeechSynthesisStreamVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *ISpeechSynthesisStream) GetMarker(markers interface{}) (err error) {
	err = sssGetMarkers(v, markers)
	return
}
