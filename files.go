package openai

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

type File struct {
	Id        string `json:"id"`
	Object    string `json:"object"`
	Bytes     int    `json:"bytes"`
	CreatedAt int    `json:"created_at"`
	Filename  string `json:"filename"`
	Purpose   string `json:"purpose"`
}

type Files struct {
	Object string `json:"object"`
	Data   []File `json:"data"`
}

type FileDeleted struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

type UploadFileRequest struct {
	File    *os.File `form:"file"`
	Purpose string   `form:"purpose"`
}

func (a *OpenAI) Files() (*Files, error) {
	files := &Files{}
	req, err := a.NewRequest(
		http.MethodGet,
		a.getUrl("files"),
		nil,
	)
	if err != nil {
		return nil, err
	}

	if err = a.Send(req, files); err != nil {
		return nil, err
	}

	return files, nil
}

func (a *OpenAI) UploadFile(cr UploadFileRequest) (*File, error) {
	file := &File{}
	req, err := a.NewFormRequest(
		http.MethodPost,
		a.getUrl("files"),
		reflect.ValueOf(cr),
	)
	if err != nil {
		return nil, err
	}

	if err = a.Send(req, file); err != nil {
		return nil, err
	}

	return file, nil
}

func (a *OpenAI) DeleteFile(fileId string) (*FileDeleted, error) {
	res := &FileDeleted{}
	req, err := a.NewRequest(
		http.MethodDelete,
		a.getUrl(fmt.Sprintf("files/%s", fileId)),
		nil,
	)
	if err != nil {
		return nil, err
	}

	if err = a.Send(req, res); err != nil {
		return nil, err
	}

	return res, nil
}
