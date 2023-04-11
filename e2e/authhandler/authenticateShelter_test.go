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

func TestAuthenticateShelter(t *testing.T) {
	database.Init()
	database.Migrate()
	gin.SetMode(gin.TestMode)
	r := router.SetupRoutes()

	t.Cleanup(func() {
		database.DropTables()
	})

	shelter := mock.Shelters()[0]
	passwordHashed, err := password.Hash(shelter.Password)
	require.NoError(t, err)
	shelter.Password = passwordHashed

	shelterCreated, err := database.CreateShelter(shelter)
	require.NoError(t, err)

	tests := []struct {
		name       string
		reqBody    schemas.AuthRequest
		wantStatus int
	}{
		{
			name: "should authenticate a shelter and return status 200",
			reqBody: schemas.AuthRequest{
				Email:    shelterCreated.Email,
				Password: mock.Shelters()[0].Password,
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "should return status 401 when email is wrong",
			reqBody: schemas.AuthRequest{
				Email:    "wrongemail@email.com",
				Password: mock.Shelters()[0].Password,
			},
			wantStatus: http.StatusUnauthorized,
		},
		{
			name: "should return status 401 when password is wrong",
			reqBody: schemas.AuthRequest{
				Email:    shelterCreated.Email,
				Password: "wrongpassword",
			},
			wantStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBody, err := json.Marshal(tt.reqBody)
			require.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, "/auth/shelter", bytes.NewBuffer(jsonBody))
			require.NoError(t, err)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			if tt.wantStatus == http.StatusOK {
				var resp schemas.AuthResponse
				err = json.Unmarshal(w.Body.Bytes(), &resp)
				require.NoError(t, err)

				assert.NotEmpty(t, resp.Token)
			}
		})
	}
}
