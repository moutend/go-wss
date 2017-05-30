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
	get_IncludeWordBoundaryMetadata     uintptr
	put_IncludeWordBoundaryMetadata     uintptr
	get_IncludeSentenceBoundaryMetadata uintptr
	put_IncludeSentenceBoundaryMetadata uintptr
}

func (v *ISpeechSynthesizerOptions) VTable() *ISpeechSynthesizerOptionsVtbl {
	return (*ISpeechSynthesizerOptionsVtbl)(unsafe.Pointer(v.RawVTable))
}
