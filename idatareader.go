package wss

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IDataReader struct {
	ole.IInspectable
}

type IDataReaderVtbl struct {
	ole.IInspectableVtbl
	GetUnconsumedBufferLength uintptr
	GetUnicodeEncoding        uintptr
	PutUnicodeEncoding        uintptr
	GetByteOrder              uintptr
	PutByteOrder              uintptr
	GetInputStreamOptions     uintptr
	PutInputStreamOptions     uintptr
	ReadByte                  uintptr
	ReadBytes                 uintptr
	ReadBuffer                uintptr
	ReadBoolean               uintptr
	ReadGuid                  uintptr
	ReadInt16                 uintptr
	ReadInt32                 uintptr
	ReadInt64                 uintptr
	ReadUInt16                uintptr
	ReadUInt32                uintptr
	ReadUInt64                uintptr
	ReadSingle                uintptr
	ReadDouble                uintptr
	ReadString                uintptr
	ReadDateTime              uintptr
	ReadTimeSpan              uintptr
	LoadAsync                 uintptr
	DetachBuffer              uintptr
	DetachStream              uintptr
}

func (v *IDataReader) VTable() *IDataReaderVtbl {
	return (*IDataReaderVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IDataReader) ReadBytes(count uint32, data *byte) (err error) {
	err = drReadBytes(v, count, data)
	return
}
