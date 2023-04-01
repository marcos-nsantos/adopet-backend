package tutorhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
)

func CreateTutor(c *gin.Context) {
	var req TutorCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	tutor, err := database.CreateTutor(req.ToEntity())
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "email is already in use"})
		return
	}

	c.JSON(http.StatusCreated, toTutorResponse(tutor))
}
