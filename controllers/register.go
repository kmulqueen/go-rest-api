package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kmulqueen/go-rest-api/models"
)

func RegisterForEvent(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve event."})
		return
	}

	if event.UserID == userID {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cannot register for your own event."})
		return
	}

	err = event.Register(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Successfully registered."})
}

func CancelRegisterForEvent(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	var event models.Event
	event.ID = eventID

	err = event.CancelRegistration(userID)
	if err != nil {
		if err.Error() == "Registration not found." {
			context.JSON(http.StatusNotFound, gin.H{"message": "Registration not found."})
			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel user's registration for the event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully canceled registration."})
}
