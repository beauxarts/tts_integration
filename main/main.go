package main

import (
	"github.com/beauxarts/tts_integration/say"
)

func main() {
	s := say.NewSynthesizer("Milena (Enhanced)", say.DefaultAudioFormat)

	if err := s.WriteText("Слава коммунистической партии Советского Союза!", nil, "test"+say.DefaultAudioExt); err != nil {
		panic(err)
	}
}
