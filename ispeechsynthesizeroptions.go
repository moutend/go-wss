package wss

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type ISpeechSynthesizerOptions struct {
	ole.IInspectable
}

type ISpeechSynthesizerOptionsVtbl struct {
	ole.IInspectableVtbl
	GetIncludeWordBoundaryMetadata     uintptr
	PutIncludeWordBoundaryMetadata     uintptr
	GetIncludeSentenceBoundaryMetadata uintptr
	PutIncludeSentenceBoundaryMetadata uintptr
}

func (v *ISpeechSynthesizerOptions) VTable() *ISpeechSynthesizerOptionsVtbl {
	return (*ISpeechSynthesizerOptionsVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *ISpeechSynthesizerOptions) GetIncludeWordBoundaryMetadata(value *bool) (err error) {
	err = ssoGetIncludeWordBoundaryMetadata(v, value)
	return
}

func (v *ISpeechSynthesizerOptions) PutIncludeWordBoundaryMetadata(value bool) (err error) {
	err = ssoPutIncludeWordBoundaryMetadata(v, value)
	return
}

func (v *ISpeechSynthesizerOptions) GetIncludeSentenceBoundaryMetadata(value *bool) (err error) {
	err = ssoGetIncludeSentenceBoundaryMetadata(v, value)
	return
}

func (v *ISpeechSynthesizerOptions) PutIncludeSentenceBoundaryMetadata(value bool) (err error) {
	err = ssoPutIncludeSentenceBoundaryMetadata(v, value)
	return
}
