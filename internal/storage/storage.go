package storage

import (
	"idealista_tg/pkg/idealista"
)

type Storage interface {
	SaveAd(ad *idealista.Ad, author string) error
	HasAd(id string) (bool, error)
	GetAds() ([]*idealista.Ad, error)
	UpdateAd(ad *idealista.Ad) error
}

type StorageType string

const (
	StorageTypeGoogleSheet StorageType = "googlesheet"
)

func NewStorage(storageType StorageType) Storage {
	switch storageType {
	case StorageTypeGoogleSheet:
		return NewGoogleSheetStorage()
	}
	return nil
}
