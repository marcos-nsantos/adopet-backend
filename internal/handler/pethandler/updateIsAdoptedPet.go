package pethandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/entity"
	"github.com/marcos-nsantos/adopet-backend/internal/schemas"
)

// UpdateIsAdoptedPet handles request to update a pet's adoption status
//
//	@Summary	update a pet's adoption status
//	@Tags		pets
//	@Security	Bearer
//	@Accept		json
//	@Produce	json
//	@Param		id	path	uint								true	"Pet id"
//	@Param		pet	body	schemas.UpdateIsAdoptPetRequests	true	"Pet data"
//	@Success	204
//	@Failure	400
//	@Failure	404
//	@Failure	422
//	@Router		/pets/{id}/adopted [patch]
func UpdateIsAdoptedPet(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid id"})
		return
	}

	var req schemas.UpdateIsAdoptPetRequests
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	pet := entity.Pet{ID: id, IsAdopted: req.IsAdopted}
	if err = database.UpdateIsAdoptedPet(pet); err != nil {
		if err == entity.ErrPetNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "pet not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "error to update pet"})
		return
	}

	c.Status(http.StatusNoContent)
}
