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
	get_AllVoices    uintptr
	get_DefaultVoice uintptr
}

func (v *IInstalledVoicesStatic) VTable() *IInstalledVoicesStaticVtbl {
	return (*IInstalledVoicesStaticVtbl)(unsafe.Pointer(v.RawVTable))
}
