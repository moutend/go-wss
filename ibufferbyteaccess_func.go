// +build !windows

package wss

import (
	"github.com/go-ole/go-ole"
)

func bbaBuffer(bba *iBufferByteAccess, bufPtr **byte) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
