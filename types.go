package wss

// VoiceGender represents gender of voice.
// VoiceGender_Male = 0x0
// VoiceGender_FEMale = 0x1
type VoiceGender uint32

// VoiceInformation represents characteristic of voice.
type VoiceInformation struct {
	Name        string
	Language    string
	Description string
	Gender      VoiceGender
}
