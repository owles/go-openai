package openai

import (
	"net/http"
	"os"
	"reflect"
)

type AudioResponseFormat string

const (
	AudioFormatJSON  AudioResponseFormat = "json"
	AudioFormatTEXT  AudioResponseFormat = "text"
	AudioFormatSRT   AudioResponseFormat = "srt"
	AudioFormatVJSON AudioResponseFormat = "verbose_json"
	AudioFormatVTT   AudioResponseFormat = "vtt"
)

type TranscriptionRequest struct {
	File           *os.File            `form:"file"`
	Model          string              `form:"model"`
	Prompt         string              `form:"prompt"`
	ResponseFormat AudioResponseFormat `form:"response_format"`
	Temperature    string              `form:"temperature"`
	Language       string              `form:"language"`
}

type TranslationRequest struct {
	File           *os.File            `form:"file"`
	Model          string              `form:"model"`
	Prompt         string              `form:"prompt"`
	ResponseFormat AudioResponseFormat `form:"response_format"`
	Temperature    string              `form:"temperature"`
}

type Transcription struct {
	Text string `json:"text"`
}

type Translation struct {
	Text string `json:"text"`
}

func (a *OpenAI) AudioTranscription(cr TranscriptionRequest) (*Transcription, error) {
	transcription := &Transcription{}
	req, err := a.NewFormRequest(
		http.MethodPost,
		a.getUrl("audio/transcriptions"),
		reflect.ValueOf(cr),
	)
	if err != nil {
		return nil, err
	}

	if err = a.Send(req, transcription); err != nil {
		return nil, err
	}

	return transcription, nil
}

func (a *OpenAI) AudioTranslation(cr TranslationRequest) (*Translation, error) {
	translation := &Translation{}
	req, err := a.NewFormRequest(
		http.MethodPost,
		a.getUrl("audio/translations"),
		reflect.ValueOf(cr),
	)
	if err != nil {
		return nil, err
	}

	if err = a.Send(req, translation); err != nil {
		return nil, err
	}

	return translation, nil
}
