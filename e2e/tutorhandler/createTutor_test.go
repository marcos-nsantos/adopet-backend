package tutorhandler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/marcos-nsantos/adopet-backend/internal/schemas"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/mock"
	"github.com/marcos-nsantos/adopet-backend/internal/router"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateTutor(t *testing.T) {
	database.InitTest()
	database.Migrate()
	gin.SetMode(gin.TestMode)
	r := router.SetupRoutes()

	t.Cleanup(func() {
		database.DropTables()
	})

	tutor := mock.Tutors()[0]

	tests := []struct {
		name       string
		reqBody    schemas.TutorCreationRequest
		wantStatus int
	}{
		{
			name: "should create a tutor and return status 201",
			reqBody: schemas.TutorCreationRequest{
				Name:     tutor.Name,
				Email:    tutor.Email,
				Password: tutor.Password,
				Phone:    tutor.Phone,
				Photo:    tutor.Photo,
				City:     tutor.City,
				About:    tutor.About,
			},
			wantStatus: http.StatusCreated,
		},
		{
			name: "should return status 422 when name is empty",
			reqBody: schemas.TutorCreationRequest{
				Name:     "",
				Email:    tutor.Email,
				Password: tutor.Password,
				Phone:    tutor.Phone,
				Photo:    tutor.Photo,
				City:     tutor.City,
				About:    tutor.About,
			},
			wantStatus: http.StatusUnprocessableEntity,
		},
		{
			name: "should return status 409 when email already exists",
			reqBody: schemas.TutorCreationRequest{
				Name:     tutor.Name,
				Email:    tutor.Email,
				Password: tutor.Password,
				Phone:    tutor.Phone,
				Photo:    tutor.Photo,
				City:     tutor.City,
				About:    tutor.About,
			},
			wantStatus: http.StatusConflict,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBody, err := json.Marshal(tt.reqBody)
			require.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, "/tutors", bytes.NewBuffer(jsonBody))
			require.NoError(t, err)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}
