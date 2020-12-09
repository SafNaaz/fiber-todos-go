package main

import "github.com/gofiber/fiber/v2"

//Todo ...
type Todo struct {
	ID        int
	Name      string
	Completed bool
}

var todos = []Todo{
	{ID: 1, Name: "Bla bla bla", Completed: false},
	{ID: 2, Name: "Blu blau blu", Completed: false},
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
