package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/joechung-msft/json-go/internal/shared"

	"github.com/gin-gonic/gin"
)

// https://gin-gonic.com/
func main() {
	router := gin.Default()
	if err := router.SetTrustedProxies(nil); err != nil {
		fmt.Println("Failed to set trusted proxies:", err)
		panic(err)
	}

	router.POST("/api/v1/parse", func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body", "code": 400})
			return
		}
		jsonString := string(body)

		var result any
		defer func() {
			if r := recover(); r != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON", "code": 400})
				result = nil
			}
		}()
		result = shared.Parse(jsonString)
		if result == nil {
			return
		}

		c.JSON(http.StatusOK, result)
	})

	if err := router.Run(); err != nil {
		panic(err)
	}
}
