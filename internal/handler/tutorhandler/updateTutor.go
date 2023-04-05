package tutorhandler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/marcos-nsantos/adopet-backend/internal/schemas"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
)

// UpdateTutor handles request to update a tutor
//
//	@Summary	Update a tutor
//	@Tags		tutor
//	@Accept		json
//	@Produce	json
//	@Param		id		path		uint						true	"Tutor id"
//	@Param		user	body		schemas.UserUpdateRequest	true	"Tutor data"
//	@Success	200		{object}	schemas.UserResponse
//	@Failure	400
//	@Failure	404
//	@Failure	422
//	@Router		/tutors/{id} [put]
func UpdateTutor(c *gin.Context) {
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

	if _, err = database.GetTutorByID(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tutor not found"})
		return
	}

	tutor := req.ToEntity()
	tutor.ID = id

	if err = database.UpdateUser(&tutor); err != nil {
		log.Println("error updating tutor", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error updating tutor"})
		return
	}

	c.JSON(http.StatusOK, schemas.ToUserResponse(tutor))
}
