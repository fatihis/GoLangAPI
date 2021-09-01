package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

//UTILS
//FUNC:Returns *mongo.Client
func GetMongoDbConnection() (*mongo.Client, error) {

	clientOptions := options.Client().
		ApplyURI("mongodb+srv://fatih:1234@cluster0.jk26p.mongodb.net/fillusers?retryWrites=true&w=majority") //Applied connection string retrieved by MongoDB Atlas
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions) //Establish connection
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary()) //Ping MongoDB Atlas servers
	if err != nil {
		log.Fatal(err)
	}
	return client, nil //Return connection client
}

//FUNC:Returns instance of desired collection of DB
func getMongoDbCollection(DbName string, CollectionName string) (*mongo.Collection, error) {
	client, err := GetMongoDbConnection() //Get *mongo.client

	if err != nil {
		return nil, err
	}

	collection := client.Database(DbName).Collection(CollectionName)
	return collection, nil
}
