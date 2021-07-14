package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//memoryStorage := NewMemoryStorage()
	//handler := NewHandler(memoryStorage)
	mongoStorage := NewMongoStorage()
	handler := NewHandler(mongoStorage)

	router := gin.Default()

	//Emploee main API
	router.POST("/employee", handler.CreateEmployee)
	router.GET("/employee/:id", handler.GetEmployee)
	router.PUT("/employee/:id", handler.UpdateEmployee)
	router.DELETE("/employee/:id", handler.DeleteEmployee)
	//Employee addition API
	router.GET("/employee/", handler.GetEmployees)

	router.Run()
}
