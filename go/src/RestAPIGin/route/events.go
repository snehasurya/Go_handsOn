package routes

import (
	"net/http"
	"strconv"

	"example.com/restApiGin/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	event, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get the event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": event})
}

func postEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}
	userID := context.GetInt64("userID")
	event.UserID = userID
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create the event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": event})
}

func getSingleEvent(context *gin.Context) {
	event := fetchEvent(context)
	if event == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create the event"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func putEvent(context *gin.Context) {
	userID := context.GetInt64("userID")
	event := fetchEvent(context)
	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User is not autorized to update the event"})
		return
	}
	var updatedEvent models.Event
	err := context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}
	updatedEvent.ID = event.ID
	err = updatedEvent.UpdateEventById()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update the event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event updated successfully"})
}

func deleteEvent(context *gin.Context) {
	userID := context.GetInt64("userID")
	event := fetchEvent(context)
	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User is not autorized to delete the event"})
		return
	}
	err := event.Delete()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not delete event ID"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event deleted successfully"})
}

func fetchEvent(context *gin.Context) *models.Event {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		return nil
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event"})
		return nil
	}
	return event
}
