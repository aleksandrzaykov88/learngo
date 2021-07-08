package main

import (
	"errors"
	"sync"
)

//Employee describes entity which characterizes the some company's employee.
type Employee struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Sex    string `json:"sex"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}

//Storage is a type for RESP API communication.
type Storage interface {
	Insert(e *Employee)
	Get(id int) (Employee, error)
	Update(id int, e Employee)
	Delete(id int)
}

//MemoryStorage is a memory-control type which reserves employees.
type MemoryStorage struct {
	counter int
	data    map[int]Employee
	sync.Mutex
}

//NewMemoryStorage constructs the MemoryStrorage object.
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data:    make(map[int]Employee),
		counter: 1,
	}
}

//Insert allows to add a new employee to the storage by his identifier.
func (s *MemoryStorage) Insert(e *Employee) {
	s.Lock()
	e.ID = s.counter
	s.data[e.ID] = *e
	s.counter++
	s.Unlock()
}

//Get employee object from store.
func (s *MemoryStorage) Get(id int) (Employee, error) {
	s.Lock()
	defer s.Unlock()
	employee, ok := s.data[id]
	if !ok {
		return employee, errors.New("employee not found")
	}
	return employee, nil
}

//Update allows to change information about employees.
func (s *MemoryStorage) Update(id int, e Employee) {
	s.Lock()
	s.data[id] = e
	s.Unlock()
}

//Delete employee from storage by his id.
func (s *MemoryStorage) Delete(id int) {
	s.Lock()
	delete(s.data, id)
	s.Unlock()
}
