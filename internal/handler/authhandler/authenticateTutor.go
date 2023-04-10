package authhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/auth"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/entity"
	"github.com/marcos-nsantos/adopet-backend/internal/password"
	"github.com/marcos-nsantos/adopet-backend/internal/schemas"
)

// AuthenticateTutor is a handler that authenticates a tutor
//
//	@Summary		Authenticate a tutor
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		schemas.AuthRequest	true	"Auth request"
//	@Success		200		{object}	schemas.AuthResponse
//	@Failure		400
//	@Failure		401
//	@Failure		500
//	@Router			/auth/tutor [post]
func AuthenticateTutor(c *gin.Context) {
	var req schemas.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, passwordGot, err := database.GetIDAndPasswordByEmailFromTutor(req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "email or password is incorrect"})
		return
	}

	if err = password.Compare(passwordGot, req.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "email or password is incorrect"})
		return
	}

	token, err := auth.GenerateToken(id, entity.TutorType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error generating token"})
		return
	}

	c.JSON(http.StatusOK, schemas.AuthResponse{Token: token})
}
