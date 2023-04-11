package tutorhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/entity"
)

// DeleteTutor handles request to delete a tutor by id
//
//	@Summary	Delete a tutor by id
//	@Tags		tutors
//	@Security	Bearer
//	@Param		id	path	uint	true	"Tutor id"
//	@Success	204
//	@Failure	400
//	@Failure	404
//	@Router		/tutors/{id} [delete]
func DeleteTutor(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err = database.DeleteTutor(id); err != nil {
		if err == entity.ErrTutorNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "tutor not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting tutor"})
		return
	}

	c.Status(http.StatusNoContent)
}
