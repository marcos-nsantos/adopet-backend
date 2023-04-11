package authhandler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/mock"
	"github.com/marcos-nsantos/adopet-backend/internal/password"
	"github.com/marcos-nsantos/adopet-backend/internal/router"
	"github.com/marcos-nsantos/adopet-backend/internal/schemas"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthenticateTutor(t *testing.T) {
	database.Init()
	database.Migrate()
	gin.SetMode(gin.TestMode)
	r := router.SetupRoutes()

	t.Cleanup(func() {
		database.DropTables()
	})

	tutor := mock.Tutors()[0]
	passwordHashed, err := password.Hash(tutor.Password)
	require.NoError(t, err)
	tutor.Password = passwordHashed

	tutorCreated, err := database.CreateTutor(tutor)
	require.NoError(t, err)

	tests := []struct {
		name       string
		reqBody    schemas.AuthRequest
		wantStatus int
	}{
		{
			name: "should authenticate a tutor and return status 200",
			reqBody: schemas.AuthRequest{
				Email:    tutorCreated.Email,
				Password: mock.Tutors()[0].Password,
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "should return status 401 when email is wrong",
			reqBody: schemas.AuthRequest{
				Email:    "wrongemail@email.com",
				Password: mock.Tutors()[0].Password,
			},
			wantStatus: http.StatusUnauthorized,
		},
		{
			name: "should return status 401 when password is wrong",
			reqBody: schemas.AuthRequest{
				Email:    tutorCreated.Email,
				Password: "wrongpassword",
			},
			wantStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBody, err := json.Marshal(tt.reqBody)
			require.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, "/auth/tutor", bytes.NewBuffer(jsonBody))
			require.NoError(t, err)

			res := httptest.NewRecorder()
			r.ServeHTTP(res, req)

			assert.Equal(t, tt.wantStatus, res.Code)
			if tt.wantStatus == http.StatusOK {
				var response schemas.AuthResponse
				err = json.NewDecoder(res.Body).Decode(&response)
				require.NoError(t, err)

				require.NotEmpty(t, response.Token)
			}
		})
	}
}
