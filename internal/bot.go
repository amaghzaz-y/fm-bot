package fmbot

import (
	"encoding/json"
	"log"
	"os"

	"github.com/amaghzaz-y/fm-bot/pkg/db/index"
	"github.com/ostafen/clover/v2"
	"github.com/ostafen/clover/v2/document"
	"github.com/sashabaranov/go-openai"
)

type Bot struct {
	Index  *index.VectorIndex[string]
	DB     *clover.DB
	OpenAI *openai.Client
}

func New() *Bot {
	db, err := clover.Open("db/")
	if err != nil {
		log.Fatal("error opening db:", err)
	}
	openai := openai.NewClient(os.Getenv("OPENAI"))
	return &Bot{
		nil,
		db,
		openai,
	}
}

func (b *Bot) Init() {
	// purge old collection
	if hasColl, _ := b.DB.HasCollection("vectors"); hasColl {
		b.DB.DropCollection("vectors")
		b.DB.CreateCollection("vectors")
	} else {
		b.DB.CreateCollection("vectors")
	}
	blob, err := os.ReadFile("assets/embeddings.json")
	if err != nil {
		panic(err)
	}
	var embeddings []Embedding
	err = json.Unmarshal(blob, &embeddings)
	if err != nil {
		panic(err)
	}
	var dps []*index.DataPoint[string]
	for _, embedding := range embeddings {
		id, err := b.insertEmbedding(&embedding)
		if err != nil {
			log.Println("error inserting embedding to db", err)
			continue
		}
		dp := index.NewDataPoint[string](id, embedding.Vector)
		dps = append(dps, dp)
	}
	idx, err := index.NewVectorIndex[string](1, 1536, 2, dps, index.NewCosineDistanceMeasure())
	if err != nil {
		panic(err)
	}

	idx.Build()
	b.Index = idx
}

func (b *Bot) insertEmbedding(emb *Embedding) (string, error) {
	doc := document.NewDocumentOf(emb)
	return b.DB.InsertOne("vectors", doc)
}

func (b *Bot) getEmbeddingByID(id string) (*Embedding, error) {
	doc, err := b.DB.FindById("vectors", id)
	if err != nil {
		return nil, err
	}
	var emb Embedding
	err = doc.Unmarshal(&emb)
	return &emb, err
}
