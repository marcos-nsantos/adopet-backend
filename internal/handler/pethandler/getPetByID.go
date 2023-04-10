package pethandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/schemas"
)

// GetPetByID handles request to get a pet by id
//
//	@Summary	Get a pet by id
//	@Tags		pets
//	@Security	Bearer
//	@Param		id	path		uint	true	"Pet id"
//	@Success	200	{object}	schemas.PetResponse
//	@Failure	404
//	@Failure	422
//	@Router		/pets/{id} [get]
func GetPetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid id"})
		return
	}

	pet, err := database.GetPetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "pet not found"})
		return
	}

	c.JSON(http.StatusOK, schemas.ToPetResponse(pet))
}
