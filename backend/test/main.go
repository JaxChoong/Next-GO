package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", handleHome)

	router.Run("localhost:8080")
}

func handleHome(c *gin.Context) {
	message := "welcome to my page"
	c.JSON(http.StatusOK, message)
}
