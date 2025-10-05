package main

import (
	"home"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Enable CORS for localhost:5173
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
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
	c.JSON(http.StatusOK, gin.H{
		"message": home.Home(),
	})
}
