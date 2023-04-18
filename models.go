package openai

import (
	"net/http"
)

type Permission struct {
	Id                 string      `json:"id"`
	Object             string      `json:"object"`
	Created            int         `json:"created"`
	AllowCreateEngine  bool        `json:"allow_create_engine"`
	AllowSampling      bool        `json:"allow_sampling"`
	AllowLogprobs      bool        `json:"allow_logprobs"`
	AllowSearchIndices bool        `json:"allow_search_indices"`
	AllowView          bool        `json:"allow_view"`
	AllowFineTuning    bool        `json:"allow_fine_tuning"`
	Organization       string      `json:"organization"`
	Group              interface{} `json:"group"`
	IsBlocking         bool        `json:"is_blocking"`
}

type Model struct {
	Id         string       `json:"id"`
	Object     string       `json:"object"`
	Created    int          `json:"created"`
	OwnedBy    string       `json:"owned_by"`
	Permission []Permission `json:"permission"`
	Root       string       `json:"root"`
	Parent     string       `json:"parent"`
}

type Models struct {
	Object string  `json:"object"`
	Data   []Model `json:"data"`
}

func (a *OpenAI) Models() (*Models, error) {
	models := &Models{}
	req, err := a.NewRequest(
		http.MethodGet,
		a.getUrl("models"),
		nil,
	)
	if err != nil {
		return nil, err
	}

	if err = a.Send(req, models); err != nil {
		return nil, err
	}

	return models, nil
}

func (a *OpenAI) RetrieveModel(name string) (*Model, error) {
	model := &Model{}
	req, err := a.NewRequest(
		http.MethodGet,
		a.getUrl("models/"+name),
		nil,
	)
	if err != nil {
		return nil, err
	}

	if err = a.Send(req, model); err != nil {
		return nil, err
	}

	return model, nil
}
