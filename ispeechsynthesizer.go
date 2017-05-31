package wss

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type ISpeechSynthesizer struct {
	ole.IInspectable
}

type ISpeechSynthesizerVtbl struct {
	ole.IInspectableVtbl
	SynthesizeTextToStreamAsync uintptr
	SynthesizeSsmlToStreamAsync uintptr
	PutVoice                    uintptr
	GetVoice                    uintptr
}

func (v *ISpeechSynthesizer) VTable() *ISpeechSynthesizerVtbl {
	return (*ISpeechSynthesizerVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *ISpeechSynthesizer) SynthesizeTextToStreamAsync(text string, stream *ISpeechSynthesisStream) (err error) {
	err = ssSynthesizeTextToStreamAsync(v, text, stream)
	return
}

func (v *ISpeechSynthesizer) SynthesizeSsmlToStreamAsync(ssml string, stream *ISpeechSynthesisStream) (err error) {
	err = ssSynthesizeSsmlToStreamAsync(v, ssml, stream)
	return
}

func (v *ISpeechSynthesizer) PutVoice(info *IVoiceInformation) (err error) {
	err = ssPutVoice(v, info)
	return
}

func (v *ISpeechSynthesizer) GetVoice(info **IVoiceInformation) (err error) {
	err = ssGetVoice(v, info)
	return
}
