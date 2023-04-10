package shelterhandler

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

func TestGetShelterByID(t *testing.T) {
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
			id:         shelterCreated.ID,
			token:      shelterToken,
			wantStatus: http.StatusOK,
		},
		{
			name:       "should return status 404 when shelter not found",
			id:         999,
			token:      shelterToken,
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "should return status 401 when token is not provided",
			id:         shelterCreated.ID,
			wantStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/shelters/%d", tt.id), nil)
			req.Header.Set("Authorization", "Bearer "+tt.token)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			if tt.wantStatus == http.StatusOK {
				var result schemas.ShelterResponse
				err = json.Unmarshal(w.Body.Bytes(), &result)
				require.NoError(t, err)

				assert.Equal(t, shelterCreated.ID, result.ID)
				assert.Equal(t, shelterCreated.Name, result.Name)
				assert.Equal(t, shelterCreated.Email, result.Email)
				assert.Equal(t, shelterCreated.Phone, result.Phone)
				assert.Equal(t, shelterCreated.Photo, result.Photo)
				assert.Equal(t, shelterCreated.City, result.City)
				assert.Equal(t, shelterCreated.About, result.About)
			}
		})
	}
}
