package handler

import (
	"tagimagefinder/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

// NewHandler is a constructor-function for Handler entity
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// InitRoutes initializes the routes
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		images := api.Group("/images") // Endpoints for image-methods
		{
			images.GET("/", h.getImages)
			images.GET("/search", h.getTagImages)
			images.POST("/", h.createImage)
			tags := images.Group("/tags") // Endpoints for image-tags-methods
			{
				tags.POST("/:id", h.createTag)
			}
		}
	}

	return router
}
