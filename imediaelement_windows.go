// +build windows
package wss

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func meGetPosterSource(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mePutPosterSource(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetSource(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mePutSource(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetIsMuted(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mePutIsMuted(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetIsAudioOnly(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func meGetAutoPlay(me *IMediaElement, state *bool) (err error) {
	hr, _, _ := syscall.Syscall(
		me.VTable().GetAutoPlay,
		2,
		uintptr(unsafe.Pointer(me)),
		uintptr(unsafe.Pointer(state)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func mePutAutoPlay(me *IMediaElement, state bool) (err error) {
	var boolValue uint32

	if state {
		boolValue = 1
	}
	hr, _, _ := syscall.Syscall(
		me.VTable().PutAutoPlay,
		2,
		uintptr(unsafe.Pointer(me)),
		uintptr(boolValue),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func meGetVolume(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mePutVolume(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetBalance(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mePutBalance(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetNaturalVideoHeight(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetNaturalVideoWidth(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetNaturalDuration(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetPosition(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mePutPosition(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetDownloadProgress(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetBufferingProgress(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetDownloadProgressOffset(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetCurrentState(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetMarkers(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetCanSeek(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetCanPause(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetAudioStreamCount(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetAudioStreamIndex(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mePutAudioStreamIndex(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetPlaybackRate(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mePutPlaybackRate(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetIsLooping(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mePutIsLooping(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetPlayToSource(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetDefaultPlaybackRate(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mePutDefaultPlaybackRate(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetAspectRatioWidth(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetAspectRatioHeight(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetRealTimePlayback(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mePutRealTimePlayback(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetAudioCategory(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mePutAudioCategory(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetAudioDeviceType(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mePutAudioDeviceType(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetProtectionManager(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mePutProtectionManager(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetStereo3DVideoPackingMode(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mePutStereo3DVideoPackingMode(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetStereo3DVideoRenderMode(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mePutStereo3DVideoRenderMode(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetIsStereo3DVideo(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meAddMediaOpened(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meRemoveMediaOpened(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meAddMediaEnded(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meRemoveMediaEnded(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meAddMediaFailed(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meRemoveMediaFailed(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meAddDownloadProgressChanged(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meRemoveDownloadProgressChanged(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meAddBufferingProgressChanged(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meRemoveBufferingProgressChanged(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meAddCurrentStateChanged(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meRemoveCurrentStateChanged(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meAddMarkerReached(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meRemoveMarkerReached(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meAddRateChanged(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meRemoveRateChanged(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meAddVolumeChanged(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meRemoveVolumeChanged(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meAddSeekCompleted(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meRemoveSeekCompleted(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meStop(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func mePlay(me *IMediaElement) (err error) {
	hr, _, _ := syscall.Syscall(
		me.VTable().Play,
		1,
		uintptr(unsafe.Pointer(me)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func mePause(me *IMediaElement) (err error) {
	hr, _, _ := syscall.Syscall(
		me.VTable().Pause,
		1,
		uintptr(unsafe.Pointer(me)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func meCanPlayType(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meSetSource(me *IMediaElement, ras *IRandomAccessStream, mimeType string) (err error) {
	var mimeTypeHString ole.HString

	if mimeTypeHString, err = ole.NewHString(mimeType); err != nil {
		return
	}
	defer ole.DeleteHString(mimeTypeHString)

	hr, _, _ := syscall.Syscall(
		me.VTable().SetSource,
		3,
		uintptr(unsafe.Pointer(me)),
		uintptr(unsafe.Pointer(ras)),
		uintptr(unsafe.Pointer(mimeTypeHString)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
func meGetAudioStreamLanguage(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meAddAudioEffect(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meAddVideoEffect(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meRemoveAllEffects(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func meGetActualStereo3DVideoPackingMode(me *IMediaElement) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
