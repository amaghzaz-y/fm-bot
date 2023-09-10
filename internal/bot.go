package internal

import (
	"encoding/json"
	"log"
	"os"

	"github.com/amaghzaz-y/fm-bot/pkg/db/index"
	"github.com/ostafen/clover/v2"
	"github.com/sashabaranov/go-openai"
)

type Bot struct {
	Index  *index.VectorIndex[string]
	DB     *clover.DB
	OpenAI *openai.Client
}

func New() *Bot {
	db, err := clover.Open("bot.db")
	if err != nil {
		log.Fatal(err)
	}
	openai := openai.NewClient(os.Getenv("OPENAI"))
	idx := indexBuild()
	return &Bot{
		idx,
		db,
		openai,
	}
}

// builds the index
func indexBuild() *index.VectorIndex[string] {
	idx, err := index.NewVectorIndex[string](1, 1536, 2, getDataPoints(), index.NewCosineDistanceMeasure())
	if err != nil {
		panic(err)
	}
	idx.Build()
	return idx
}

func getDataPoints() []*index.DataPoint[string] {
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
		dp := index.NewDataPoint[string](embedding.Text, embedding.Vector)
		dps = append(dps, dp)
	}
	return dps
}
