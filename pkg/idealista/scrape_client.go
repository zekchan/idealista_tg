package idealista

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ScrapeClient struct{}

func getHtml(url string) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (c *ScrapeClient) GetAd(id string) (Ad, error) {
	htmlReader, err := getHtml(fmt.Sprintf("https://www.idealista.pt/imovel/%s", id))
	defer htmlReader.Close()
	doc, err := goquery.NewDocumentFromReader(htmlReader)
	if err != nil {
		return Ad{}, err
	}

	price := 0
	priceText := doc.Find("div.info-data>span.info-data-price>span.txt-bold").First().Text()
	priceText = strings.TrimSpace(priceText)
	priceText = strings.NewReplacer(
		"€", "",
		".", "",
		",", ".",
	).Replace(priceText)

	if priceText != "" {
		if p, err := strconv.ParseFloat(priceText, 64); err == nil {
			price = int(p)
		}

	}
	return Ad{Id: id, Price: price}, nil
}

var _ Client = (*ScrapeClient)(nil) // Ensure ScrapeClient implements the Client interface
