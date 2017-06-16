package wss

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IVectorView struct {
	ole.IInspectable
}

type IVectorViewVtbl struct {
	ole.IInspectableVtbl
	GetAt   uintptr
	GetSize uintptr
	IndexOf uintptr
	GetMany uintptr
}

func (v *IVectorView) VTable() *IVectorViewVtbl {
	return (*IVectorViewVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IVectorView) GetAt(index uint16, voiceInformation **IVoiceInformation) (err error) {
	err = vvGetAt(v, index, voiceInformation)
	return
}

func (v *IVectorView) GetSize(size *uint16) (err error) {
	err = vvGetSize(v, size)
	return
}

func (v *IVectorView) IndexOf(voiceInformation *IVoiceInformation, index *uint16, found *bool) (err error) {
	err = vvIndexOf(v, voiceInformation, index, found)
	return
}
