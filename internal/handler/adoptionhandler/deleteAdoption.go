package adoptionhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/entity"
)

// DeleteAdoption handles request to delete a adoption
//
//	@Summary	Delete an adoption
//	@Tags		adoptions
//	@Produce	json
//	@Param		id	path	uint64	true	"Adoption id"
//	@Success	204
//	@Failure	400
//	@Failure	422
//	@Router		/adoption/{petId} [delete]
func DeleteAdoption(c *gin.Context) {
	userType, ok := c.Get("user_type")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while getting user type"})
		return
	}

	if userType != entity.ShelterType {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "only shelters can delete adoptions"})
		return
	}

	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	adoption := entity.Adoption{ID: id}
	if err = database.DeleteAdoption(&adoption); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while deleting adoption"})
		return
	}

	c.Status(http.StatusNoContent)
}
