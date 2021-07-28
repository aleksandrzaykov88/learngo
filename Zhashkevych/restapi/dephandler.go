package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//DepHandler handles rest api events for Storage.
type DepHandler struct {
	departments Departments
}

//NewDepHandler constructs the DepHandler object.
func NewDepHandler(departments Departments) *DepHandler {
	return &DepHandler{departments: departments}
}

//CreateDepartment describes rest api POST method.
func (d *DepHandler) CreateDepartment(c *gin.Context) {
	var department Department

	if err := c.BindJSON(&department); err != nil {
		fmt.Printf("failed to bind department: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	d.departments.Insert(&department)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": department.ID,
	})
}

//UpdateDepartment describes rest api PUT method.
func (d *DepHandler) UpdateDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	var department Department

	if err := c.BindJSON(&department); err != nil {
		fmt.Printf("failed to bind department: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	d.departments.Update(id, department)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": department.ID,
	})
}

//GetDepartment describes rest api GET method.
func (d *DepHandler) GetDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	department, err := d.departments.Get(id)
	if err != nil {
		fmt.Printf("failed to get department %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, department)
}

//GetDepartments also describes rest api GET method but this one return slice of all departments.
func (d *DepHandler) GetDepartments(c *gin.Context) {
	departments := d.departments.GetAll()

	c.JSON(http.StatusOK, departments)
}

//GetDepartment describes rest api DELETE method.
func (d *DepHandler) DeleteDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	d.departments.Delete(id)

	c.String(http.StatusOK, "department deleted")
}

//InsertEmployee describes rest api POST method whitch adds new employee to department.
func (d *DepHandler) InsertEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	var employee Employee
	var department Department

	if err := c.BindJSON(&employee); err != nil {
		fmt.Printf("failed to bind employee: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	d.departments.InsertEmployee(id, &employee)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": department.ID,
	})
}

//RemoveEmployee describes rest api POST method whitch adds new employee to department.
func (d *DepHandler) RemoveEmployee(c *gin.Context) {
	did, err := strconv.Atoi(c.Param("did"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	eid, err := strconv.Atoi(c.Param("eid"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	d.departments.RemoveEmployee(did, eid)

	c.String(http.StatusOK, "employee deleted")
}
