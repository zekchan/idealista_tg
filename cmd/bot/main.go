package main

import (
	"log"

	"idealista_tg/internal/bot"
	"idealista_tg/internal/config"
	"idealista_tg/pkg/idealista"
)

func main() {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize services
	idealistaClient := idealista.NewClient(idealista.ScrapeClientType)
	botService := bot.NewService(cfg, &idealistaClient)
	// Start the bot
	if err := botService.Start(); err != nil {
		log.Fatal(err)
	}
}
