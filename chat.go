package openai

import "net/http"

type Role string

const (
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
	RoleSystem    Role = "system"
)

type Message struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionRequest struct {
	Model            string         `json:"model"`
	Messages         []Message      `json:"messages"`
	Temperature      int            `json:"temperature,omitempty"`
	TopP             int            `json:"top_p,omitempty"`
	N                int            `json:"n,omitempty"`
	Stream           bool           `json:"stream,omitempty"`
	Stop             string         `json:"stop,omitempty"`
	MaxTokens        int            `json:"max_tokens,omitempty"`
	PresencePenalty  int            `json:"presence_penalty,omitempty"`
	FrequencyPenalty int            `json:"frequency_penalty,omitempty"`
	LogitBias        map[string]int `json:"logit_bias,omitempty"`
	User             string         `json:"user,omitempty"`
}

type ChatChoice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type ChatCompletion struct {
	Id      string       `json:"id"`
	Object  string       `json:"object"`
	Created int          `json:"created"`
	Choices []ChatChoice `json:"choices"`
	Usage   Usage        `json:"usage"`
}

func (a *OpenAI) ChatCompletion(cr ChatCompletionRequest) (*ChatCompletion, error) {
	completion := &ChatCompletion{}
	req, err := a.NewRequest(
		http.MethodPost,
		a.getUrl("chat/completions"),
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
