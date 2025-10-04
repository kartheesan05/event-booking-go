package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents )

	err := server.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server")
	}
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message":"hello"})
}