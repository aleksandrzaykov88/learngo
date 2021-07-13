package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Employee struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Sex    string `json:"sex"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}

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

func main() {
	client := connectMongoClient()
	collection := client.Database("storage").Collection("employees")

	//ADD
	newEmployee := Employee{12, "Michael", "Male", 22, 36000}
	b, err := bson.Marshal(newEmployee)
	if err != nil {
		log.Fatal(err)
	}
	collection.InsertOne(context.Background(), b)

	/*
		//GET
		cur, err := collection.Find(context.Background(), bson.D{})
		if err != nil {
			log.Fatal(err)
		}
		defer cur.Close(context.Background())
		for cur.Next(context.Background()) {
			raw := cur.Current
			fmt.Println(raw)
		}
		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}
	*/
}
