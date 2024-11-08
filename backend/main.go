package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}

var collection *mongo.Collection

func main() {
	fmt.Println("hello world!")

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error while loading .env file")
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal("error while connected to mongodb")
	}

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to mongodb")

	collection = client.Database("golang_db").Collection("todos")

	app := fiber.New()

	app.Get("/api/todos", getTodos)
	app.Post("/api/todos", createTodo)
	app.Patch("/api/todos/:id", updateTodo)
	app.Delete("/api/todos/:id", deleteTodo)

	PORT := os.Getenv("PORT")

	log.Fatal(app.Listen("0.0.0.0:" + PORT))
}

func getTodos(c *fiber.Ctx) error {
	var todos []Todo

	cursor, err := collection.Find(context.Background(), bson.M{})
	
	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			return err
		}

		todos = append(todos, todo)
	}

	return c.JSON(todos)
}

func createTodo(c *fiber.Ctx) error {
	todo := new(Todo)

	if err := c.BodyParser(todo); err != nil {
		return err
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{ "error" : "Todo body can't be empty"})
	}

	insertResult, err := collection.InsertOne(context.Background(), todo)

	if err != nil {
		return err
	}

	todo.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(todo)
}

func updateTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{ "error" : "Invalid todo ID"})
	}

	filter := bson.M{"_id" : objectID}
	update := bson.M{"$set" : bson.M{"completed": true}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"success": true})
}

func deleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error" : "Invalid todo ID"})
	}

	filter := bson.M{"_id" : objectID}
	_, err = collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"success": true})
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