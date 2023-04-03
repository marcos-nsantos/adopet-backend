package shelterhandler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/mock"
	"github.com/marcos-nsantos/adopet-backend/internal/router"
	"github.com/marcos-nsantos/adopet-backend/internal/schemas"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateShelter(t *testing.T) {
	database.Init()
	database.Migrate()
	gin.SetMode(gin.TestMode)
	r := router.SetupRoutes()

	t.Cleanup(func() {
		database.DropTables()
	})

	shelter := mock.Shelters()[0]
	shelter.Type = ""

	tests := []struct {
		name       string
		reqBody    schemas.UserCreateRequest
		wantStatus int
	}{
		{
			name: "should create a shelter and return status 201",
			reqBody: schemas.UserCreateRequest{
				Name:     shelter.Name,
				Email:    shelter.Email,
				Password: shelter.Password,
				Phone:    shelter.Phone,
				Photo:    shelter.Photo,
				City:     shelter.City,
				About:    shelter.About,
			},
			wantStatus: http.StatusCreated,
		},
		{
			name: "should return status 422 when name is empty",
			reqBody: schemas.UserCreateRequest{
				Name:     "",
				Email:    shelter.Email,
				Password: shelter.Password,
				Phone:    shelter.Phone,
				Photo:    shelter.Photo,
				City:     shelter.City,
				About:    shelter.About,
			},
			wantStatus: http.StatusUnprocessableEntity,
		},
		{
			name: "should return status 409 when email already exists",
			reqBody: schemas.UserCreateRequest{
				Name:     shelter.Name,
				Email:    shelter.Email,
				Password: shelter.Password,
				Phone:    shelter.Phone,
				Photo:    shelter.Photo,
				City:     shelter.City,
				About:    shelter.About,
			},
			wantStatus: http.StatusConflict,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBody, err := json.Marshal(tt.reqBody)
			require.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, "/shelters", bytes.NewBuffer(jsonBody))
			require.NoError(t, err)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}
