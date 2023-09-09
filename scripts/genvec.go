package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	text, err := os.ReadFile("assets/rules.txt")
	if err != nil {
		panic(err)
	}
	parts := splitPartN(string(text))
	fmt.Println(len(parts))
}

type Embedding struct {
	Text   string
	Vector float32
}

func splitPartN(text string) []string {
	words := strings.Split(text, ".")
	fmt.Println("tokens:", len(strings.Fields(text)))
	var result []string
	for _, word := range words {
		if len(word) > 4 {
			result = append(result, word)
		}
	}
	return result
}
