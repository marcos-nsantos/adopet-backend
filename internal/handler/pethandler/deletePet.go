package pethandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/entity"
)

// DeletePet handles request to delete a pet
//
//	@Summary	Delete a pet
//	@Tags		pets
//	@Security	Bearer
//	@Param		id	path	uint	true	"Pet id"
//	@Success	204
//	@Failure	400
//	@Failure	404
//	@Router		/pets/{id} [delete]
func DeletePet(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err = database.DeletePet(id); err != nil {
		if err == entity.ErrPetNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "pet not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting a pet"})
		return
	}

	c.Status(http.StatusNoContent)
}
