package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//ErrorResponse stores the error message.
type ErrorResponse struct {
	Message string `json:"message"`
}

//Handler handles rest api events for Storage.
type Handler struct {
	storage Storage
}

//NewHandler constructs the handler object.
func NewHandler(storage Storage) *Handler {
	return &Handler{storage: storage}
}

//CreateEmployee describes rest api POST method.
func (h *Handler) CreateEmployee(c *gin.Context) {
	var employee Employee

	if err := c.BindJSON(&employee); err != nil {
		fmt.Printf("failed to bind employee: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	h.storage.Insert(&employee)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": employee.ID,
	})
}

//UpdateEmployee describes rest api PUT method.
func (h *Handler) UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	var employee Employee

	if err := c.BindJSON(&employee); err != nil {
		fmt.Printf("failed to bind employee: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	h.storage.Update(id, employee)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": employee.ID,
	})
}

//GetEmployee describes rest api GET method.
func (h *Handler) GetEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	employee, err := h.storage.Get(id)
	if err != nil {
		fmt.Printf("failed to get employee %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, employee)
}

//GetEmployee describes rest api DELETE method.
func (h *Handler) DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	h.storage.Delete(id)

	c.String(http.StatusOK, "employee deleted")
}
