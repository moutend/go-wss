package wss

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type iBufferByteAccess struct {
	ole.IUnknown
}

type iBufferByteAccessVtbl struct {
	ole.IUnknownVtbl
	Buffer uintptr
}

func (v *iBufferByteAccess) VTable() *iBufferByteAccessVtbl {
	return (*iBufferByteAccessVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iBufferByteAccess) Buffer(bufPtr **byte) error {
	return bbaBuffer(v, bufPtr)
}
