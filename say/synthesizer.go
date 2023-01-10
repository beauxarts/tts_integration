package say

import (
	"errors"
	"fmt"
	"github.com/beauxarts/tts_integration"
	"io"
	"os/exec"
	"strings"
	"time"
)

type Synthesizer struct {
	sayCmd      string
	voice       string
	audioFormat AudioFormat
}

func NewSynthesizer(voice, audioFormat string) tts_integration.Synthesizer {
	af := Parse(audioFormat)
	if af == "" {
		af = DefaultAudioFormat
	}
	sayCmd := ""
	if path, err := exec.LookPath("say"); err == nil {
		sayCmd = path
	}
	return &Synthesizer{
		sayCmd:      sayCmd,
		voice:       voice,
		audioFormat: af,
	}
}

func (s *Synthesizer) WriteText(text string, _ io.Writer, name string) error {

	args := []string{"-v", s.voice, "-o", name, "--file-format", s.audioFormat.String(), "--data-format", "aac", text}
	cmd := exec.Command(s.sayCmd, args...)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (s *Synthesizer) WriteSSML(ssml string, _ io.Writer, name string) error {
	return errors.New("SSML synthesis is not supported")
}

func (s *Synthesizer) VoicesStrings(_ ...string) ([]string, error) {
	sb := &strings.Builder{}

	args := []string{"-v", "?"}
	cmd := exec.Command(s.sayCmd, args...)
	cmd.Stdout = sb

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return strings.Split(sb.String(), "\n"), nil
}

func (s *Synthesizer) IsSSMLSupported() bool {
	return false
}

func (s *Synthesizer) IsWriterRequired() bool {
	return false
}

func (s *Synthesizer) IsNameRequired() bool {
	return true
}

func (s *Synthesizer) Pause(dur time.Duration) (string, tts_integration.SynthesisInputType) {
	return fmt.Sprintf("[[slnc %d]]", dur.Milliseconds()),
		tts_integration.Text
}
