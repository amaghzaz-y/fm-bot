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
	reps, err := bot.Reply("how far is the sun for earth ?")
	if err != nil {
		panic(err)
	}
	fmt.Println(reps.Replies)
}
