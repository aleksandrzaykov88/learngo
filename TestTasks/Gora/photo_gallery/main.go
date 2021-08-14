package main

import (
	"github.com/gin-gonic/gin"
)

//Apply setups from config-file.
var Configuration = NewConfig()

func main() {
	router := gin.Default()

	gallery := NewGallery()
	handler := NewHandler(gallery)

	router.POST("/add", handler.AddPhoto)
	router.GET("/gallery", handler.GetPhotos)
	router.DELETE("/gallery/:id", handler.DeletePhoto)

	router.Run()
}
