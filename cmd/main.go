package main

import (
	. "github.com/MarlikAlmighty/picbot/bot"
	. "github.com/MarlikAlmighty/picbot/models"
	"log"
	"os"
)

func main() {

	cfg := Config{}

	cfg.BotToken = os.Getenv("API_KEY")
	if cfg.BotToken == "" {
		log.Fatalf("API_KEY needed!")
	}

	cfg.WebHook = os.Getenv("WEB_HOOK")
	if cfg.WebHook == "" {
		log.Fatalf("WEB_HOOK needed!")
	}

	cfg.Host = "0.0.0.0"
	cfg.Port = os.Getenv("PORT")
	if cfg.Port == "" {
		log.Fatalf("PORT needed!")
	}

	if err := Run(&cfg); err != nil {
		log.Fatalf("Error bot run: %s \n", err)
	}
}
