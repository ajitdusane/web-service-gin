package main

import (
	"net/http"

	"github.com/ajitdusane/web-service-gin/app_config"

	"github.com/case-framework/case-backend/pkg/study/types"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var participants = []types.Participant{
	{
		ID:                  primitive.NewObjectID(),
		ParticipantID:       "P1",
		CurrentStudySession: "CSS1",
		EnteredAt:           11,
		StudyStatus:         "SS11",
		Flags:               make(map[string]string),
		AssignedSurveys:     []types.AssignedSurvey{},
		LastSubmissions:     make(map[string]int64),
		Messages:            []types.ParticipantMessage{},
	},
}

func main() {
	app_config.ReadConfig()
	router := gin.Default()
	router.GET("/participants", getParticipants)
	router.GET("/participants/:id", getParticipantByID)
	router.POST("/participants", postParticipants)
	router.Run(app_config.AppConfig.Server.Host + ":" + app_config.AppConfig.Server.Port)
}

// responds list of all participants as JSON.
func getParticipants(c *gin.Context) {

	c.JSON(http.StatusOK, participants)
}

// adds a participant from JSON received in the request body.
func postParticipants(c *gin.Context) {
	var newParticipant types.Participant

	if err := c.BindJSON(&newParticipant); err != nil {
		return
	}

	participants = append(participants, newParticipant)
	c.JSON(http.StatusCreated, newParticipant)
}

// locates the participant whose ParticipantID value matches the id
func getParticipantByID(c *gin.Context) {
	participantID := c.Param("id")

	for _, a := range participants {
		if a.ParticipantID == participantID {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Participant could not be found"})
}
