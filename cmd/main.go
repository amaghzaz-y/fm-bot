package main

import (
	"fmt"

	fmbot "github.com/amaghzaz-y/fm-bot/internal"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	bot := fmbot.New()
	bot.Init()
	reps, err := bot.Reply("les PAMIS")
	if err != nil {
		panic(err)
	}
	fmt.Println(reps.Reply)
}
