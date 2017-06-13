package wss

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IBuffer struct {
	ole.IInspectable
}

type IBufferVtbl struct {
	ole.IInspectableVtbl
	GetCapacity uintptr
	GetLength   uintptr
	PutLength   uintptr
}

func (v *IBuffer) VTable() *IBufferVtbl {
	return (*IBufferVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IBuffer) GetCapacity(capacity *uint32) (err error) {
	err = bGetCapacity(v, capacity)
	return
}

func (v *IBuffer) GetLength(length *uint32) (err error) {
	err = bGetLength(v, length)
	return
}

func (v *IBuffer) PutLength(length uint32) (err error) {
	err = bPutLength(v, length)
	return
}
