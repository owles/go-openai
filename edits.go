package openai

import "net/http"

type EditRequest struct {
	Model       string `json:"model"`
	Input       string `json:"input"`
	Instruction string `json:"instruction"`
	Temperature int    `json:"temperature,omitempty"`
	TopP        int    `json:"top_p,omitempty"`
	N           int    `json:"n,omitempty"`
}

type EditChoice struct {
	Text  string `json:"text"`
	Index int    `json:"index"`
}

type Edit struct {
	Object  string       `json:"object"`
	Created int          `json:"created"`
	Choices []EditChoice `json:"choices"`
	Usage   Usage        `json:"usage"`
}

func (a *OpenAI) Edit(cr EditRequest) (*Edit, error) {
	edit := &Edit{}
	req, err := a.NewRequest(
		http.MethodPost,
		a.getUrl("edits"),
		cr,
	)
	if err != nil {
		return nil, err
	}

	if err = a.Send(req, edit); err != nil {
		return nil, err
	}

	return edit, nil
}
