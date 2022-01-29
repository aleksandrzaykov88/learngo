package repository

import (
	"tagimagefinder/models"

	"github.com/jmoiron/sqlx"
)

type ImageFinder interface {
	Create(image models.Image) (int, error)
	Get() ([]models.Image, error)
	CreateTag(imageID int, tag string) (int, error)
	GetTagImage(tag string, page int) ([]models.Image, error)
}

type Repository struct {
	ImageFinder
}

// NewRepository is a constructor-function for Repository entity
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		ImageFinder: NewImageFinderRepository(db),
	}
}
