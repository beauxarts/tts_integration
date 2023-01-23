# tts-integration
Tiny Cloud and local services text-to-speech integration module. 

For cloud versions - most likely you would be better served by official modules:
  - `gcp`: `cloud.google.com/go/texttospeech/apiv1`
  - `acs`: `github.com/Microsoft/cognitive-services-speech-sdk-go`

## How to use

Add module to your project: `go get github.com/beauxarts/tts-integration`.

## Usage

This package defines a common `Synthesizer` interface ([source](https://github.com/beauxarts/tts_integration/blob/main/synthesizer.go)), that is realized with `gcp` (Google Cloud), `acs` (Azure Cognitive Services), `say` (macOS say command).

[lego](https://github.com/beauxarts/lego) is a reference implementation of the `Synthesizer` and provider specifics ([source](https://github.com/beauxarts/lego/blob/main/chapter_paragraph/synthesizer.go)). 