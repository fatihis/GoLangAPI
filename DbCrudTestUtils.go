package main

import (
	"context"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type User struct {
	_id       primitive.ObjectID
	firstname string
	id        int
	lastname  string
}
type UserMongo struct {
	id        int    `bson:"_id,omitempty"`
	firstname string `bson:"firstname,omitempty"`
	lastname  string `bson:"lastname,omitempty"`
}

//DB CRUD UTILS
func createUser(user UserMongo, collection *mongo.Collection, ctx context.Context) error {
	createUserResult, err := collection.InsertOne(ctx, bson.D{
		{"id", user.id},
		{"firstname", user.firstname},
		{"lastname", user.lastname},
	})
	fmt.Println(createUserResult)
	return err
}

func editUser(user UserMongo, collection *mongo.Collection, ctx context.Context) error {
	updateD := bson.D{{"firstname", user.firstname}, {"lastname", user.lastname}}
	update := bson.M{
		"$set": updateD,
	}
	editUserResult, err := collection.UpdateOne(ctx, bson.M{"id": user.id}, update)
	fmt.Println(editUserResult)
	return err
}
func getUser(id int, collection *mongo.Collection, ctx context.Context) {
	filterCursor, err := collection.Find(ctx, bson.M{"id": id})
	if err != nil {
		log.Fatal(err)
	}
	var user []bson.M
	if err = filterCursor.All(ctx, &user); err != nil {
		log.Fatal(err)
	}
	var foundUser User

	mapstructure.Decode(user[0], &foundUser)
	fmt.Println(foundUser)
}

func getAll(collection *mongo.Collection, ctx context.Context) {
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var documents []bson.M
	if err = cursor.All(ctx, &documents); err != nil {
		log.Fatal(err)
	}
	fmt.Println(documents)
}
