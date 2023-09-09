package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

type Embedding struct {
	Text   string    `json:"text"`
	Vector []float32 `json:"vector"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	text, err := os.ReadFile("assets/rules.txt")
	if err != nil {
		panic(err)
	}
	parts := splitPartN(string(text))
	fmt.Println("vectorizing...")
	vecs, err := vectorize(parts)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("exporting...")
	err = export(vecs)
	if err != nil {
		log.Fatalln(err)
	}
}

func vectorize(parts []string) ([]Embedding, error) {
	var embeddings []Embedding
	client := openai.NewClient(os.Getenv("OPENAI"))
	res, err := client.CreateEmbeddings(context.TODO(), openai.EmbeddingRequestStrings{
		Input: parts,
		Model: openai.AdaEmbeddingV2,
		User:  "FM-BOT",
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("total tokens:", res.Usage.TotalTokens,
		"-",
		"price:", float32(res.Usage.TotalTokens/1000)*0.0001, "$",
		"-",
		"model", res.Model,
	)
	for _, data := range res.Data {
		embeddings = append(embeddings, Embedding{
			Text:   parts[data.Index],
			Vector: data.Embedding,
		})
	}
	return embeddings, nil
}

func export(embeddings []Embedding) error {
	payload, err := json.Marshal(&embeddings)
	if err != nil {
		return err
	}
	return os.WriteFile("assets/embeddings.json", payload, os.ModePerm)
}

func splitPartN(text string) []string {
	words := strings.Split(text, ".")
	var result []string
	for _, word := range words {
		if len(word) > 4 {
			result = append(result, word)
		}
	}
	return result
}
