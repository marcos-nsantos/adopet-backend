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
//	@Param		shelter	body		schemas.TutorUpdateRequest	true	"Shelter data"
//	@Success	200		{object}	schemas.TutorResponse
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

	var req schemas.ShelterUpdateRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if _, err = database.GetShelterByID(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shelter not found"})
		return
	}

	shelter := req.ToEntity()
	shelter.ID = id

	if err = database.UpdateShelter(&shelter); err != nil {
		log.Println("error updating shelter", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error updating shelter"})
		return
	}

	c.JSON(http.StatusOK, schemas.ToShelterResponse(shelter))
}
