package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var todos = make([]string, 0)

func Lists(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"list": todos})
}

func main() {
	todos = append(todos, "Write the application")
	r := gin.Default()
	r.GET("/api/lists", Lists)
	r.Run()
}
