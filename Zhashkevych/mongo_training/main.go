package main

import (
	"context"
	"log"
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

func main() {
	client := connectMongoClient()

	collection := client.Database("storage").Collection("employees")
	collection.InsertOne(context.Background(), bson.M{"name": "John2"})

	/*
		cur, err := collection.Find(context.Background(), bson.D{})
		if err != nil {
			log.Fatal(err)
		}
		defer cur.Close(context.Background())
		for cur.Next(context.Background()) {
			result := struct {
				Foo string
				Bar int32
			}{}
			err := cur.Decode(&result)
			if err != nil {
				log.Fatal(err)
			}
			raw := cur.Current
			fmt.Println(raw)
		}
		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}*/
}
