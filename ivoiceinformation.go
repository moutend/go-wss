package wss

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IVoiceInformation struct {
	ole.IInspectable
}

type IVoiceInformationVtbl struct {
	ole.IInspectableVtbl
	GetDisplayName uintptr
	GetId          uintptr
	GetLanguage    uintptr
	GetDescription uintptr
	GetGender      uintptr
}

func (v *IVoiceInformation) VTable() *IVoiceInformationVtbl {
	return (*IVoiceInformationVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IVoiceInformation) GetDisplayName(name *ole.HString) (err error) {
	err = viGetDisplayName(v, name)
	return
}

func (v *IVoiceInformation) GetId(id *ole.HString) (err error) {
	err = viGetId(v, id)
	return
}

func (v *IVoiceInformation) GetLanguage(language *ole.HString) (err error) {
	err = viGetLanguage(v, language)
	return
}

func (v *IVoiceInformation) GetDescription(description *ole.HString) (err error) {
	err = viGetDescription(v, description)
	return
}

func (v *IVoiceInformation) GetGender(gender *VoiceGender) (err error) {
	err = viGetGender(v, gender)
	return
}
