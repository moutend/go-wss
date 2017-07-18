// +build !windows

package wss

import (
	"github.com/go-ole/go-ole"
)

func viGetDisplayName(vi *iVoiceInformation, name *ole.HString) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func viGetId(vi *iVoiceInformation, id *ole.HString) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func viGetLanguage(vi *iVoiceInformation, language *ole.HString) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func viGetDescription(vi *iVoiceInformation, description *ole.HString) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func viGetGender(vi *iVoiceInformation, gender *VoiceGender) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
