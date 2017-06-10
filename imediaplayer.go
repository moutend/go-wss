package wss

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IMediaPlayer struct {
	ole.IInspectable
}

type IMediaPlayerVtbl struct {
	ole.IInspectableVtbl
	GetAutoPlay                      uintptr
	PutAutoPlay                      uintptr
	GetNaturalDuration               uintptr
	GetPosition                      uintptr
	PutPosition                      uintptr
	GetBufferingProgress             uintptr
	GetCurrentState                  uintptr
	GetCanSeek                       uintptr
	GetCanPause                      uintptr
	GetIsLoopingEnabled              uintptr
	PutIsLoopingEnabled              uintptr
	GetIsProtected                   uintptr
	GetIsMuted                       uintptr
	PutIsMuted                       uintptr
	GetPlaybackRate                  uintptr
	PutPlaybackRate                  uintptr
	GetVolume                        uintptr
	PutVolume                        uintptr
	GetPlaybackMediaMarkers          uintptr
	AddMediaOpened                   uintptr
	RemoveMediaOpened                uintptr
	AddMediaEnded                    uintptr
	RemoveMediaEnded                 uintptr
	AddMediaFailed                   uintptr
	RemoveMediaFailed                uintptr
	AddCurrentStateChanged           uintptr
	RemoveCurrentStateChanged        uintptr
	AddPlaybackMediaMarkerReached    uintptr
	RemovePlaybackMediaMarkerReached uintptr
	AddMediaPlayerRateChanged        uintptr
	RemoveMediaPlayerRateChanged     uintptr
	AddVolumeChanged                 uintptr
	RemoveVolumeChanged              uintptr
	AddSeekCompleted                 uintptr
	RemoveSeekCompleted              uintptr
	AddBufferingStarted              uintptr
	RemoveBufferingStarted           uintptr
	AddBufferingEnded                uintptr
	RemoveBufferingEnded             uintptr
	Play                             uintptr
	Pause                            uintptr
}

func (v *IMediaPlayer) VTable() *IMediaPlayerVtbl {
	return (*IMediaPlayerVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IMediaPlayer) GetAutoPlay(state *bool) (err error) {
	err = mpGetAutoPlay(v, state)
	return
}

func (v *IMediaPlayer) PutAutoPlay(state bool) (err error) {
	err = mpPutAutoPlay(v, state)
	return
}

func (v *IMediaPlayer) Play() (err error) {
	err = mpPlay(v)
	return
}

func (v *IMediaPlayer) Pause() (err error) {
	err = mpPause(v)
	return
}
