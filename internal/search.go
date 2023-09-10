package fmbot

import (
	"context"
	"log"

	"github.com/sashabaranov/go-openai"
)

type Replies struct {
	Prompt  interface{}
	Replies []string
}

func (b *Bot) Reply(query string) (*Replies, error) {
	ids, err := b.getIDs(query)
	if err != nil {
		return nil, err
	}
	
	info, err := b.getInfo(ids)
	if err != nil {
		return nil, err
	}
	prompt := buildPrompt(query, info)
	answers, err := b.getAnswer(prompt)
	if err != nil {
		return nil, err
	}
	return &Replies{
		Prompt:  prompt,
		Replies: answers,
	}, nil
}

func (b *Bot) getIDs(query string) ([]string, error) {
	vec, err := b.vectorize(query)
	if err != nil {
		return nil, nil
	}
	res, err := b.Index.SearchByVector(vec, 10, 10)
	if err != nil {
		return nil, err
	}
	var ids []string
	for _, r := range res {
		ids = append(ids, r.ID)
	}
	return ids, nil
}

func (b *Bot) getInfo(ids []string) ([]string, error) {
	var texts []string
	for _, id := range ids {
		emb, err := b.getEmbeddingByID(id)
		if err != nil {
			log.Println("error getting embbedding by id", err)
			continue
		}
		texts = append(texts, emb.Text)
	}
	return texts, nil
}

func (b *Bot) getAnswer(prompt []openai.ChatCompletionMessage) ([]string, error) {
	res, err := b.OpenAI.CreateChatCompletion(context.TODO(), openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo,
		Messages: prompt,
	})
	if err != nil {
		return nil, err
	}
	var results []string
	for _, comp := range res.Choices {
		results = append(results, comp.Message.Content)
	}
	return results, nil
}

func buildPrompt(query string, info []string) []openai.ChatCompletionMessage {
	prompt := []openai.ChatCompletionMessage{}
	instruction := "using only the given information and a little of general knowledge, answer the question with a simple and a concise answer"
	for _, text := range info {
		prompt = append(prompt, openai.ChatCompletionMessage{
			Role:    "system",
			Content: "information: " + text,
		})
	}
	prompt = append(prompt, openai.ChatCompletionMessage{
		Role:    "system",
		Content: instruction,
	})
	prompt = append(prompt, openai.ChatCompletionMessage{
		Role:    "user",
		Content: query,
	})
	return prompt
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
