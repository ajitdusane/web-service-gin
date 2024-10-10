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

	ReadParticipants()

	if len(Participants) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Participants not defined"})
		return
	}

	c.JSON(http.StatusOK, Participants)
}

// adds a participant from JSON received in the request body and puts it inside a directory (see path in config.yml) as a .json file
func postParticipants(c *gin.Context) {
	var newParticipant types.Participant

	if err := c.BindJSON(&newParticipant); err != nil {
		slog.Error("error parsing json received in the request body.")
		c.JSON(http.StatusAlreadyReported, gin.H{"message": "unexpected error occured"})
		return
	}

	if ParticipantExists(newParticipant.ParticipantID + ".json") {
		c.JSON(http.StatusAlreadyReported, gin.H{"message": "Participant already exists."})
		return
	}

	WriteParticipantToFile(newParticipant)
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

	var participant types.Participant
	if json.Unmarshal(ReadParticipantAsJson(participantFileName), &participant) != nil {
		slog.Error("error while converting json to Participant struct")
		c.JSON(http.StatusNotFound, gin.H{"message": "Participant could not be loaded"})
		return
	}
	c.JSON(http.StatusOK, participant)
}
