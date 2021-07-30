package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	memoryStorage := NewMemoryStorage()
	handler := NewHandler(memoryStorage)
	//mongoStorage := NewMongoStorage()
	//handler := NewHandler(mongoStorage)

	router := gin.Default()

	//Emploee main API.
	router.POST("/employee", handler.CreateEmployee)
	router.GET("/employee/:id", handler.GetEmployee)
	router.PUT("/employee/:id", handler.UpdateEmployee)
	router.DELETE("/employee/:id", handler.DeleteEmployee)
	//Employee addition API.
	router.GET("/employee/", handler.GetEmployees)

	//memDepts := NewMemoryDepartments()
	//dHandler := NewDepHandler(memDepts)
	mongoDepts := NewMongoDepts()
	dHandler := NewDepHandler(mongoDepts)

	//Department main API.
	router.POST("/department", dHandler.CreateDepartment)
	router.GET("/department/:id", dHandler.GetDepartment)
	router.PUT("/department/:id", dHandler.UpdateDepartment)
	router.DELETE("/department/:id", dHandler.DeleteDepartment)
	//Department addition API.
	router.GET("/department/", dHandler.GetDepartments)
	router.POST("/department/:id", dHandler.InsertEmployee)
	router.DELETE("/department/remove/:did/:eid", dHandler.RemoveEmployee)

	router.Run()
}
