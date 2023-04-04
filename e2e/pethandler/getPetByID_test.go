package pethandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/mock"
	"github.com/marcos-nsantos/adopet-backend/internal/router"
	"github.com/marcos-nsantos/adopet-backend/internal/schemas"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPetByID(t *testing.T) {
	database.Init()
	database.Migrate()
	gin.SetMode(gin.TestMode)
	r := router.SetupRoutes()

	t.Cleanup(func() {
		database.DropTables()
	})

	pet := mock.Pet()[0]
	pet, err := database.CreatePet(pet)
	require.NoError(t, err)

	tests := []struct {
		name       string
		id         uint64
		wantStatus int
	}{
		{
			name:       "should return status 200",
			id:         pet.ID,
			wantStatus: http.StatusOK,
		},
		{
			name:       "should return status 404 when pet not found",
			id:         999,
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/pets/%d", tt.id), nil)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			if tt.wantStatus == http.StatusOK {
				var bodyResult schemas.PetResponse
				err = json.Unmarshal(w.Body.Bytes(), &bodyResult)
				require.NoError(t, err)

				assert.Equal(t, pet.Name, bodyResult.Name)
				assert.Equal(t, pet.Description, bodyResult.Description)
				assert.Equal(t, pet.Photo, bodyResult.Photo)
				assert.Equal(t, pet.Age, bodyResult.Age)
				assert.Equal(t, pet.IsAdopt, bodyResult.IsAdopt)
				assert.Equal(t, pet.UF, bodyResult.UF)
				assert.Equal(t, pet.City, bodyResult.City)
			}
		})
	}
}
