package wss

import (
	"unsafe"

	"github.com/go-olego-ole"
)

type IRandomAccessStream struct {
	ole.IInspectable
}

type IRandomAccessStreamVtbl struct {
	ole.IInspectableVtbl
	GetSize          uintptr
	PutSize          uintptr
	GetInputStreamAt uintptr
	GetOutputStreamAt
	GetPosition uintptr
	Seek        uintptr
	CloneStream uintptr
	GetCanRead  uintptr
	GetCanWrite uintptr
}
