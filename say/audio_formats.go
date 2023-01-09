package say

import "strings"

type AudioFormat string

const (
	//3gp2  3GPP-2 Audio         (.3g2,.3gp2) [Qclp,aac,aace,aacf,aacg,aach,aacl,aacp]
	//3gpp  3GP Audio            (.3gp,.3gpp) [Qclp,aac,aace,aacf,aacg,aach,aacl,aacp]
	AIFC = "AIFC" //  AIFC                 (.aifc,.aiff,.aif) [lpcm,ulaw,alaw,ima4,Qclp]
	AIFF = "AIFF" //  AIFF                 (.aiff,.aif) [lpcm]
	BW64 = "BW64" //  WAVE (BW64 for length over 4 GB) (.wav) [lpcm,ulaw,alaw]
	NEXT = "NeXT" //  NeXT/Sun             (.snd,.au) [lpcm,ulaw,alaw]
	RF64 = "RF64" //  WAVE (RF64 for length over 4 GB) (.wav) [lpcm,ulaw,alaw]
	SD2F = "Sd2f" //  Sound Designer II    (.sd2) [lpcm]
	W64F = "W64f" //  Wave64               (.w64) [lpcm,ulaw,alaw]
	WAVE = "WAVE" //  WAVE                 (.wav) [lpcm,ulaw,alaw]
	ADTS = "adts" //  AAC ADTS             (.aac,.adts) [aac,aach,aacp]
	CAFF = "caff" // CAF                  (.caf,.caff) [Qclp,aac,aace,aacf,aacg,aach,aacl,aacp,alac,alaw,flac,ilbc,ima4,lpcm,opus,ulaw]
	FLAC = "flac" //  FLAC                 (.flac) [flac]
	LOAS = "loas" //  LATM/LOAS            (.loas,.latm,.xhe) [aac,aace,aacf,aacg,aach,aacl,aacp]
	M4AF = "m4af" //  Apple MPEG-4 Audio   (.m4a,.m4r) [aac,aace,aacf,aacg,aach,aacl,aacp,alac,flac,lpcm]
	M4BF = "m4bf" //  Apple MPEG-4 AudioBooks (.m4b) [aac,aace,aacf,aacg,aach,aacl,aacp]
	MP4F = "mp4f" //  MPEG-4 Audio         (.mp4,.mpg4) [aac,aace,aacf,aacg,aach,aacl,aacp,alac,flac,lpcm]
)

const (
	DefaultAudioFormat = M4AF
	DefaultAudioExt    = ".m4a"
)

func Parse(afs string) AudioFormat {

	options := []string{
		AIFC,
		AIFF,
		BW64,
		NEXT,
		RF64,
		SD2F,
		W64F,
		WAVE,
		ADTS,
		CAFF,
		FLAC,
		LOAS,
		M4AF,
		M4BF,
		MP4F,
	}

	optionsMap := make(map[string]AudioFormat, len(options))
	for _, o := range options {
		optionsMap[strings.ToLower(o)] = AudioFormat(o)
	}

	if af, ok := optionsMap[strings.ToLower(afs)]; ok {
		return af
	}

	return ""
}

func (af AudioFormat) String() string {
	return string(af)
}
