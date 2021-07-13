package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//connectMongoClient connects to MongoDB and returns client-entity;
func connectMongoClient() *mongo.Client {
	context.WithTimeout(context.Background(), 20*time.Second)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

//Employee describes entity which characterizes the some company's employee.
type Employee struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Sex    string `json:"sex"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}

//Storage is an interface with methods for RESP API communication.
type Storage interface {
	Insert(e *Employee)
	Get(id int) (Employee, error)
	Update(id int, e Employee)
	Delete(id int)
}

//MongoStorage is a type which implements Storage for work with Mongo DB.
type MongoStorage struct {
	data map[int]Employee
}

//NewMongoStorage constructs the MongoStorage object.
func NewMongoStorage() *MongoStorage {
	return &MongoStorage{
		data: make(map[int]Employee),
	}
}

//Insert allows to add a new employee to the DB.
func (s *MongoStorage) Insert(e *Employee) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := connectMongoClient()
	defer client.Disconnect(ctx)
	collection := client.Database("storage").Collection("employees")

	var temp bson.M
	if checkID := collection.FindOne(ctx, bson.M{"id": e.ID}).Decode(&temp); checkID == nil {
		log.Fatal(errors.New("There is already such an element"))
	}

	b, err := bson.Marshal(&e)
	if err != nil {
		log.Fatal(err)
	}
	collection.InsertOne(context.Background(), b)
}

//Get employee object from DB.
func (s *MongoStorage) Get(id int) (Employee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := connectMongoClient()
	defer client.Disconnect(ctx)
	collection := client.Database("storage").Collection("employees")

	var employeeEntry bson.M
	if checkID := collection.FindOne(ctx, bson.M{"id": id}).Decode(&employeeEntry); checkID != nil {
		log.Fatal(errors.New("There is no such an element"))
	}

	bsonBytes, _ := bson.Marshal(employeeEntry)
	var n Employee
	bson.Unmarshal(bsonBytes, &n)
	return n, nil
}

//Update allows to change information about employees in DB.
func (s *MongoStorage) Update(id int, e Employee) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := connectMongoClient()
	defer client.Disconnect(ctx)
	collection := client.Database("storage").Collection("employees")

	var employeeEntry bson.M
	if checkID := collection.FindOne(ctx, bson.M{"id": id}).Decode(&employeeEntry); checkID != nil {
		log.Fatal(errors.New("There is no such an element"))
	}

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"id": id},
		bson.D{
			{"$set", bson.D{{"name", e.Name}, {"sex", e.Sex}, {"age", e.Age}, {"salary", e.Salary}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

//Delete employee from DB.
func (s *MongoStorage) Delete(id int) {
	delete(s.data, id)
}

//MemoryStorage is a type which reserves employees.
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
