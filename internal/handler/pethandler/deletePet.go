package pethandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
)

func DeletePet(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if _, err := database.GetPetByID(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "pet not found"})
		return
	}

	if err := database.DeletePet(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting a pet"})
		return
	}

	c.Status(http.StatusNoContent)
}
