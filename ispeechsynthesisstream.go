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
	get_Markers uintptr
}

func (v *ISpeechSynthesisStream) VTable() *ISpeechSynthesisStreamVtbl {
	return (*ISpeechSynthesisStreamVtbl)(unsafe.Pointer(v.RawVTable))
}
