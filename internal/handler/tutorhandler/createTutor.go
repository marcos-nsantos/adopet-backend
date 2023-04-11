package tutorhandler

import (
	"log"
	"net/http"

	"github.com/marcos-nsantos/adopet-backend/internal/password"
	"github.com/marcos-nsantos/adopet-backend/internal/schemas"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
)

// CreateTutor handles request to create a new tutor
//
//	@Summary	Create a new tutor
//	@Tags		tutors
//	@Accept		json
//	@Produce	json
//	@Param		tutor	body		schemas.TutorCreationRequest	true	"User data"
//	@Success	201		{object}	schemas.TutorResponse
//	@Failure	400
//	@Failure	409
//	@Router		/tutors [post]
func CreateTutor(c *gin.Context) {
	var req schemas.TutorCreationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	user := req.ToEntity()
	passwordHashed, err := password.Hash(user.Password)
	if err != nil {
		log.Println("error while hashing password", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while hashing password"})
		return
	}

	user.Password = passwordHashed
	tutor, err := database.CreateTutor(user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "email is already in use"})
		return
	}

	c.JSON(http.StatusCreated, schemas.ToTutorResponse(tutor))
}
