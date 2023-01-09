package gcp

type AudioEncoding string

const (
	//AudioEncodingUnspecified - Not specified. Will return result google.rpc.Code.INVALID_ARGUMENT.
	AudioEncodingUnspecified AudioEncoding = "AUDIO_ENCODING_UNSPECIFIED"
	//Linear16 - Uncompressed 16-bit signed little-endian samples (Linear PCM). Audio content returned as LINEAR16 also contains a WAV header.
	Linear16 = "LINEAR16"
	//MP3 audio at 32kbps.
	MP3 = "MP3"
	//OggOpus - Opus encoded audio wrapped in an ogg container. The result will be a file which can be played natively on Android, and in browsers (at least Chrome and Firefox). The quality of the encoding is considerably higher than MP3 while using approximately the same bitrate.
	OggOpus = "OGG_OPUS"
	//Mulaw -  8-bit samples that compand 14-bit audio samples using G.711 PCMU/mu-law. Audio content returned as MULAW also contains a WAV header.
	Mulaw = "MULAW"
	//Alaw - 8-bit samples that compand 14-bit audio samples using G.711 PCMU/A-law. Audio content returned as ALAW also contains a WAV header.
	Alaw = "ALAW"
)
