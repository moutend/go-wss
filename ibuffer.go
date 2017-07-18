package wss

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type iBuffer struct {
	ole.IInspectable
}

type iBufferVtbl struct {
	ole.IInspectableVtbl
	GetCapacity uintptr
	GetLength   uintptr
	PutLength   uintptr
}

func (v *iBuffer) VTable() *iBufferVtbl {
	return (*iBufferVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iBuffer) GetCapacity(capacity *uint32) error {
	return bGetCapacity(v, capacity)
}

func (v *iBuffer) GetLength(length *uint32) error {
	return bGetLength(v, length)
}

func (v *iBuffer) PutLength(length uint32) error {
	return bPutLength(v, length)
}
