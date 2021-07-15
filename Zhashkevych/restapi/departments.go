package main

import (
	"errors"
	"sync"
)

//Department describes entity which characterizes the some company's department.
type Department struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Employees []Employee `json:"employees"`
}

//Departments is an interface with methods for RESP API.
type Departments interface {
	Insert(d *Department)
	Get(id int) (Department, error)
	GetAll() []Department
	Update(id int, d Department)
	Delete(id int)
	InsertEmployee(id int, e *Employee)
	RemoveEmployee(id int, empID int)
}

//MemoryDepartments is a type which implements the interface Departments.
type MemoryDepartments struct {
	counter int
	data    map[int]Department
	sync.Mutex
}

//NewMemoryDepartments constructs the MemoryDepartments object.
func NewMemoryDepartments() *MemoryDepartments {
	return &MemoryDepartments{
		data:    make(map[int]Department),
		counter: 1,
	}
}

//Insert allows to add a new department.
func (m *MemoryDepartments) Insert(d *Department) {
	m.Lock()
	d.ID = m.counter
	m.data[d.ID] = *d
	m.counter++
	m.Unlock()
}

//Get employee object from store.
func (m *MemoryDepartments) Get(id int) (Department, error) {
	m.Lock()
	defer m.Unlock()
	department, ok := m.data[id]
	if !ok {
		return department, errors.New("such department not found")
	}
	return department, nil
}

//Get all department-objects from store.
func (m *MemoryDepartments) GetAll() []Department {
	m.Lock()
	defer m.Unlock()
	departments := make([]Department, 0)
	for _, v := range m.data {
		departments = append(departments, v)
	}
	return departments
}

//Update allows to change information about department.
func (m *MemoryDepartments) Update(id int, d Department) {
	m.Lock()
	m.data[id] = d
	m.Unlock()
}

//Delete department from storage by its id.
func (m *MemoryDepartments) Delete(id int) {
	m.Lock()
	delete(m.data, id)
	m.Unlock()
}

//InsertEmployee to selected department.
func (m *MemoryDepartments) InsertEmployee(id int, e *Employee) {
	m.Lock()
	m.data[id].Employees = append(m.data[id].Employees, e)
	m.Unlock()
}

//RemoveEmployee from selected department.
func (m *MemoryDepartments) RemoveEmployee(id int) {
	m.Lock()
	delete(m.data, id)
	m.Unlock()
}
