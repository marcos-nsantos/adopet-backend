package shelterhandler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/entity"
	"github.com/marcos-nsantos/adopet-backend/internal/schemas"
	"github.com/marcos-nsantos/adopet-backend/pkg/password"
)

// CreateShelter handles request to create a shelter
//
//	@Summary	Create a shelter
//	@Tags		shelter
//	@Accept		json
//	@Produce	json
//	@Param		shelter	body		schemas.UserCreateRequest	true	"Shelter data"
//	@Success	201		{object}	schemas.UserResponse
//	@Failure	400
//	@Failure	409
//	@Failure	422
//	@Router		/shelters [post]
func CreateShelter(c *gin.Context) {
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
	user.Type = entity.Shelter

	shelter, err := database.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "email is already in use"})
		return
	}

	c.JSON(http.StatusCreated, schemas.ToUserResponse(shelter))
}
