package openai

import "net/http"

type CompletionRequest struct {
	Model            string         `json:"model"`
	Prompt           string         `json:"prompt"`
	Suffix           string         `json:"suffix,omitempty"`
	MaxTokens        int            `json:"max_tokens,omitempty"`
	Temperature      int            `json:"temperature,omitempty"`
	TopP             int            `json:"top_p,omitempty"`
	N                int            `json:"n,omitempty"`
	Stream           bool           `json:"stream,omitempty"`
	Echo             bool           `json:"echo,omitempty"`
	Stop             string         `json:"stop,omitempty"`
	PresencePenalty  int            `json:"presence_penalty,omitempty"`
	FrequencyPenalty int            `json:"frequency_penalty,omitempty"`
	BestOf           int            `json:"best_of,omitempty"`
	LogitBias        map[string]int `json:"logit_bias,omitempty"`
	Logprobs         int            `json:"logprobs,omitempty"`
	User             string         `json:"user,omitempty"`
}

type Choice struct {
	Text         string      `json:"text"`
	Index        int         `json:"index"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason string      `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Completion struct {
	Id      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

func (a *OpenAI) Completion(cr CompletionRequest) (*Completion, error) {
	completion := &Completion{}
	req, err := a.NewRequest(
		http.MethodPost,
		a.getUrl("completions"),
		cr,
	)
	if err != nil {
		return nil, err
	}

	if err = a.Send(req, completion); err != nil {
		return nil, err
	}

	return completion, nil
}
