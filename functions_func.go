// +build !windows

package wss

import (
	"github.com/go-ole/go-ole"
)

func textToStream(index int, text string) ([]byte, error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}

func ssmlToStream(index int, ssml string) ([]byte, error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}

func getVoices() ([]VoiceInformation, error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}
