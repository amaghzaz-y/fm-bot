package fmbot

type Embedding struct {
	Text   string    `json:"text" clover:"text"`
	Vector []float32 `json:"vector" clover:"vector"`
}
