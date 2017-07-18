// +build !windows

package wss

import (
	"github.com/go-ole/go-ole"
)

func vvGetAt(vv *iVectorView, index uint16, voiceInformation **iVoiceInformation) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func vvGetSize(vv *iVectorView, size *uint16) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func vvIndexOf(vv *iVectorView, voiceInformation *iVoiceInformation, index *uint16, found *bool) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func vvGetMany(vv *iVectorView) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
