package pethandler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/schemas"
)

func UpdatePet(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid id"})
		return
	}

	var req schemas.PetUpdateRequests
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if _, err = database.GetPetByID(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "pet not found"})
		return
	}

	pet := req.ToEntity()
	pet.ID = id

	if err = database.UpdatePet(pet); err != nil {
		log.Println("error to update pet", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error to update pet"})
		return
	}

	c.JSON(http.StatusOK, schemas.ToPetResponse(pet))
}
