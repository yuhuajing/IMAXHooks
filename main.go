package main

import (
	"fmt"
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Post("/", func(c *fiber.Ctx) error {
		fmt.Println(string(c.Body()))
		return c.SendString("POST request")
	})

	log.Fatal(app.Listen(":3010"))
}
