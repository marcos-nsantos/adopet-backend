package shelterhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/entity"
)

// DeleteShelter handles request to delete a shelter
//
//	@Summary	Delete a shelter
//	@Tags		shelters
//	@Security	Bearer
//	@Param		id	path	uint	true	"Shelter id"
//	@Success	204
//	@Failure	400
//	@Failure	404
//	@Router		/shelters/{id} [delete]
func DeleteShelter(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err = database.DeleteShelter(id); err != nil {
		if err == entity.ErrShelterNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "shelter not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting shelter"})
		return
	}

	c.Status(http.StatusNoContent)
}
