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

	tutor, err := database.CreateTutor(req.ToEntity())
	if err != nil {
		errMessage := fmt.Sprintf("email %s already exists", req.Email)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMessage})
		return
	}

	c.JSON(http.StatusCreated, toTutorResponse(tutor))
}
