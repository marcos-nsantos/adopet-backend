package adoptionhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/entity"
	"github.com/marcos-nsantos/adopet-backend/internal/schemas"
)

// CreateAdoption handles request to create a adoption
//
//	@Summary	Create an adoption
//	@Tags		adoptions
//	@Produce	json
//	@Param		tutorId	path		uint64	true	"Tutor id"
//	@Param		petId	path		uint64	true	"Pet id"
//	@Success	201		{object}	schemas.AdoptionResponse
//	@Failure	400
//	@Failure	422
//	@Router		/adoption/{tutorId}/{petId} [post]
func CreateAdoption(c *gin.Context) {
	tutorIDParam := c.Param("tutorId")
	if tutorIDParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tutor id is required"})
		return
	}

	tutorId, err := strconv.ParseUint(tutorIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tutor id"})
		return
	}

	petIDParam := c.Param("petId")
	if petIDParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "pet id is required"})
		return
	}

	petId, err := strconv.ParseUint(petIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid pet id"})
		return
	}

	adoption := entity.Adoption{TutorID: tutorId, PetID: petId}

	if err = database.Adopt(&adoption); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while adopting pet"})
		return
	}

	c.JSON(http.StatusCreated, schemas.ToAdoptionResponse(adoption))
}
