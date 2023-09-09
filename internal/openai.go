package internal

import (
	"context"
	"os"

	"github.com/sashabaranov/go-openai"
)

func GetEmbedding(query string) ([]float32, error) {
	client := openai.NewClient(os.Getenv("OPENAI"))
	res, err := client.CreateEmbeddings(context.TODO(), openai.EmbeddingRequestStrings{
		Input: []string{query},
		Model: openai.AdaEmbeddingV2,
		User:  "FM-BOT",
	})
	if err != nil {
		return nil, err
	}
	return res.Data[0].Embedding, nil
}
