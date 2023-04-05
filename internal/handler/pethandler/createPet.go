package pethandler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/schemas"
)

// CreatePet handles request to create a pet
//
//	@Summary	Create a pet
//	@Tags		pet
//	@Accept		json
//	@Produce	json
//	@Param		pet	body		schemas.PetCreateRequests	true	"Pet data"
//	@Success	201	{object}	schemas.PetResponse
//	@Failure	400
//	@Failure	422
//	@Router		/pets [post]
func CreatePet(c *gin.Context) {
	var req schemas.PetCreateRequests
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	pet, err := database.CreatePet(req.ToEntity())
	if err != nil {
		log.Println("error while creating pet", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while creating pet"})
		return
	}

	c.JSON(http.StatusCreated, schemas.ToPetResponse(pet))
}
