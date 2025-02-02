package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"idealista_tg/internal/bot"
	"idealista_tg/internal/config"
	"idealista_tg/internal/storage"
	"idealista_tg/pkg/idealista"
)

func main() {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Check if user wants to run the "rerun-scrape" command
	if len(os.Args) > 1 && os.Args[1] == "rerun-scrape" {
		rerunScrape(*cfg)
		return
	}

	// Initialize services
	idealistaClient := idealista.NewClient(idealista.ScrapeClientType)
	botService := bot.NewService(cfg, idealistaClient)
	// Start the bot

	if err := botService.Start(); err != nil {
		log.Fatal(err)
	}
}

// function that runs array of callbacks in parallel
func runCallbacks(callbacks []func() error) []error {
	wg := sync.WaitGroup{}
	errs := make([]error, len(callbacks))

	for i, callback := range callbacks {
		wg.Add(1)
		go func(i int, callback func() error) {
			defer wg.Done()
			errs[i] = callback()
		}(i, callback)
	}
	wg.Wait()
	return errs
}

// new helper function that fetches all ads from storage, re-scrapes them, and prints results
func rerunScrape(cfg config.Config) {
	storage := storage.NewGoogleSheetStorage()
	ads, err := storage.GetAds()
	if err != nil {
		log.Printf("Error reading all ads from storage: %v", err)
		return
	}
	fmt.Println("Found", len(ads), "ads in storage")

	client := idealista.NewClient(idealista.ScrapeClientType)
	callbacks := make([]func() error, 0, len(ads))

	for _, ad := range ads {
		callbacks = append(callbacks, func() error {
			newAd, scrapeErr := client.GetAd(ad.Id)
			fmt.Println("Scraped ad", newAd)
			if scrapeErr != nil {
				log.Printf("Error scraping ad %s: %v", ad.Id, scrapeErr)
				return scrapeErr
			}
			storage.UpdateAd(&newAd)
			return nil
		})
	}

	errs := runCallbacks(callbacks)
	for _, err := range errs {
		if err != nil {
			log.Printf("%v", err)
		}
	}
}
