package tutorhandler

import (
	"github.com/marcos-nsantos/adopet-backend/internal/entity"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/pkg/password"
)

func CreateTutor(c *gin.Context) {
	var req UserCreateRequest
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

	c.JSON(http.StatusCreated, toUserResponse(tutor))
}
