package wss

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IBufferByteAccess struct {
	ole.IUnknown
}

type IBufferByteAccessVtbl struct {
	ole.IUnknownVtbl
	Buffer uintptr
}

func (v *IBufferByteAccess) VTable() *IBufferByteAccessVtbl {
	return (*IBufferByteAccessVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IBufferByteAccess) Buffer(bufPtr **byte) (err error) {
	err = bbaBuffer(v, bufPtr)
	return
}
