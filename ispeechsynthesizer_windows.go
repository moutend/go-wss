// +buil windows
package wss

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func ssSynthesizeTextToStreamAsync(ss *ISpeechSynthesizer, text string, stream *ISpeechSynthesisStream) (err error) {
	var hstr ole.HString

	if hstr, err = ole.NewHString(text); err != nil {
		return
	}
	hr, _, _ := syscall.Syscall(
		ss.VTable().SynthesizeTextToStreamAsync,
		3,
		uintptr(unsafe.Pointer(ss)),
		uintptr(hstr),
		uintptr(unsafe.Pointer(stream)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func ssSynthesizeSsmlToStreamAsync(ss *ISpeechSynthesizer, text string, stream *ISpeechSynthesisStream) (err error) {
	var hstr ole.HString

	if hstr, err = ole.NewHString(text); err != nil {
		return
	}
	hr, _, _ := syscall.Syscall(
		ss.VTable().SynthesizeSsmlToStreamAsync,
		3,
		uintptr(unsafe.Pointer(ss)),
		uintptr(hstr),
		uintptr(unsafe.Pointer(stream)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func ssPutVoice(ss *ISpeechSynthesizer, info *IVoiceInformation) (err error) {
	hr, _, _ := syscall.Syscall(
		ss.VTable().PutVoice,
		2,
		uintptr(unsafe.Pointer(ss)),
		uintptr(unsafe.Pointer(info)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func ssGetVoice(ss *ISpeechSynthesizer, info **IVoiceInformation) (err error) {
	hr, _, _ := syscall.Syscall(
		ss.VTable().GetVoice,
		2,
		uintptr(unsafe.Pointer(ss)),
		uintptr(unsafe.Pointer(info)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
