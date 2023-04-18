package openai

import "net/http"

type EmbeddingRequest struct {
	Model string `json:"model"`
	Input string `json:"input"`
	User  string `json:"user"`
}

type Embedding struct {
	Object    string    `json:"object"`
	Embedding []float64 `json:"embedding"`
	Index     int       `json:"index"`
}

type EmbeddingUsage struct {
	PromptTokens int `json:"prompt_tokens"`
	TotalTokens  int `json:"total_tokens"`
}

type Embeddings struct {
	Object string         `json:"object"`
	Data   []Embedding    `json:"data"`
	Model  string         `json:"model"`
	Usage  EmbeddingUsage `json:"usage"`
}

func (a *OpenAI) Embeddings(cr EmbeddingRequest) (*Embeddings, error) {
	embeddings := &Embeddings{}
	req, err := a.NewRequest(
		http.MethodPost,
		a.getUrl("embeddings"),
		cr,
	)
	if err != nil {
		return nil, err
	}

	if err = a.Send(req, embeddings); err != nil {
		return nil, err
	}

	return embeddings, nil
}
