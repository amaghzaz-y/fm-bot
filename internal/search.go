package internal

import (
	"context"

	"github.com/amaghzaz-y/fm-bot/pkg/db/index"
	"github.com/sashabaranov/go-openai"
)

func (b *Bot) Search(query string) (*[]index.SearchResult[string], error) {
	vec, err := b.vectorize(query)
	if err != nil {
		return nil, nil
	}
	res, err := b.Index.SearchByVector(vec, 10, 10)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (b *Bot) vectorize(query string) ([]float32, error) {
	res, err := b.OpenAI.CreateEmbeddings(context.TODO(), openai.EmbeddingRequestStrings{
		Input: []string{query},
		Model: openai.AdaEmbeddingV2,
		User:  "FM-BOT",
	})
	if err != nil {
		return nil, err
	}
	return res.Data[0].Embedding, nil
}
