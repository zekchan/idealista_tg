package storage

import (
	"idealista_tg/pkg/idealista"
)

type Storage interface {
	SaveAd(ad *idealista.Ad) error
	HasAd(id string) (bool, error)
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
