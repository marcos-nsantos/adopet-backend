package tutorhandler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
)

func CreateTutor(c *gin.Context) {
	var req TutorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	tutor := req.ToEntity()
	if err := database.CreateTutor(tutor); err != nil {
		errMessage := fmt.Errorf("email %s is already in use", tutor.Email)
		c.JSON(http.StatusConflict, gin.H{"error": errMessage.Error()})
		return
	}

	c.JSON(http.StatusCreated, toTutorResponse(tutor))
}
