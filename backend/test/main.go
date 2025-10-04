package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Allow simple CORS for local development
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	router.GET("/", handleHome)
	router.GET("/other", handleOther)

	// listen on all interfaces on port 8080
	router.Run(":8080")
}

func handleHome(c *gin.Context) {
	// return a JSON object with a message field
	c.JSON(http.StatusOK, gin.H{"message": "welcome to my page"})
}

func handleOther(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "This is another message"})
}
