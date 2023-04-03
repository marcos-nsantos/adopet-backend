package tutorhandler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
)

func UpdateTutor(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req TutorUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid request"})
		return
	}

	if _, err = database.GetTutorByID(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tutor not found"})
		return
	}

	tutor := req.ToEntity()
	tutor.ID = id

	if err = database.UpdateUser(&tutor); err != nil {
		log.Println("error updating tutor", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error updating tutor"})
		return
	}

	c.JSON(http.StatusOK, toUserResponse(tutor))
}
