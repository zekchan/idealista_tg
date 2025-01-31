package storage

import (
	"fmt"
	"idealista_tg/pkg/idealista"
	"testing"

	"github.com/joho/godotenv"
)

func TestGoogleSheetStorage_HasAd(t *testing.T) {
	godotenv.Load("../../.env") // TODO use separate env for tests
	storage := NewGoogleSheetStorage()
	fmt.Println(storage.HasAd("1234"))
	fmt.Println(storage.HasAd("12wefwefwef34"))
}

func TestGoogleSheetStorage_SaveAd(t *testing.T) {
	godotenv.Load("../../.env") // TODO use separate env for tests
	storage := NewGoogleSheetStorage()
	fmt.Println(storage.SaveAd(&idealista.Ad{
		Id:    "1234sdfssdfd",
		Title: "Test Ad",
		Price: 100000,
		Area:  100,
		Rooms: "T3",
	}))
}
