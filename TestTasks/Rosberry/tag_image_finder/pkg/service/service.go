package service

import (
	"tagimagefinder/models"
	"tagimagefinder/pkg/repository"
)

type ImageFinder interface {
	Create(image models.Image) (int, error)
	Get() ([]models.Image, error)
	CreateTag(imageID int, tag string) (int, error)
	GetTagImage(tag string, page int) ([]models.Image, error)
}

type Service struct {
	ImageFinder
}

// NewService is a constructor-function for Service entity
func NewService(repos *repository.Repository) *Service {
	return &Service{
		ImageFinder: NewImageFInderService(repos.ImageFinder),
	}
}
