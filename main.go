package main

import (
	"event-booking-go/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	err := server.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server")
	}
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could Not Parse Request Data."})
		return
	}

	event.Id = 1
	event.UserID = 1

	context.JSON(http.StatusCreated, gin.H{"message": "Event Successfully Created.", "event": event})
}
