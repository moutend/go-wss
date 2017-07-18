package wss

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type iVoiceInformation struct {
	ole.IInspectable
}

type iVoiceInformationVtbl struct {
	ole.IInspectableVtbl
	GetDisplayName uintptr
	GetId          uintptr
	GetLanguage    uintptr
	GetDescription uintptr
	GetGender      uintptr
}

func (v *iVoiceInformation) VTable() *iVoiceInformationVtbl {
	return (*iVoiceInformationVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iVoiceInformation) GetDisplayName(name *ole.HString) error {
	return viGetDisplayName(v, name)
}

func (v *iVoiceInformation) GetId(id *ole.HString) error {
	return viGetId(v, id)
}

func (v *iVoiceInformation) GetLanguage(language *ole.HString) error {
	return viGetLanguage(v, language)
}

func (v *iVoiceInformation) GetDescription(description *ole.HString) error {
	return viGetDescription(v, description)
}

func (v *iVoiceInformation) GetGender(gender *VoiceGender) error {
	return viGetGender(v, gender)
}
