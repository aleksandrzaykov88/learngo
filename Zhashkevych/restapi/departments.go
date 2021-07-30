package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
)

//Department describes entity which characterizes the some company's department.
type Department struct {
	ID        int              `json:"id"`
	Name      string           `json:"name"`
	Employees map[int]Employee `json:"employees"`
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

//MongoDepts is a collection of departments for work with Mongo DB.
type MongoDepts struct {
	data map[int]*Department
}

//NewMongoDepts constructs the MongoDepts object.
func NewMongoDepts() *MongoDepts {
	return &MongoDepts{
		data: make(map[int]*Department),
	}
}

//Insert allows to add a new department to the DB.
func (m *MongoDepts) Insert(d *Department) {
	collection := connectDB("storage", "depts")
	ctx, _ := getCTX(10)

	var temp bson.M
	if checkID := collection.FindOne(ctx, bson.M{"id": d.ID}).Decode(&temp); checkID == nil {
		log.Fatal(errors.New("There is already such an element"))
	}

	b, err := bson.Marshal(&d)
	if err != nil {
		log.Fatal(err)
	}
	collection.InsertOne(context.Background(), b)
}

//Get department object from DB.
func (m *MongoDepts) Get(id int) (Department, error) {
	collection := connectDB("storage", "depts")
	ctx, _ := getCTX(10)

	var departmentEntry bson.M
	if checkID := collection.FindOne(ctx, bson.M{"id": id}).Decode(&departmentEntry); checkID != nil {
		log.Fatal(errors.New("There is no such an element"))
	}

	bsonBytes, _ := bson.Marshal(departmentEntry)
	var n Department
	bson.Unmarshal(bsonBytes, &n)
	return n, nil
}

//Get all departments from DB.
func (m *MongoDepts) GetAll() []Department {
	collection := connectDB("storage", "depts")
	ctx, _ := getCTX(10)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var n []Department
	err = cursor.All(ctx, &n)
	if err != nil {
		log.Fatal(err)
	}

	return n
}

//Update allows to change information about departments in DB.
func (m *MongoDepts) Update(id int, d Department) {
	collection := connectDB("storage", "depts")
	ctx, _ := getCTX(10)

	var departmentEntry bson.M
	if checkID := collection.FindOne(ctx, bson.M{"id": id}).Decode(&departmentEntry); checkID != nil {
		log.Fatal(errors.New("There is no such an element"))
	}

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"id": id},
		bson.D{
			{"$set", bson.D{{"name", d.Name}, {"employees", d.Employees}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

//Delete department from DB.
func (m *MongoDepts) Delete(id int) {
	collection := connectDB("storage", "depts")
	ctx, _ := getCTX(10)

	var departmentEntry bson.M
	if checkID := collection.FindOne(ctx, bson.M{"id": id}).Decode(&departmentEntry); checkID != nil {
		log.Fatal(errors.New("There is no such an element"))
	}

	_, err := collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		log.Fatal(err)
	}
}

//InsertEmployee to selected department in DB.
func (m *MongoDepts) InsertEmployee(id int, e *Employee) {
	collection := connectDB("storage", "depts")
	ctx, _ := getCTX(10)

	var departmentEntry bson.M
	if checkID := collection.FindOne(ctx, bson.M{"id": id}).Decode(&departmentEntry); checkID != nil {
		log.Fatal(errors.New("There is no such an element"))
	}

	bsonBytes, _ := bson.Marshal(departmentEntry)
	var d Department
	bson.Unmarshal(bsonBytes, &d)

	ID := len(d.Employees) + 1

	d.Employees[ID] = *e

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"id": id},
		bson.D{
			{"$set", bson.D{{"employees", d.Employees}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

//RemoveEmployee from selected department in DB.
func (m *MongoDepts) RemoveEmployee(depid int, eid int) {
	collection := connectDB("storage", "depts")
	ctx, _ := getCTX(10)

	var departmentEntry bson.M
	if checkID := collection.FindOne(ctx, bson.M{"id": depid}).Decode(&departmentEntry); checkID != nil {
		log.Fatal(errors.New("There is no such an element"))
	}

	bsonBytes, _ := bson.Marshal(departmentEntry)
	var d Department
	bson.Unmarshal(bsonBytes, &d)

	delete(d.Employees, eid)

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"id": depid},
		bson.D{
			{"$set", bson.D{{"employees", d.Employees}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

//STRUCT IMPLEMENTATION
//MemoryDepartments is a type which implements the interface Departments.
type MemoryDepartments struct {
	counter int
	data    map[int]*Department
	sync.Mutex
}

//NewMemoryDepartments constructs the MemoryDepartments object.
func NewMemoryDepartments() *MemoryDepartments {
	return &MemoryDepartments{
		data:    make(map[int]*Department),
		counter: 1,
	}
}

//Insert allows to add a new department.
func (m *MemoryDepartments) Insert(d *Department) {
	m.Lock()
	d.ID = m.counter
	m.data[d.ID] = d
	m.counter++
	m.Unlock()
}

//Get employee object from store.
func (m *MemoryDepartments) Get(id int) (Department, error) {
	m.Lock()
	defer m.Unlock()
	department, ok := m.data[id]
	if !ok {
		department := Department{}
		return department, errors.New("such department not found")
	}
	return *department, nil
}

//Get all department-objects from store.
func (m *MemoryDepartments) GetAll() []Department {
	m.Lock()
	defer m.Unlock()
	departments := make([]Department, 0)
	for _, v := range m.data {
		departments = append(departments, *v)
	}
	return departments
}

//Update allows to change information about department.
func (m *MemoryDepartments) Update(id int, d Department) {
	m.Lock()
	m.data[id] = &d
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
	ID := len(m.data[id].Employees) + 1
	m.data[id].Employees[ID] = *e
	m.Unlock()
}

//RemoveEmployee from selected department.
func (m *MemoryDepartments) RemoveEmployee(depid int, eid int) {
	m.Lock()
	delete(m.data[depid].Employees, eid)
	m.Unlock()
}
