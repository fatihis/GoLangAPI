package main

import (
	"context"
	json "encoding/json"
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)
//types
type User struct {
	_id primitive.ObjectID
	firstname string
	id int
	lastname string
}
type UserMongo struct {
	id int `bson:"_id,omitempty"`
	firstname  string             `bson:"firstname,omitempty"`
	lastname string             `bson:"lastname,omitempty"`
}

type Person struct {
	_id string `json:”id,omitempty”`
	idNum int `json:”idnum,omitempty”`
	FirstName string `json:”firstname,omitempty”`
	LastName string `json:”lastname,omitempty”`
	Email string `json:”email,omitempty”`
	Age int `json:”age,omitempty”`
}


func main() {
	fmt.Println("go")

	/*ctx, _ := context.WithTimeout(context.Background(),10*time.Second)
	client, _ = mongo.Connect(ctx, "mongodb://localhost:27017")
	router := mux.NewRouter()
	http.ListenAndServe(":12345",router)
*/
	app := fiber.New()
	app.Get("/person/:id?", getPerson)
	app.Post("/person", createPerson)
	//app.Put("/person/:id", updatePerson)
	//app.Delete("/person/:id", deletePerson)
	port := 27027
	app.Listen(port)


	//databases, err := client.ListDatabaseNames(ctx, bson.M{})
	//fmt.Println(databases)





	/*userResult, err := userCollection.InsertOne(ctx, bson.D{
		{ "firstname", newDocument.firstname},
		{"lastname", newDocument.lastname},
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userResult.InsertedID)*/
	//add doc test
	//newDocument  := UserMongo {id:12,firstname: "sssss", lastname: "lssss"}
	//fmt.Println(createUser(newDocument,userCollection,ctx))
	//update doc test
	//updateDocument  := UserMongo {id:12,firstname: "ss", lastname: "ls"}
	//fmt.Println(editUser(updateDocument,userCollection,ctx))

	//get doc test

 //getAll(userCollection,ctx)
 //getUser(12,userCollection,ctx)
}


func GetMongoDbConnection() (*mongo.Client, error) {

	clientOptions := options.Client().
		ApplyURI("mongodb+srv://fatih:1234@cluster0.jk26p.mongodb.net/fillusers?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx,readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	return client, nil
}

func getMongoDbCollection(DbName string, CollectionName string) (*mongo.Collection, error) {
	client, err := GetMongoDbConnection()

	if err != nil {
		return nil, err
	}

	collection := client.Database(DbName).Collection(CollectionName)
	/*userDatabase := client.Database("fillusers")
	userCollection := userDatabase.Collection("users")*/

	return collection, nil
}

//REST API FUNCTIONS
func getPerson(c *fiber.Ctx) {
	collection, err := getMongoDbCollection("fillusers", "users")
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	var filter bson.M = bson.M{}

	if c.Params("id") != "" {
		id := c.Params("id")
		filter = bson.M{"idNUm": id}
	}

	var results []bson.M
	cur, err := collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	cur.All(context.Background(), &results)

	if results == nil {
		c.SendStatus(404)
		return
	}

	json, _ := json.Marshal(results)
	c.Send(json)
}
func createPerson(c *fiber.Ctx) {
	collection, err := getMongoDbCollection("fillusers", "users")
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	var person Person
	json.Unmarshal([]byte(c.Body()), &person)

	res, err := collection.InsertOne(context.Background(), person)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	response, _ := json.Marshal(res)
	c.Send(response)
}
func updatePerson(c *fiber.Ctx) {
}
func deletePerson(c *fiber.Ctx) {
}


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
	updateD := bson.D{{"firstname", user.firstname},{"lastname", user.lastname}}
	update := bson.M{
		"$set": updateD,
	}
	editUserResult, err := collection.UpdateOne(ctx,bson.M{"id":user.id},update)
	fmt.Println(editUserResult)
	return err
}
func getUser(id int, collection *mongo.Collection, ctx context.Context){
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

func getAll(collection *mongo.Collection, ctx context.Context){
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