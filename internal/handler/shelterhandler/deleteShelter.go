package shelterhandler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
)

func DeleteShelter(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if _, err = database.GetShelterByID(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shelter not found"})
		return
	}

	if err = database.DeleteUser(id); err != nil {
		log.Println("error deleting shelter", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting shelter"})
		return
	}

	c.Status(http.StatusNoContent)
}
