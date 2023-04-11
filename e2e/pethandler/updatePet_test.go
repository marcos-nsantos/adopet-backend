package pethandler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/auth"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/entity"
	"github.com/marcos-nsantos/adopet-backend/internal/mock"
	"github.com/marcos-nsantos/adopet-backend/internal/router"
	"github.com/marcos-nsantos/adopet-backend/internal/schemas"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdatePet(t *testing.T) {
	database.Init()
	database.Migrate()
	gin.SetMode(gin.TestMode)
	r := router.SetupRoutes()

	t.Cleanup(func() {
		database.DropTables()
	})

	shelter := mock.Shelters()[0]
	shelter, err := database.CreateShelter(shelter)
	require.NoError(t, err)

	pet := mock.Pet()[0]
	pet.ShelterID = shelter.ID
	pet, err = database.CreatePet(pet)
	require.NoError(t, err)

	shelterToken, err := auth.GenerateToken(shelter.ID, entity.ShelterType)
	require.NoError(t, err)

	tests := []struct {
		name       string
		id         uint64
		reqBody    schemas.PetUpdateRequests
		token      string
		wantStatus int
	}{
		{
			name: "should return status 200",
			id:   pet.ID,
			reqBody: schemas.PetUpdateRequests{
				Name:        "Pet One Updated",
				Description: "Description of pet one updated",
				Age:         1,
				Photo:       "https://somephoto.com/photo.png",
				UF:          "SP",
				City:        "São Paulo",
			},
			token:      shelterToken,
			wantStatus: http.StatusOK,
		},
		{
			name: "should return status 422 when name is empty",
			id:   pet.ID,
			reqBody: schemas.PetUpdateRequests{
				Name:        "",
				Description: "Description of pet one updated",
				Age:         1,
				Photo:       "https://somephoto.com/photo.png",
				UF:          "SP",
				City:        "São Paulo",
			},
			token:      shelterToken,
			wantStatus: http.StatusUnprocessableEntity,
		},
		{
			name: "should return status 404 when pet is not found",
			id:   999,
			reqBody: schemas.PetUpdateRequests{
				Name:        "Pet One Updated",
				Description: "Description of pet one updated",
				Age:         1,
				Photo:       "https://somephoto.com/photo.png",
				UF:          "SP",
				City:        "São Paulo",
			},
			token:      shelterToken,
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "should return status 401 when token is not provided",
			id:         pet.ID,
			wantStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody, err := json.Marshal(tt.reqBody)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/pets/%d", tt.id), bytes.NewReader(reqBody))
			req.Header.Set("Authorization", "Bearer "+tt.token)
			require.NoError(t, err)

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}
