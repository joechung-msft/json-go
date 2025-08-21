package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joechung-msft/json-go/internal/shared"
)

func main() {
	app := fiber.New()

	app.Post("/api/v1/parse", func(c *fiber.Ctx) error {
		bodyBytes := c.Body()
		if len(bodyBytes) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Failed to read request body",
				"code":  400,
			})
		}

		jsonString := string(bodyBytes)
		var parsed any
		defer func() {
			if r := recover(); r != nil {
				_ = c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Invalid JSON",
					"code":  400,
				})
				parsed = nil
			}
		}()
		parsed = shared.Parse(jsonString)
		if parsed == nil {
			return nil // response already sent in recover
		}
		return c.JSON(parsed)
	})

	if err := app.Listen(":8080"); err != nil {
		panic(fmt.Sprintf("Unable to start server: %v", err))
	}
}
