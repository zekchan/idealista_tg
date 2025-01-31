package idealista

import (
	"log"
	"testing"
)

func TestGetAdScrape(t *testing.T) {
	client := NewClient(ScrapeClientType)
	ad, err := client.GetAd("33878574")
	if err != nil {
		log.Printf("Error fetching ad: %v", err)
	} else {
		log.Printf("Fetched ad: %+v", ad)
	}
}
