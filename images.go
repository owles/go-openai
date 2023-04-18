package openai

import (
	"net/http"
	"os"
	"reflect"
)

type ImageResponseFormat string
type ImageSize string

const (
	ImageFormatURL ImageResponseFormat = "url"
	ImageFormatB64 ImageResponseFormat = "b64_json"

	ImageSize256  ImageSize = "256x256"
	ImageSize512  ImageSize = "512x512"
	ImageSize1024 ImageSize = "1024x1024"
)

type GenerateImageRequest struct {
	Prompt         string              `json:"prompt"`
	N              int                 `json:"n,omitempty"`
	Size           ImageSize           `json:"size,omitempty"`
	ResponseFormat ImageResponseFormat `json:"response_format,omitempty"`
	User           string              `json:"user,omitempty"`
}

type EditImageRequest struct {
	Image          *os.File            `form:"image"`
	Mask           *os.File            `form:"mask"`
	Prompt         string              `form:"prompt"`
	N              int                 `form:"n"`
	Size           ImageSize           `form:"size"`
	ResponseFormat ImageResponseFormat `form:"response_format"`
	User           string              `form:"user"`
}

type VariationImageRequest struct {
	Image          *os.File            `form:"image"`
	N              int                 `form:"n"`
	Size           ImageSize           `form:"size"`
	ResponseFormat ImageResponseFormat `form:"response_format"`
	User           string              `form:"user"`
}

type ImageData struct {
	Url     string `json:"url"`
	B64Json string `json:"b64_json"`
}

type Image struct {
	Created int         `json:"created"`
	Data    []ImageData `json:"data"`
}

func (a *OpenAI) GenerateImage(cr GenerateImageRequest) (*Image, error) {
	image := &Image{}
	req, err := a.NewRequest(
		http.MethodPost,
		a.getUrl("images/generations"),
		cr,
	)
	if err != nil {
		return nil, err
	}

	if err = a.Send(req, image); err != nil {
		return nil, err
	}

	return image, nil
}

func (a *OpenAI) EditImage(cr EditImageRequest) (*Image, error) {
	image := &Image{}
	req, err := a.NewFormRequest(
		http.MethodPost,
		a.getUrl("images/edits"),
		reflect.ValueOf(cr),
	)
	if err != nil {
		return nil, err
	}

	if err = a.Send(req, image); err != nil {
		return nil, err
	}

	return image, nil
}

func (a *OpenAI) VariationImage(cr VariationImageRequest) (*Image, error) {
	image := &Image{}
	req, err := a.NewFormRequest(
		http.MethodPost,
		a.getUrl("images/variations"),
		reflect.ValueOf(cr),
	)
	if err != nil {
		return nil, err
	}

	if err = a.Send(req, image); err != nil {
		return nil, err
	}

	return image, nil
}
