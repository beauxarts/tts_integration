package acs

type AudioOutput string

const (
	//Streaming

	//amr-wb-16000hz

	Audio16khz32kbpsMonoOpus AudioOutput = "audio-16khz-16bit-32kbps-mono-opus"

	Audio24khz24kbpsMonoOpus AudioOutput = "audio-24khz-16bit-24kbps-mono-opus"
	Audio24khz48kbpsMonoOpus AudioOutput = "audio-24khz-16bit-48kbps-mono-opus"

	Audio16khz32kbpsMonoMp3  AudioOutput = "audio-16khz-32kbitrate-mono-mp3"
	Audio16khz64kbpsMonoMp3  AudioOutput = "audio-16khz-64kbitrate-mono-mp3"
	Audio16khz128kbpsMonoMp3 AudioOutput = "audio-16khz-128kbitrate-mono-mp3"

	Audio24khz48kbpsMonoMp3  AudioOutput = "audio-24khz-48kbitrate-mono-mp3"
	Audio24khz96kbpsMonoMp3  AudioOutput = "audio-24khz-96kbitrate-mono-mp3"
	Audio24khz160kbpsMonoMp3 AudioOutput = "audio-24khz-160kbitrate-mono-mp3"

	Audio48khz96kbpsMonoMp3  AudioOutput = "audio-48khz-96kbitrate-mono-mp3"
	Audio48khz192kbpsMonoMp3 AudioOutput = "audio-48khz-192kbitrate-mono-mp3"

	Ogg16khzOpus AudioOutput = "ogg-16khz-16bit-mono-opus"
	Ogg24khzOpus AudioOutput = "ogg-24khz-16bit-mono-opus"
	Ogg48khzOpus AudioOutput = "ogg-48khz-16bit-mono-opus"

	Raw8khz8bitAlaw  AudioOutput = "raw-8khz-8bit-mono-alaw"
	Raw8khz8bitMulaw AudioOutput = "raw-8khz-8bit-mono-mulaw"
	Raw8khz16bitPCM  AudioOutput = "raw-8khz-16bit-mono-pcm"
	Raw16khzPCM      AudioOutput = "raw-16khz-16bit-mono-pcm"
	Raw16khzTrueSilk AudioOutput = "raw-16khz-16bit-mono-truesilk"
	Raw22khzPCM      AudioOutput = "raw-22050hz-16bit-mono-pcm"
	Raw24khzPCM      AudioOutput = "raw-24khz-16bit-mono-pcm"
	Raw24khzTrueSilk AudioOutput = "raw-24khz-16bit-mono-truesilk"
	Raw44khzPCM      AudioOutput = "raw-44100hz-16bit-mono-pcm"
	Raw48khzPCM      AudioOutput = "raw-48khz-16bit-mono-pcm"

	WebM16khzOpus       AudioOutput = "webm-16khz-16bit-mono-opus"
	WebM24khz24kbpsOpus AudioOutput = "webm-24khz-16bit-24kbps-mono-opus"
	WebM24khzOpus       AudioOutput = "webm-24khz-16bit-mono-opus"

	//NonStreaming

	Riff8khzAlaw  AudioOutput = "riff-8khz-8bit-mono-alaw"
	Riff8khzMulaw AudioOutput = "riff-8khz-8bit-mono-mulaw"
	Riff8khzPCM   AudioOutput = "riff-8khz-16bit-mono-pcm"
	Riff22khzPCM  AudioOutput = "riff-22050hz-16bit-mono-pcm"
	Riff24khzPCM  AudioOutput = "riff-24khz-16bit-mono-pcm"
	Riff44khzPCM  AudioOutput = "riff-44100hz-16bit-mono-pcm"
	Riff48khzPCM  AudioOutput = "riff-48khz-16bit-mono-pcm"
)

const (
	DefaultAudioOutput    = Audio48khz192kbpsMonoMp3
	DefaultAudioOutputExt = ".mp3"
)

func (ao AudioOutput) String() string {
	return string(ao)
}
