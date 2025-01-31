package storage

import (
	"context"
	"fmt"
	"idealista_tg/pkg/idealista"
	"log"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type GoogleSheetStorage struct {
	sheetsService *sheets.Service
	spreadsheetID string
}

func (s *GoogleSheetStorage) SaveAd(ad *idealista.Ad) error {
	call := s.sheetsService.Spreadsheets.Values.Append(s.spreadsheetID, "DATABASE!A:A", &sheets.ValueRange{
		Values: [][]interface{}{{
			ad.Id,
			fmt.Sprintf("https://www.idealista.pt/imovel/%s", ad.Id),
			ad.Title,
			ad.Price,
			ad.Area,
			ad.Rooms,
			ad.Description,
		}},
	})
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
