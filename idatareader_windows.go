// +build windows
package wss

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func drGetUnconsumedBufferLength(dr *IDataReader) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func drGetUnicodeEncoding(dr *IDataReader) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func drPutUnicodeEncoding(dr *IDataReader) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func drGetByteOrder(dr *IDataReader) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func drPutByteOrder(dr *IDataReader) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func drGetInputStreamOptions(dr *IDataReader) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func drPutInputStreamOptions(dr *IDataReader) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func drReadByte(dr *IDataReader) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func drReadBytes(dr *IDataReader, count uint32, data *byte) (err error) {
	hr, _, _ := syscall.Syscall(
		dr.VTable().ReadBytes,
		3,
		uintptr(unsafe.Pointer(dr)),
		uintptr(count),
		uintptr(unsafe.Pointer(data)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func drReadBuffer(dr *IDataReader) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func drReadBoolean(dr *IDataReader) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func drReadGuid(dr *IDataReader) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func drReadInt16(dr *IDataReader) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func drReadInt32(dr *IDataReader) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func drReadInt64(dr *IDataReader) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func drReadUInt16(dr *IDataReader) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func drReadUInt32(dr *IDataReader) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func drReadUInt64(dr *IDataReader) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func drReadSingle(dr *IDataReader) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func drReadDouble(dr *IDataReader) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func drReadString(dr *IDataReader) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func drReadDateTime(dr *IDataReader) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func drReadTimeSpan(dr *IDataReader) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func drLoadAsync(dr *IDataReader) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func drDetachBuffer(dr *IDataReader) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func drDetachStream(dr *IDataReader) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
