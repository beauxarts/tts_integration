# google-tts-integration
Tiny Google Client Text-to-Speech integration module. 

Most likely you would be better served by an official module `cloud.google.com/go/texttospeech/apiv1`.

## How to use

Add module to your project: `go get github.com/beauxarts/google-tts-integration`.

## Prerequisites

- Create (or use existing) project on Google Cloud Console
- Enable `texttospeech.googleapis.com` API
- Navigate to the the `API` > `Credentials` section:
  - Create an API key
  - Create a service account (source: https://cloud.google.com/text-to-speech/docs/create-audio-text-client-libraries)

## Getting available voices for a locale

NOTE: error handling omitted for brevity.

```go
//key := "Text-to-speech API key"
voices, _ := google_tts_integration.GetVoices(http.DefaultClient, "en-US", key)
```

## Synthesizing audio from text

NOTE: error handling omitted for brevity.

```go
//key := "Text-to-speech API key"
text := "Hello, World!" //The input size is limited to 5000 characters.
voice := google_tts_integration.NewVoiceSelectionParams("en-US", "en-US-WaveNet-A", google_tts_integration.Male)

tsr, _ := google_tts_integration.PostTextSynthesize(http.DefaultClient, text, voice, key)

//tsr.Bytes() is an OggOpus encoded audio (default format used by this module)
```