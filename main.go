package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct{
	ID int `json:"id"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}

func main(){
	fmt.Println("Hello bitchess")

	app:= fiber.New()
	err := godotenv.Load(".env")
	if err!=nil{
		log.Fatal("Error loading .env file")
	}
	
	// import the godotenv 
	PORT := os.Getenv("PORT")

	todos := [] Todo{}


	app.Get("/", func(c *fiber.Ctx) error{
		return c.Status(200).JSON(fiber.Map{"msg":"Home page is being returned nig"})
	})


	// creating a post request for adding a todo

	app.Post("/api/todos",func(c *fiber.Ctx) error{
		todo := &Todo{}

		if err:=c.BodyParser(todo); err!=nil{
			return err
		}
		if todo.Body == ""{
			return c.Status(400).JSON(fiber.Map{"msg":"Give Todo Desciption"})
		}

		todo.ID = len(todos)+1
		todos = append(todos,*todo)

		return c.Status(201).JSON(todo)


	})

	// mark a todo as completed

	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error{
		id := c.Params("id");

		for i,todo:= range todos{
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true;
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"error":"Todo not found"})

	})

	// try getting all todos 
	app.Get("/api/todos", func(c *fiber.Ctx) error{

		return c.Status(200).JSON(todos)

	})

	// delete a todo 

	app.Delete("/api/todos/:id",func(c *fiber.Ctx) error{

		id:= c.Params("id")

		for i,todo := range todos{
			if fmt.Sprint(todo.ID) == id{
				todos = append(todos[:i],todos[i+1:]...)
				return c.Status(201).JSON(fiber.Map{"msg":"Todo is deleted"}) 
			}
		}

		return c.Status(400).JSON(fiber.Map{"error":"Todo not found"})
	})


	log.Fatal(app.Listen(":"+PORT))
	
}