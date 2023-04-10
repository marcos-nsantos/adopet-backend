package tutorhandler

import (
	"net/http"
	"strconv"

	"github.com/marcos-nsantos/adopet-backend/internal/schemas"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
)

// GetTutorByID handles request to get a tutor by id
//
//	@Summary	Get a tutor by id
//	@Tags		tutors
//	@Security	Bearer
//	@Produce	json
//	@Param		id	path		uint	true	"Tutor id"
//	@Success	200	{object}	schemas.TutorResponse
//	@Failure	400
//	@Failure	404
//	@Router		/tutors/{id} [get]
func GetTutorByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	tutor, err := database.GetTutorByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tutor not found"})
		return
	}

	c.JSON(http.StatusOK, schemas.ToTutorResponse(tutor))
}
