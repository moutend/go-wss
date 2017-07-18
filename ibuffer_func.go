// +build !windows

package wss

import (
	"github.com/go-ole/go-ole"
)

func bGetCapacity(b *iBuffer, capacity *uint32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func bGetLength(b *iBuffer, length *uint32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func bPutLength(b *iBuffer, length uint32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
