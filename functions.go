// Package wss is Go bindings for Windows SpeechSynthesizer (WinRT) API.
package wss

// TextToStream generates WAV file which contains synthesized voice from plain text.
// The WAV file is encoded as 22050 kHz / 16 bit monoral audio.
func TextToStream(index int, text string) ([]byte, error) {
	return textToStream(index, text)
}

// TextToStream generates WAV file which contains synthesized voice from SSML.
// The WAV file is encoded as 22050 kHz / 16 bit monoral audio.
func SsmlToStream(index int, ssml string) ([]byte, error) {
	return ssmlToStream(index, ssml)
}

// GetVoices returns slice of VoiceInformation.
func GetVoices() ([]VoiceInformation, error) {
	return getVoices()
}
