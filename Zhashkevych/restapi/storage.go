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
	GetAll() []Employee
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

//Insert allows to add a new employee to the DB.
func (s *MongoStorage) Insert(e *Employee) {
	collection := connectDB()
	ctx, _ := getCTX(10)

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
	collection := connectDB()
	ctx, _ := getCTX(10)

	var employeeEntry bson.M
	if checkID := collection.FindOne(ctx, bson.M{"id": id}).Decode(&employeeEntry); checkID != nil {
		log.Fatal(errors.New("There is no such an element"))
	}

	bsonBytes, _ := bson.Marshal(employeeEntry)
	var n Employee
	bson.Unmarshal(bsonBytes, &n)
	return n, nil
}

//Get all employees from DB.
func (s *MongoStorage) GetAll() []Employee {
	collection := connectDB()
	ctx, _ := getCTX(10)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var n []Employee
	err = cursor.All(ctx, &n)
	if err != nil {
		log.Fatal(err)
	}

	return n
}

//Update allows to change information about employees in DB.
func (s *MongoStorage) Update(id int, e Employee) {
	collection := connectDB()
	ctx, _ := getCTX(10)

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
	collection := connectDB()
	ctx, _ := getCTX(10)

	var employeeEntry bson.M
	if checkID := collection.FindOne(ctx, bson.M{"id": id}).Decode(&employeeEntry); checkID != nil {
		log.Fatal(errors.New("There is no such an element"))
	}

	_, err := collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		log.Fatal(err)
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

//Get all employees objects from store.
func (s *MemoryStorage) GetAll() []Employee {
	s.Lock()
	defer s.Unlock()
	employees := make([]Employee, 0)
	for _, v := range s.data {
		employees = append(employees, v)
	}
	return employees
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

//connectMongoClient connects to MongoDB and returns client-entity;
func connectMongoClient() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := getCTX(20)
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

//connectDB connects and returns employees DB.
func connectDB() *mongo.Collection {
	client := connectMongoClient()
	collection := client.Database("storage").Collection("employees")
	return collection
}

//getCTX returns ctx-value.
func getCTX(seconds int) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(seconds)*time.Second)
}
