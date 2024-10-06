package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/case-framework/case-backend/pkg/study/types"

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
	router := gin.Default()
	router.GET("/participants", getParticipants)
	router.GET("/participants/:id", getParticipantByID)
	router.POST("/participants", postParticipants)
	router.Run("localhost:8080")
}

// responds list of all participants as JSON.
func getParticipants(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, participants)
}

// adds an album from JSON received in the request body.
func postParticipants(c *gin.Context) {
	var newParticipant types.Participant

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newParticipant); err != nil {
		return
	}

	// Add the new album to the slice.
	participants = append(participants, newParticipant)
	c.IndentedJSON(http.StatusCreated, newParticipant)
}

// locates the participant whose ParticipantID value matches the id
func getParticipantByID(c *gin.Context) {
	participantID := c.Param("id")

	for _, a := range participants {
		if a.ParticipantID == participantID {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Participant could not be found"})
}
