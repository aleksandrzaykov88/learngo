package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"tagimagefinder/models"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type ImageFinderRepository struct {
	db *sqlx.DB
}

// NewImageFinderRepository is a constructor-function for ImageFinderRepository entity
func NewImageFinderRepository(db *sqlx.DB) *ImageFinderRepository {
	return &ImageFinderRepository{db: db}
}

// Create adds new image
func (r *ImageFinderRepository) Create(image models.Image) (int, error) {
	tx, err := r.db.Begin() // Starts transaction
	if err != nil {
		return 0, err
	}

	var id int // Variable to get result from db-query
	createImageQuery := fmt.Sprintf("select * from %s ($1, $2)", fnImageIns)
	// Creates new image in db with link and alt params and returns its id
	row := tx.QueryRow(createImageQuery, image.Link, image.Alt)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit() // Ends transaction
}

// Get gets all images
func (r *ImageFinderRepository) Get() ([]models.Image, error) {
	// Gers all images from view
	query := fmt.Sprintf("select * from %s", vImage)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var images []models.Image // Result slice
	for rows.Next() {
		var img models.Image
		// Scanning query to relevant fields
		if err := rows.Scan(&img.ID, &img.Link, &img.Alt, &img.CreatedAt, &img.UpdatedAt, &img.Tags); err != nil {
			return nil, err
		}

		if img.Tags != "" { // Tags need to be unmarshalled
			var m json.RawMessage
			err := json.Unmarshal(img.Tags.([]byte), &m)
			if err != nil {
				return nil, err
			}
			img.Tags = m
		}
		images = append(images, img) // Adds image to result slice
	}

	return images, err
}

// CreateTag adds new tag to image
func (r *ImageFinderRepository) CreateTag(imageID int, tag string) (int, error) {
	tx, err := r.db.Begin() // Starts transaction
	if err != nil {
		return 0, err
	}

	var id int // Variable to get result from db-query
	// Adds tag to image with imageID
	createImageQuery := fmt.Sprintf("select * from %s ($1, $2)", fnImageTagIns)
	row := tx.QueryRow(createImageQuery, imageID, tag)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit() // Ends transaction
}

var ( // Auxiliary variables for GetTagImage func
	images  []models.Image // Global slice for storage images from query
	saveTag string         // Global string for storage tag from last query
)

// GetTagImage searching images by tags
func (r *ImageFinderRepository) GetTagImage(tag string, page int) ([]models.Image, error) {
	// Get param from config-file
	// Variable items shows the number of images per page
	items, err := strconv.Atoi(viper.GetString("images.items"))
	if err != nil {
		return nil, err
	}

	// Find all images for input tag
	query := fmt.Sprintf("select * from %s ($1)", fnImageTagGet)
	rows, err := r.db.Query(query, tag)
	if err != nil {
		return nil, err
	}

	var result []models.Image                            // Result array
	if len(images) == 0 || saveTag != tag || page == 0 { // If the tag has changed or the image array is empty
		images = nil      // Clear image slice
		for rows.Next() { // Read query response
			var img models.Image
			// Scanning query to relevant fields
			if err := rows.Scan(&img.ID, &img.Link, &img.Alt, &img.CreatedAt, &img.UpdatedAt, &img.Tags); err != nil {
				return nil, err
			}

			if img.Tags != "" { // Tags need to be unmarshalled
				var m json.RawMessage
				err := json.Unmarshal(img.Tags.([]byte), &m)
				if err != nil {
					return nil, err
				}
				img.Tags = m
			}
			images = append(images, img) // Adds image to global image-slice
		}
	}

	saveTag = tag // Saving the current tag to a global variable

	if items > len(images) { // If there are more elements on the page than there are elements in the query result
		items = len(images)
	}
	if page != 0 { // If there is pagination in the request
		temp := float64(len(images)) / float64(items)
		pages := int(math.Ceil(temp)) // Determine the number of pages
		if page < 0 || page > pages {
			return nil, errors.New("такой страницы не существует")
		}
		// Determine the limits of image-slice for every page
		from := (page - 1) * items
		to := (page-1)*items + items
		if to > len(images) {
			to = len(images)
		}
		result = images[from:to]
	} else {
		result = images[0:items] // If the page variable is not specified, the first page is returned.
	}

	return result, err
}
