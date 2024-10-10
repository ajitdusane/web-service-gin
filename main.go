package main

import (
	"net/http"

	"log/slog"

	"github.com/ajitdusane/web-service-gin/app_config"

	"github.com/case-framework/case-backend/pkg/study/types"

	"github.com/gin-gonic/gin"

	"encoding/json"
)

var Participants = []types.Participant{}

func init() {
	app_config.ReadConfig()
}

func main() {

	router := gin.Default()
	router.GET("/participants", getParticipants)
	router.GET("/participants/:id", getParticipantByID)
	router.POST("/participants", postParticipants)
	slog.Info("starting server " + app_config.AppConfig.Server.Host + " at port " + app_config.AppConfig.Server.Port)
	router.Run(app_config.AppConfig.Server.Host + ":" + app_config.AppConfig.Server.Port)
}

// responds list of all participants as JSON. i.e. by reading all json files inside a directory (see path in config.yml)
func getParticipants(c *gin.Context) {

	if err := ReadAllParticipants(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error retrieving participants: " + err.Error()})
		return
	}

	if len(Participants) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No participants found."})
		return
	}

	c.JSON(http.StatusOK, Participants)
}

// adds a participant from JSON received in the request body and puts it inside a directory (see path in config.yml) as a .json file
func postParticipants(c *gin.Context) {
	var newParticipant types.Participant

	if err := c.BindJSON(&newParticipant); err != nil {
		slog.Error("error parsing json received in the request body.")
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid JSON format"})
		return
	}

	if ParticipantExists(newParticipant.ParticipantID + ".json") {
		c.JSON(http.StatusConflict, gin.H{"message": "Participant already exists."})
		return
	}

	if err := WriteParticipantToFile(newParticipant); err != nil {
		slog.Error("error writing participant to file", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to save participant."})

	}
	c.JSON(http.StatusCreated, newParticipant)
}

// locates the participant whose ParticipantID value matches the id
func getParticipantByID(c *gin.Context) {
	participantID := c.Param("id")
	participantFileName := participantID + ".json"

	if !ParticipantExists(participantFileName) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Participant could not be found"})
		return
	}

	content, err := ReadFileAsBytes(participantFileName)
	if err != nil {
		slog.Error("error reading Participant file", slog.String("file", participantFileName))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Participant could not be loaded"})
		return
	}

	var participant types.Participant
	if err := json.Unmarshal(content, &participant); err != nil {
		slog.Error("error while converting json to Participant struct")
		c.JSON(http.StatusNotFound, gin.H{"message": "Participant could not be loaded"})
		return
	}

	c.JSON(http.StatusOK, participant)
}
