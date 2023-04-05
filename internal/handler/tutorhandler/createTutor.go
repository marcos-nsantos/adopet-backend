package tutorhandler

import (
	"log"
	"net/http"

	"github.com/marcos-nsantos/adopet-backend/internal/entity"
	"github.com/marcos-nsantos/adopet-backend/internal/password"
	"github.com/marcos-nsantos/adopet-backend/internal/schemas"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
)

// CreateTutor handles request to create a new tutor
//
//	@Summary	Create a new tutor
//	@Tags		tutor
//	@Accept		json
//	@Produce	json
//	@Param		tutor	body		schemas.UserCreateRequest	true	"User data"
//	@Success	201		{object}	schemas.UserResponse
//	@Failure	400
//	@Failure	409
//	@Router		/tutors [post]
func CreateTutor(c *gin.Context) {
	var req schemas.UserCreateRequest
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
	user.Type = entity.Tutor

	tutor, err := database.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "email is already in use"})
		return
	}

	c.JSON(http.StatusCreated, schemas.ToUserResponse(tutor))
}
