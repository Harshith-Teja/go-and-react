package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID int `json: "_id" bson: "_id"`
	Completed bool `json: "completed"`
	Body string `json: "body"`
}

var collection *mongo.Collection

func main() {
	fmt.Println("hello world!")
}







/*
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID int `json: "id"`
	Completed bool `json: "completed"`
	Body string `json: "body"`
}

func main() {
	fmt.Println("Hello World");
	app := fiber.New()

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("err loading the .env file")
	}

	PORT := os.Getenv("PORT")

	todos := []Todo{}

	//GET Todos
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	//CREATE a Todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{} //{id: 0, completed: false, body: ""}

		if err := c.BodyParser(todo); err != nil {
			return err;
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error" : "Todo body is required"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})

	//UPDATE a Todo
	app.Patch("/api/todos/:id", func (c *fiber.Ctx) error {
		id := c.Params("id")

		for ind, todo := range todos {

			if fmt.Sprint(todo.ID) == id {
				todos[ind].Completed = true
				return c.Status(200).JSON(todos[ind])
			}
		}

		return c.Status(404).JSON(fiber.Map{ "error" : "Todo not found"})
	})


	//DELETE a Todo
	app.Delete("/api/todos/:id", func (c *fiber.Ctx) error {

		id := c.Params("id")

		for ind, todo := range todos {

			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:ind], todos[ind+1:]...)
				return c.Status(200).JSON(fiber.Map{"success" : "true"})
			}
		}

		return c.Status(400).JSON(fiber.Map{"error": "Todo not found"})
	})

	log.Fatal(app.Listen(":" + PORT))	
}

*/