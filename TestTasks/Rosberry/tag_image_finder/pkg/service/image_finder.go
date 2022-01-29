package service

import (
	"tagimagefinder/models"
	"tagimagefinder/pkg/repository"
)

type ImageFinderService struct {
	repo repository.ImageFinder
}

// NewImageFInderService is a constructor-function for ImageFinderService entity
func NewImageFInderService(repo repository.ImageFinder) *ImageFinderService {
	return &ImageFinderService{repo: repo}
}

// Create adds new image
func (s *ImageFinderService) Create(image models.Image) (int, error) {
	return s.repo.Create(image)
}

// Get gets all images
func (s *ImageFinderService) Get() ([]models.Image, error) {
	return s.repo.Get()
}

// CreateTag adds new tag to image
func (s *ImageFinderService) CreateTag(imageID int, tag string) (int, error) {
	return s.repo.CreateTag(imageID, tag)
}

// GetTagImage searching images by tags
func (s *ImageFinderService) GetTagImage(tag string, page int) ([]models.Image, error) {
	return s.repo.GetTagImage(tag, page)
}
