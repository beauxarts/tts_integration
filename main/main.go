package main

import (
	"fmt"
	tts_integration "github.com/beauxarts/tts-integration"
	"io"
	"net/http"
	"os"
)

func main() {

	fk := "./secret_api_key.txt"
	key := ""
	if _, err := os.Stat(fk); err == nil {
		f, err := os.Open(fk)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		bts, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}
		key = string(bts)
	}

	if vcs, err := tts_integration.GetVoices(http.DefaultClient, "ru", key); err == nil {
		fmt.Println(vcs)
	}
}
