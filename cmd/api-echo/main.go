package main

import (
	"io"
	"net/http"

	"github.com/joechung-msft/json-go/internal/shared"

	"github.com/labstack/echo/v4"
)

// https://echo.labstack.com/
func main() {
	e := echo.New()

	e.POST("/api/v1/echo", func(c echo.Context) error {
		bodyBytes, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{"error": "Failed to read request body", "code": 400})
		}
		jsonString := string(bodyBytes)

		var parsed any
		defer func() {
			if r := recover(); r != nil {
				c.JSON(http.StatusBadRequest, map[string]any{"error": "Invalid JSON", "code": 400})
				parsed = nil
			}
		}()
		parsed = shared.Parse(jsonString)
		if parsed == nil {
			return nil // response already sent in recover
		}
		return c.JSON(http.StatusOK, parsed)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
