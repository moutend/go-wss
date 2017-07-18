// +build windows

package wss

import (
	"fmt"
	"reflect"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

var (
	modSpeechSynthesizer, _ = syscall.LoadDLL("speechsynthesizer.dll")

	procSsmlToStream, _ = modSpeechSynthesizer.FindProc("SsmlToStream")
	procTextToStream, _ = modSpeechSynthesizer.FindProc("TextToStream")
	procGetVoices, _    = modSpeechSynthesizer.FindProc("GetVoices")
)

func createStream(index int, input, mime string) ([]byte, error) {
	ole.RoInitialize(1)
	//defer ole.RoUninitialize()

	var err error
	var hString ole.HString
	if hString, err = ole.NewHString(input); err != nil {
		return nil, err
	}
	defer ole.DeleteHString(hString)

	var hr uintptr
	var b *iBuffer
	var proc *syscall.Proc

	if mime == "text/plain" {
		proc = procTextToStream
	} else if mime == "application/ssml+xml" {
		proc = procSsmlToStream
	} else {
		return nil, fmt.Errorf("%v is not supported", mime)
	}
	hr, _, _ = proc.Call(
		uintptr(uint32(index)),
		uintptr(unsafe.Pointer(hString)),
		uintptr(unsafe.Pointer(&b)))
	if hr != 0 {
		return nil, fmt.Errorf("failed to synthesize voice")
	}
	defer b.Release()

	var length uint32
	if err = b.GetLength(&length); err != nil {
		return nil, err
	}

	var u *ole.IUnknown
	if b.PutQueryInterface(ole.IID_IUnknown, &u); err != nil {
		return nil, err
	}

	var bba *iBufferByteAccess
	var IID_IBufferByteAccess = ole.NewGUID("905a0fef-bc53-11df-8c49-001e4fc686da")
	if err = u.PutQueryInterface(IID_IBufferByteAccess, &bba); err != nil {
		return nil, err
	}

	var bufPtr *byte
	if err = bba.Buffer(&bufPtr); err != nil {
		return nil, err
	}

	// Convert native byte array to byte slice.
	sliceHeader := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(bufPtr)),
		Len:  int(length),
		Cap:  int(length),
	}
	return *(*[]byte)(unsafe.Pointer(&sliceHeader)), nil
}

func textToStream(index int, text string) ([]byte, error) {
	return createStream(index, text, "text/plain")
}

func ssmlToStream(index int, ssml string) ([]byte, error) {
	return createStream(index, ssml, "application/ssml+xml")
}

func getVoices() ([]VoiceInformation, error) {
	ole.RoInitialize(1)
	//defer ole.RoUninitialize()

	var err error
	var vv *iVectorView
	var hr uintptr

	hr, _, _ = procGetVoices.Call(
		uintptr(unsafe.Pointer(&vv)))
	if hr != 0 {
		return nil, fmt.Errorf("failed to get voices")
	}
	defer vv.Release()

	var size uint16
	if err = vv.GetSize(&size); err != nil {
		return nil, err
	}

	var vi *iVoiceInformation
	var hString ole.HString
	var gender VoiceGender
	var v VoiceInformation
	var vs = make([]VoiceInformation, int(size))

	for i := 0; i < int(size); i++ {
		if err = vv.GetAt(uint16(i), &vi); err != nil {
			return nil, err
		}
		v = VoiceInformation{}

		if err = vi.GetDisplayName(&hString); err != nil {
			return nil, err
		}
		v.Name = hString.String()

		if err = vi.GetLanguage(&hString); err != nil {
			return nil, err
		}
		v.Language = hString.String()

		if err = vi.GetDescription(&hString); err != nil {
			return nil, err
		}
		v.Description = hString.String()

		if err = vi.GetGender(&gender); err != nil {
			return nil, err
		}
		v.Gender = gender

		vs[i] = v
		vi.Release()
	}
	return vs, nil
}
