// +build windows
package wss

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func mpGetAutoPlay(mp *IMediaPlayer, state *bool) (err error) {
	hr, _, _ := syscall.Syscall(
		mp.VTable().GetAutoPlay,
		2,
		uintptr(unsafe.Pointer(mp)),
		uintptr(unsafe.Pointer(state)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func mpPutAutoPlay(mp *IMediaPlayer, state bool) (err error) {
	var boolValue uint32
	if state {
		boolValue = 1
	}
	hr, _, _ := syscall.Syscall(
		mp.VTable().PutAutoPlay,
		2,
		uintptr(unsafe.Pointer(mp)),
		uintptr(boolValue),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
func mpGetNaturalDuration(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpGetPosition(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpPutPosition(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpGetBufferingProgress(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpGetCurrentState(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpGetCanSeek(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpGetCanPause(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpGetIsLoopingEnabled(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpPutIsLoopingEnabled(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpGetIsProtected(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpGetIsMuted(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpPutIsMuted(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpGetPlaybackRate(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpPutPlaybackRate(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpGetVolume(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpPutVolume(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpGetPlaybackMediaMarkers(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpAddMediaOpened(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpRemoveMediaOpened(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpAddMediaEnded(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpRemoveMediaEnded(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpAddMediaFailed(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpRemoveMediaFailed(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpAddCurrentStateChanged(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpRemoveCurrentStateChanged(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpAddPlaybackMediaMarkerReached(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpRemovePlaybackMediaMarkerReached(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpAddMediaPlayerRateChanged(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpRemoveMediaPlayerRateChanged(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpAddVolumeChanged(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpRemoveVolumeChanged(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpAddSeekCompleted(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpRemoveSeekCompleted(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpAddBufferingStarted(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpRemoveBufferingStarted(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpAddBufferingEnded(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
func mpRemoveBufferingEnded(mp *IMediaPlayer) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func mpPlay(mp *IMediaPlayer) (err error) {
	hr, _, _ := syscall.Syscall(
		mp.VTable().Play,
		1,
		uintptr(unsafe.Pointer(mp)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return

}

func mpPause(mp *IMediaPlayer) (err error) {
	hr, _, _ := syscall.Syscall(
		mp.VTable().Pause,
		1,
		uintptr(unsafe.Pointer(mp)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
