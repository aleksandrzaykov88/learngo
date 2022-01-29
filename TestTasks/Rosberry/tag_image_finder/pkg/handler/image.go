package handler

import (
	"net/http"
	"strconv"
	"tagimagefinder/models"

	"github.com/gin-gonic/gin"
)

type getImagesResponse struct {
	Data []models.Image `json:"data"`
}

// getImages gets all images
func (h *Handler) getImages(c *gin.Context) {
	// Getting query params
	images, err := h.services.ImageFinder.Get()
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getImagesResponse{
		Data: images,
	})
}

// createImage adds new image
func (h *Handler) createImage(c *gin.Context) {
	var input models.Image

	// Getting query params
	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.ImageFinder.Create(input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// createTag adds new tag to image
func (h *Handler) createTag(c *gin.Context) {
	// Getting query params
	temp, ok := c.Params.Get("id")
	if !ok {
		newErrorResponce(c, http.StatusInternalServerError, "некорректный id изображения")
		return
	}

	var (
		input   models.ImageTag
		imageID int
	)

	imageID, err := strconv.Atoi(temp)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.ImageFinder.CreateTag(imageID, input.Tag)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// getImages searching images by tags
func (h *Handler) getTagImages(c *gin.Context) {
	// Getting query params
	input := c.Query("tag")
	temp := c.Query("page")

	var (
		page int
		err  error
	)
	if temp != "" {
		page, err = strconv.Atoi(temp)
		if err != nil {
			newErrorResponce(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	images, err := h.services.ImageFinder.GetTagImage(input, page)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getImagesResponse{
		Data: images,
	})
}
