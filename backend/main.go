package main

import (
	"fmt"
	"home"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// CORS middleware: reflect Origin for local development and log requests.
	// This allows the frontend dev server (Vite) to fetch the backend during development.
	router.Use(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin == "" {
			// no origin (e.g. server-side or curl) — allow all
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		} else {
			// reflect the provided origin (required if credentials are used)
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			// recommended to vary on Origin
			c.Writer.Header().Set("Vary", "Origin")
		}

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")

		// log for debugging
		fmt.Printf("[CORS] %s %s from Origin=%s\n", c.Request.Method, c.Request.URL.Path, origin)

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	router.GET("/", handleHome)

	router.Run(":8080")
}

func handleHome(c *gin.Context) {
	// write JSON response (no return value from c.JSON)
	print("Get message from home")
	c.JSON(http.StatusOK, gin.H{
		"message": home.Home(),
	})
}
