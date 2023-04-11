package pethandler

import (
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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeletePet(t *testing.T) {
	database.Init()
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
	petCreated, err := database.CreatePet(pet)
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
			name:       "should return status 204",
			id:         petCreated.ID,
			token:      shelterToken,
			wantStatus: http.StatusNoContent,
		},
		{
			name:       "should return status 404 when pet not found",
			id:         petCreated.ID,
			token:      shelterToken,
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "should return status 401 when token is not provided",
			id:         petCreated.ID,
			wantStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/pets/%d", tt.id), nil)
			req.Header.Set("Authorization", "Bearer "+tt.token)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}
