package tutorhandler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
)

// DeleteTutor handles request to delete a tutor by id
//
//	@Summary	Delete a tutor by id
//	@Tags		tutor
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

	if _, err = database.GetTutorByID(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tutor not found"})
		return
	}

	if err = database.DeleteTutor(id); err != nil {
		log.Println("error deleting tutor", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting tutor"})
		return
	}

	c.Status(http.StatusNoContent)
}
