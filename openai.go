package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"reflect"
)

const APIBase = "https://api.openai.com/v1/"

type OpenAI struct {
	apiKey string
	Client *http.Client
	ctx    context.Context
}

type Error struct {
	Message string      `json:"message"`
	Type    string      `json:"type"`
	Param   interface{} `json:"param"`
	Code    interface{} `json:"code"`
}

type ErrorResponse struct {
	Error Error `json:"error"`
}

func NewClient(ctx context.Context, apiKey string) (*OpenAI, error) {
	if apiKey == "" {
		return nil, errors.New("ApiKey are required to create a Client")
	}

	return &OpenAI{
		ctx:    ctx,
		apiKey: apiKey,
		Client: &http.Client{},
	}, nil
}

func (a *OpenAI) SetClient(client *http.Client) {
	a.Client = client
}

func (a *OpenAI) getUrl(uri string) string {
	return fmt.Sprintf("%s%s", APIBase, uri)
}

func (a *OpenAI) NewRequest(method, url string, payload interface{}) (*http.Request, error) {
	var buf io.Reader
	if payload != nil {
		b, err := json.Marshal(&payload)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}
	return http.NewRequestWithContext(a.ctx, method, url, buf)
}

func (a *OpenAI) NewFormRequest(method, url string, v reflect.Value) (*http.Request, error) {
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		val := v.Field(i)
		if val.IsValid() && !val.IsZero() {
			tag := field.Tag.Get("form")
			if field.Type.Kind() == reflect.Ptr {
				part, _ := writer.CreateFormFile(tag, path.Base(val.Interface().(*os.File).Name()))
				io.Copy(part, val.Interface().(*os.File))
			}
			if field.Type.Kind() == reflect.Int {
				writer.WriteField(tag, fmt.Sprintf("%d", val.Int()))
			}
			if field.Type.Kind() == reflect.String {
				writer.WriteField(tag, fmt.Sprintf("%s", val.String()))
			}
		}
	}

	writer.Close()

	req, err := http.NewRequestWithContext(a.ctx, method, url, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())

	return req, nil
}

func (a *OpenAI) Send(req *http.Request, v interface{}) error {
	var (
		err  error
		resp *http.Response
	)

	if req.Header.Get("Content-type") == "" {
		req.Header.Set("Content-type", "application/json")
	}
	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/json")
	}
	req.Header.Add("Authorization", "Bearer "+a.apiKey)

	resp, err = a.Client.Do(req)

	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) error {
		return Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		var errResp ErrorResponse

		if err := json.NewDecoder(resp.Body).Decode(&errResp); err == nil {
			return errors.New(errResp.Error.Message)
		}

		return errors.New(resp.Status)
	}

	if reflect.ValueOf(v).Elem().Kind() == reflect.String {
		if buf, err := io.ReadAll(resp.Body); err == nil {
			rv := reflect.ValueOf(v)
			rv.Elem().SetString(string(buf))
			return nil
		}
	}

	return json.NewDecoder(resp.Body).Decode(v)
}
