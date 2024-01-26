package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pedro-hca/go-studies/09-rest-api/db"
	"github.com/pedro-hca/go-studies/09-rest-api/models"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents)    // GET, POST, PUT, PATCH, DELETE
	server.GET("/events/:id", getEvent) // /events/1, /5, /6
	server.POST("/events", createEvent)
	server.Run(":8080") //localhost 8080
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse id."})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}
	context.JSON(http.StatusOK, event)
}
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	event.ID = 1
	event.UserID = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create events."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}