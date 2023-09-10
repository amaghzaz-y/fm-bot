package main

import (
	"github.com/amaghzaz-y/fm-bot/api"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	api := api.New()
	api.Start()
}
