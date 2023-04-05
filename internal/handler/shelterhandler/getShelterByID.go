package shelterhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/schemas"
)

// GetShelterByID handles request to get a shelter by id
//
//	@Summary	Get a shelter by id
//	@Tags		shelter
//	@Produce	json
//	@Param		id	path		uint	true	"Shelter id"
//	@Success	200	{object}	schemas.UserResponse
//	@Failure	400
//	@Failure	404
//	@Router		/shelters/{id} [get]
func GetShelterByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	shelter, err := database.GetShelterByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shelter not found"})
		return
	}

	c.JSON(http.StatusOK, schemas.ToUserResponse(shelter))
}
