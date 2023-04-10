package pethandler

import (
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

func TestGetPetByID(t *testing.T) {
	database.InitTest()
	database.Migrate()
	gin.SetMode(gin.TestMode)
	r := router.SetupRoutes()

	t.Cleanup(func() {
		database.DropTables()
	})

	shelter := mock.Shelters()[0]
	shelterCreated, err := database.CreateShelter(shelter)
	require.NoError(t, err)

	pet := mock.Pet()[0]
	pet.ShelterID = shelterCreated.ID
	pet, err = database.CreatePet(pet)
	require.NoError(t, err)

	shelterToken, err := auth.GenerateToken(shelterCreated.ID, entity.ShelterType)
	require.NoError(t, err)

	tests := []struct {
		name       string
		id         uint64
		token      string
		wantStatus int
	}{
		{
			name:       "should return status 200",
			id:         pet.ID,
			token:      shelterToken,
			wantStatus: http.StatusOK,
		},
		{
			name:       "should return status 404 when pet not found",
			id:         999,
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
			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/pets/%d", tt.id), nil)
			req.Header.Set("Authorization", "Bearer "+tt.token)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			if tt.wantStatus == http.StatusOK {
				var result schemas.PetResponse
				err = json.Unmarshal(w.Body.Bytes(), &result)
				require.NoError(t, err)

				assert.NotEmpty(t, result.ID)
				assert.Equal(t, pet.Name, result.Name)
				assert.Equal(t, pet.Description, result.Description)
				assert.Equal(t, pet.Photo, result.Photo)
				assert.Equal(t, pet.Age, result.Age)
				assert.Equal(t, pet.IsAdopted, result.IsAdopt)
				assert.Equal(t, pet.UF, result.UF)
				assert.Equal(t, pet.City, result.City)
				assert.Equal(t, pet.ShelterID, result.ShelterID)
			}
		})
	}
}
