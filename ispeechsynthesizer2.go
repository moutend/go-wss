package wss

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type ISpeechSynthesizer2 struct {
	ole.IInspectable
}

type ISpeechSynthesizer2Vtbl struct {
	ole.IInspectableVtbl
	GetOptions uintptr
}

func (v *ISpeechSynthesizer2) VTable() *ISpeechSynthesizer2Vtbl {
	return (*ISpeechSynthesizer2Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *ISpeechSynthesizer2) GetOptions(options *ISpeechSynthesizerOptions) (err error) {
	err = ss2GetOptions(v, options)
	return
}
