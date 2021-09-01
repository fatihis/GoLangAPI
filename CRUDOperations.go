package main

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//REST API FUNCTIONS
/*HTTP GET: Get single person <<START>> */
func getPerson(c *fiber.Ctx) {
	collection, err := getMongoDbCollection("fillusers", "users") //get MongoDB connection : users collection
	if err != nil { //checksum connection
		c.Status(500).Send(err) //Return 500 Internal Server Error in order to !(error)!
		return
	}

	var filter bson.M = bson.M{}


	if c.Params("id") == "" {
		c.Status(400) //Return 400 Bad Request if "_id" query parameter empty
		return
	}

	id := c.Params("id")
	objID, _ := primitive.ObjectIDFromHex(id) //convert int id type into MongoDB related type ObjectID
	filter = bson.M{"_id": objID} // construct filter, _id == (id query parameter)

	var results []bson.M
	cur, err := collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())

	if err != nil {
		c.Status(500).Send(err) //Return 500 Internal Server Error in order to !(error)!
		return
	}

	cur.All(context.Background(), &results)

	if results == nil {
		c.SendStatus(404) //Return 404 Not Found if data can not found
		return
	}

	json, _ := json.Marshal(results) //Convert response into JSON
	c.Send(json) //Return JSON formatted response
}
/*<<END>> HTTP GET: Get single person <<END>> */

/*HTTP GET: Get all persons <<START>> */
func getAllPersons(c *fiber.Ctx) {
	collection, err := getMongoDbCollection("fillusers", "users")  //get MongoDB connection : users collection
	if err != nil {
		c.Status(500).Send(err)  //Return 500 Internal Server Error in order to !(error)!
		return
	}

	var results []bson.M
	cur, err := collection.Find(context.Background(), bson.M{})
	defer cur.Close(context.Background())

	if err != nil {
		c.Status(500).Send(err)  //Return 500 Internal Server Error in order to !(error)!
		return
	}

	cur.All(context.Background(), &results)

	if results == nil {
		c.SendStatus(404) //Return 404 Not Found if data can not found
		return
	}

	json, _ := json.Marshal(results) //Convert response into JSON
	c.Send(json) //Return JSON formatted response
}
/*<<END>> HTTP GET: Get all persons <<END>> */

/*HTTP POST: Create person <<START>> */
func createPerson(c *fiber.Ctx) {
	collection, err := getMongoDbCollection("fillusers", "users")  //get MongoDB connection : users collection
	if err != nil {
		c.Status(500).Send(err) //Return 500 Internal Server Error in order to !(error)!
		return
	}

	var person Person
	json.Unmarshal([]byte(c.Body()), &person)

	results, err := collection.InsertOne(context.Background(), person)
	if err != nil {
		c.Status(500).Send(err) //Return 500 Internal Server Error in order to !(error)!
		return
	}

	response, _ := json.Marshal(results) //Convert response into JSON
	c.Send(response) //Return JSON formatted response
}
/*<<END>> HTTP POST: Create person <<END>> */

/*HTTP PUT: Update person <<START>> */
func updatePerson(c *fiber.Ctx) {
	collection, err := getMongoDbCollection("fillusers", "users")  //get MongoDB connection : users collection
	if err != nil {
		c.Status(500).Send(err) //Return 500 Internal Server Error in order to !(error)!
		return
	}
	var person Person
	json.Unmarshal([]byte(c.Body()), &person) //Get data from JSON body into person object

	update := bson.M{
		"$set": person,
	}

	objID, _ := primitive.ObjectIDFromHex(c.Params("id")) //convert int id type into MongoDB related type ObjectID
	results, err := collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update) //Filter "_id": (id query parameter)

	if err != nil {
		c.Status(500).Send(err) //Return 500 Internal Server Error in order to !(error)!
		return
	}

	response, _ := json.Marshal(results) //Convert response into JSON
	c.Send(response) //Return JSON formatted response
}
/*<<END>> HTTP PUT: Create person <<END>> */

/*HTTP DELETE: Delete person <<START>> */
func deletePerson(c *fiber.Ctx) {

	collection, err := getMongoDbCollection("fillusers", "users")  //get MongoDB connection : users collection

	if err != nil {
		c.Status(500).Send(err) //Return 500 Internal Server Error in order to !(error)!
		return
	}

	objID, _ := primitive.ObjectIDFromHex(c.Params("id")) //convert int id type into MongoDB related type ObjectID
	results, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID}) //Filter "_id": (id query parameter)

	if err != nil {
		c.Status(500).Send(err) //Return 500 Internal Server Error in order to !(error)!
		return
	}

	response, _ := json.Marshal(results) //Convert response into JSON
	c.Send(response) //Return JSON formatted response
}
/*<<END>> HTTP DELETE: Delete person <<END>> */
