package tutorhandler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
)

func GetTutorByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	tutor, err := database.GetTutorByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("tutor with id %d not found", id)})
		return
	}

	c.JSON(http.StatusOK, toTutorResponse(tutor))
}
