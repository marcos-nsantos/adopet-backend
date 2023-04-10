package shelterhandler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/password"
	"github.com/marcos-nsantos/adopet-backend/internal/schemas"
)

// CreateShelter handles request to create a shelter
//
//	@Summary	Create a shelter
//	@Tags		shelter
//	@Accept		json
//	@Produce	json
//	@Param		shelter	body		schemas.ShelterCreationRequest	true	"Shelter data"
//	@Success	201		{object}	schemas.ShelterResponse
//	@Failure	400
//	@Failure	409
//	@Failure	422
//	@Router		/shelters [post]
func CreateShelter(c *gin.Context) {
	var req schemas.ShelterCreationRequest
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
	shelter, err := database.CreateShelter(user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "email is already in use"})
		return
	}

	c.JSON(http.StatusCreated, schemas.ToShelterResponse(shelter))
}
