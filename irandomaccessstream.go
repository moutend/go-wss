package wss

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IRandomAccessStream struct {
	ole.IInspectable
}

type IRandomAccessStreamVtbl struct {
	ole.IInspectableVtbl
	GetSize           uintptr
	PutSize           uintptr
	GetInputStreamAt  uintptr
	GetOutputStreamAt uintptr
	GetPosition       uintptr
	Seek              uintptr
	CloneStream       uintptr
	GetCanRead        uintptr
	GetCanWrite       uintptr
}

func (v *IRandomAccessStream) VTable() *IRandomAccessStreamVtbl {
	return (*IRandomAccessStreamVtbl)(unsafe.Pointer(v.RawVTable))
}
