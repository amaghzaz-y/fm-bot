package internal

type Embedding struct {
	Text   string    `json:"text"`
	Vector []float32 `json:"vector"`
}
