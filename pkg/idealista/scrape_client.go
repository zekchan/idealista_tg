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
func getAdPrice(doc *goquery.Document) int {
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
	return price
}

func getAdTitle(doc *goquery.Document) string {
	title := doc.Find("head>title").First().Text()
	title = strings.TrimSpace(title)
	return title
}

func getAdArea(doc *goquery.Document) int {
	area := 0
	areaText := doc.Find("div.main-info>p.info-data>span.price-container+span>span").First().Text()
	areaText = strings.TrimSpace(areaText)

	if areaText != "" {
		if p, err := strconv.Atoi(areaText); err == nil {
			area = int(p)
		}
	}
	return area
}
func getAdImageURL(doc *goquery.Document) string {
	imageURLs := doc.Find("head>meta[property='og:image']").First().AttrOr("content", "")
	return imageURLs
}
func getAdRooms(doc *goquery.Document) string {
	roomsText := doc.Find("div.main-info>p.info-data>span.price-container+span+span>span").First().Text()
	roomsText = strings.TrimSpace(roomsText)
	return roomsText
}
func getAdDescription(doc *goquery.Document) string {
	description := doc.Find("div.comment>div>p").First().Text()
	description = strings.TrimSpace(description)
	return description
}

func (c *ScrapeClient) GetAd(id string) (Ad, error) {
	htmlReader, err := getHtml(fmt.Sprintf("https://www.idealista.pt/imovel/%s", id))
	defer htmlReader.Close()
	doc, err := goquery.NewDocumentFromReader(htmlReader)
	if err != nil {
		return Ad{}, err
	}

	return Ad{Id: id,
		Price:       getAdPrice(doc),
		Title:       getAdTitle(doc),
		Area:        getAdArea(doc),
		Rooms:       getAdRooms(doc),
		Description: getAdDescription(doc),
		ImageURL:    getAdImageURL(doc),
	}, nil
}

var _ Client = (*ScrapeClient)(nil) // Ensure ScrapeClient implements the Client interface
