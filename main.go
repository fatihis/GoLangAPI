package main

import (
	"github.com/gofiber/fiber"
)

//types
type Person struct {
	_id       string `json:”id,omitempty”`
	FirstName string `json:”firstname,omitempty”`
	LastName  string `json:”lastname,omitempty”`
	Email     string `json:”email,omitempty”`
	Age       int    `json:”age,omitempty”`
}

func main() {
	//Construct REST API with fiber infrastructure
	app := fiber.New()

	app.Get("/get/:id", getPerson)
	app.Get("/getAll", getAllPersons)
	app.Post("/create", createPerson)
	app.Put("/update/:id", updatePerson)
	app.Delete("/delete/:id", deletePerson)
	port := 27027 //API will be served localhost:27027 || 0.0.0.0:27027
	app.Listen(port)
}
