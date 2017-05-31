package wss

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IInstalledVoicesStatic struct {
	ole.IInspectable
}

type IInstalledVoicesStaticVtbl struct {
	ole.IInspectableVtbl
	GetAllVoices    uintptr
	GetDefaultVoice uintptr
}

func (v *IInstalledVoicesStatic) VTable() *IInstalledVoicesStaticVtbl {
	return (*IInstalledVoicesStaticVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IInstalledVoicesStatic) GetAllVoices(voices *IVoiceInformation) (err error) {
	err = ivsGetAllVoices(v, voices)
	return
}

func (v *IInstalledVoicesStatic) GetDefaultVoice(voice *IVoiceInformation) (err error) {
	err = ivsGetDefaultVoice(v, voice)
	return
}
