package wss

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IMediaElement struct {
	ole.IInspectable
}

type IMediaElementVtbl struct {
	ole.IInspectableVtbl
	GetPosterSource                   uintptr
	PutPosterSource                   uintptr
	GetSource                         uintptr
	PutSource                         uintptr
	GetIsMuted                        uintptr
	PutIsMuted                        uintptr
	GetIsAudioOnly                    uintptr
	GetAutoPlay                       uintptr
	PutAutoPlay                       uintptr
	GetVolume                         uintptr
	PutVolume                         uintptr
	GetBalance                        uintptr
	PutBalance                        uintptr
	GetNaturalVideoHeight             uintptr
	GetNaturalVideoWidth              uintptr
	GetNaturalDuration                uintptr
	GetPosition                       uintptr
	PutPosition                       uintptr
	GetDownloadProgress               uintptr
	GetBufferingProgress              uintptr
	GetDownloadProgressOffset         uintptr
	GetCurrentState                   uintptr
	GetMarkers                        uintptr
	GetCanSeek                        uintptr
	GetCanPause                       uintptr
	GetAudioStreamCount               uintptr
	GetAudioStreamIndex               uintptr
	PutAudioStreamIndex               uintptr
	GetPlaybackRate                   uintptr
	PutPlaybackRate                   uintptr
	GetIsLooping                      uintptr
	PutIsLooping                      uintptr
	GetPlayToSource                   uintptr
	GetDefaultPlaybackRate            uintptr
	PutDefaultPlaybackRate            uintptr
	GetAspectRatioWidth               uintptr
	GetAspectRatioHeight              uintptr
	GetRealTimePlayback               uintptr
	PutRealTimePlayback               uintptr
	GetAudioCategory                  uintptr
	PutAudioCategory                  uintptr
	GetAudioDeviceType                uintptr
	PutAudioDeviceType                uintptr
	GetProtectionManager              uintptr
	PutProtectionManager              uintptr
	GetStereo3DVideoPackingMode       uintptr
	PutStereo3DVideoPackingMode       uintptr
	GetStereo3DVideoRenderMode        uintptr
	PutStereo3DVideoRenderMode        uintptr
	GetIsStereo3DVideo                uintptr
	AddMediaOpened                    uintptr
	RemoveMediaOpened                 uintptr
	AddMediaEnded                     uintptr
	RemoveMediaEnded                  uintptr
	AddMediaFailed                    uintptr
	RemoveMediaFailed                 uintptr
	AddDownloadProgressChanged        uintptr
	RemoveDownloadProgressChanged     uintptr
	AddBufferingProgressChanged       uintptr
	RemoveBufferingProgressChanged    uintptr
	AddCurrentStateChanged            uintptr
	RemoveCurrentStateChanged         uintptr
	AddMarkerReached                  uintptr
	RemoveMarkerReached               uintptr
	AddRateChanged                    uintptr
	RemoveRateChanged                 uintptr
	AddVolumeChanged                  uintptr
	RemoveVolumeChanged               uintptr
	AddSeekCompleted                  uintptr
	RemoveSeekCompleted               uintptr
	Stop                              uintptr
	Play                              uintptr
	Pause                             uintptr
	CanPlayType                       uintptr
	SetSource                         uintptr
	GetAudioStreamLanguage            uintptr
	AddAudioEffect                    uintptr
	AddVideoEffect                    uintptr
	RemoveAllEffects                  uintptr
	GetActualStereo3DVideoPackingMode uintptr
}

func (v *IMediaElement) VTable() *IMediaElementVtbl {
	return (*IMediaElementVtbl)(unsafe.Pointer(v.RawVTable))
}
