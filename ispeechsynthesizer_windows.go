// +buil windows
package wss

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func synthSynthesizeTextToStreamAsync(synth *ISpeechSynthesizer, text string, stream *ISpeechSynthesisStream) (err error) {
	var hstr ole.HString

	if hstr, err = ole.NewHString(text); err != nil {
		return
	}
	hr, _, _ := syscall.Syscall(
		synth.VTable().SynthesizeTextToStreamAsync,
		3,
		uintptr(unsafe.Pointer(synth)),
		uintptr(hstr),
		uintptr(unsafe.Pointer(stream)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func synthSynthesizeSsmlToStreamAsync(synth *ISpeechSynthesizer, text string, stream *ISpeechSynthesisStream) (err error) {
	var hstr ole.HString

	if hstr, err = ole.NewHString(text); err != nil {
		return
	}
	hr, _, _ := syscall.Syscall(
		synth.VTable().SynthesizeSsmlToStreamAsync,
		3,
		uintptr(unsafe.Pointer(synth)),
		uintptr(hstr),
		uintptr(unsafe.Pointer(stream)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func synthPutVoice(synth *ISpeechSynthesizer, info *IVoiceInformation) (err error) {
	hr, _, _ := syscall.Syscall(
		synth.VTable().PutVoice,
		2,
		uintptr(unsafe.Pointer(synth)),
		uintptr(unsafe.Pointer(info)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func synthGetVoice(synth *ISpeechSynthesizer, info **IVoiceInformation) (err error) {
	hr, _, _ := syscall.Syscall(
		synth.VTable().GetVoice,
		2,
		uintptr(unsafe.Pointer(synth)),
		uintptr(unsafe.Pointer(info)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
