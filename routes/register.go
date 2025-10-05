package routes

import (
	"event-booking-go/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Event does not exist."})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could Not Register."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered Successfully."})
}

func cancelRegistration(context *gin.Context) {

}
