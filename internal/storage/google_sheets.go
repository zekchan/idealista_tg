package storage

import (
	"context"
	"fmt"
	"idealista_tg/pkg/idealista"
	"log"
	"os"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type GoogleSheetStorage struct {
	sheetsService *sheets.Service
	spreadsheetID string
}

func (s *GoogleSheetStorage) SaveAd(ad *idealista.Ad, author string) error {
	call := s.sheetsService.Spreadsheets.Values.Append(s.spreadsheetID, "DATABASE!A:A", adToValueRange(ad))
	call.ValueInputOption("USER_ENTERED")
	_, err := call.Do()
	if err != nil {
		return fmt.Errorf("unable to append ad: %v", err)
	}
	return nil
}

func (s *GoogleSheetStorage) HasAd(id string) (bool, error) {
	resp, err := s.sheetsService.Spreadsheets.Values.Get(s.spreadsheetID, "DATABASE!A:A").Do()
	if err != nil {
		return false, fmt.Errorf("unable to read sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		return false, nil
	}

	// Check if ID exists in first column
	for _, row := range resp.Values {
		if len(row) > 0 {
			if row[0].(string) == id {
				return true, nil
			}
		}
	}

	return false, nil
}

func (s *GoogleSheetStorage) GetAds() ([]*idealista.Ad, error) {
	resp, err := s.sheetsService.Spreadsheets.Values.Get(s.spreadsheetID, "DATABASE!A:A").Do()
	if err != nil {
		return nil, fmt.Errorf("unable to read sheet: %v", err)
	}

	ads := make([]*idealista.Ad, 0)
	for _, row := range resp.Values {
		if len(row) > 0 {
			ads = append(ads, &idealista.Ad{Id: row[0].(string)})
		}
	}
	return ads, nil
}
func adToValueRange(ad *idealista.Ad) *sheets.ValueRange {
	return &sheets.ValueRange{
		Values: [][]interface{}{{
			ad.Id,
			time.Now().Format(time.RFC3339),
			ad.Title,
			ad.Price,
			ad.Area,
			ad.Rooms,
			ad.Description,
			ad.ImageURL,
		}},
	}
}
func (s *GoogleSheetStorage) UpdateAd(ad *idealista.Ad) error {
	// find row in sheet
	resp, err := s.sheetsService.Spreadsheets.Values.Get(s.spreadsheetID, "DATABASE!A:A").Do()
	if err != nil {
		return fmt.Errorf("unable to read sheet: %v", err)
	}

	for index, row := range resp.Values {
		if len(row) > 0 {
			if row[0].(string) == ad.Id {
				fmt.Println("Updating ad", ad.Id, fmt.Sprintf("DATABASE!A%v:Z%v", index+1, index+1))
				// update row
				call := s.sheetsService.Spreadsheets.Values.Update(s.spreadsheetID, fmt.Sprintf("DATABASE!A%v:Z%v", index+1, index+1), adToValueRange(ad))
				call.ValueInputOption("RAW")
				_, err := call.Do()
				if err != nil {
					return fmt.Errorf("unable to update ad: %v", err)
				}
				return nil
			}
		}
	}
	return fmt.Errorf("ad not found")
}

func NewGoogleSheetStorage() *GoogleSheetStorage {
	sheetsService, err := sheets.NewService(context.Background(), option.WithCredentialsFile(os.Getenv("GOOGLE_CREDENTIALS")))

	if err != nil {
		log.Fatalf("Unable to create sheets client: %v", err)
	}

	return &GoogleSheetStorage{
		sheetsService: sheetsService,
		spreadsheetID: os.Getenv("GOOGLE_SPREADSHEET_ID"),
	}
}

var _ Storage = (*GoogleSheetStorage)(nil) // Ensure GoogleSheetStorage implements the Storage interface
