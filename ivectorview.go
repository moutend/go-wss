package wss

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type iVectorView struct {
	ole.IInspectable
}

type iVectorViewVtbl struct {
	ole.IInspectableVtbl
	GetAt   uintptr
	GetSize uintptr
	IndexOf uintptr
	GetMany uintptr
}

func (v *iVectorView) VTable() *iVectorViewVtbl {
	return (*iVectorViewVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iVectorView) GetAt(index uint16, voiceInformation **iVoiceInformation) error {
	return vvGetAt(v, index, voiceInformation)
}

func (v *iVectorView) GetSize(size *uint16) error {
	return vvGetSize(v, size)
}

func (v *iVectorView) IndexOf(voiceInformation *iVoiceInformation, index *uint16, found *bool) error {
	return vvIndexOf(v, voiceInformation, index, found)
}
