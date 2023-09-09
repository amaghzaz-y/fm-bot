package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/amaghzaz-y/fm-bot/internal"
	"github.com/amaghzaz-y/fm-bot/pkg/vector-db/index"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	idx, err := index.NewVectorIndex[string](1, 1536, 2, GetDataPoints(), index.NewCosineDistanceMeasure())
	if err != nil {
		panic(err)
	}
	idx.Build()
	x, err := internal.GetEmbedding("robot limites de dimensions")
	if err != nil {
		panic(err)
	}
	res, err := idx.SearchByVector(x, 10, 10)
	if err != nil {
		panic(err)
	}
	for _, r := range res {
		fmt.Println(r.ID)
	}
}

func GetDataPoints() []*index.DataPoint[string] {
	blob, err := os.ReadFile("assets/embeddings.json")
	if err != nil {
		panic(err)
	}
	var embeddings []internal.Embedding
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
