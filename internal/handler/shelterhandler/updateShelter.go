package shelterhandler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/schemas"
)

// UpdateShelter handles request to update a shelter
//
//	@Summary	Update a shelter
//	@Tags		shelter
//	@Accept		json
//	@Produce	json
//	@Param		id		path		uint						true	"Shelter id"
//	@Param		shelter	body		schemas.UserUpdateRequest	true	"Shelter data"
//	@Success	200		{object}	schemas.UserResponse
//	@Failure	400
//	@Failure	404
//	@Router		/shelters/{id} [put]
func UpdateShelter(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req schemas.UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid request"})
		return
	}

	if _, err = database.GetShelterByID(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shelter not found"})
		return
	}

	shelter := req.ToEntity()
	shelter.ID = id

	if err = database.UpdateUser(&shelter); err != nil {
		log.Println("error updating shelter", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error updating shelter"})
		return
	}

	c.JSON(http.StatusOK, schemas.ToUserResponse(shelter))
}
