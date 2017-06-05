package wss

import (
	"github.com/go-ole/go-ole"
)

var (
	IID_IInstalledVoicesStatic    = ole.NewGUID("{7D526ECC-7533-4C3F-85BE-888C2BAEEBDC}")
	IID_ISpeechSynthesisStream    = ole.NewGUID("{83E46E93-244C-4622-BA0B-6229C4D0D65D}")
	IID_ISpeechSynthesizer        = ole.NewGUID("{CE9F7C76-97F4-4CED-AD68-D51C458E45C6}")
	IID_ISpeechSynthesizer2       = ole.NewGUID("{A7C5ECB2-4339-4D6A-BBF8-C7A4F1544C2E}")
	IID_ISpeechSynthesizerOptions = ole.NewGUID("{A0E23871-CC3D-43C9-91B1-EE185324D83D}")
	IID_IVoiceInformation         = ole.NewGUID("{B127D6A4-1291-4604-AA9C-83134083352C}")
	IID_IMediaElement             = ole.NewGUID("{A38ED2CF-13DE-4299-ADE2-AE18F74ED353}")
)
