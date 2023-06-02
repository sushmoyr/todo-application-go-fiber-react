package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func main() {
	fmt.Print("Hello World")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:4001",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	var todos []Todo

	app.Get("/health-check", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Ok")
	})

	app.Post("/api/todos", func(ctx *fiber.Ctx) error {
		todo := &Todo{}

		err := ctx.BodyParser(todo)

		if err != nil {
			return err
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return ctx.JSON(todos)
	})

	app.Patch("/api/todos/:id/done", func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")

		if err != nil {
			return ctx.Status(401).SendString("Invalid Id")
		}

		for i, todo := range todos {
			if todo.ID == id {
				todos[i].Done = true
				break
			}
		}

		return ctx.JSON(todos)
	})

	app.Get("/api/todos", func(ctx *fiber.Ctx) error {
		return ctx.JSON(todos)
	})

	app.Get("/api/todos/:id", func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")

		if err != nil {
			return ctx.Status(404).JSON(ErrorResponse{
				Status:  404,
				Message: "Todo not found",
			})
		}

		for _, todo := range todos {
			if todo.ID == id {
				return ctx.JSON(todo)
			}
		}

		return ctx.Status(404).JSON(ErrorResponse{
			Status:  404,
			Message: "Todo not found",
		})
	})

	log.Fatal(app.Listen(":4000"))
}
