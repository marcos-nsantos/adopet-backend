package shelterhandler

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

func TestGetShelterByID(t *testing.T) {
	database.Init()
	database.Migrate()
	gin.SetMode(gin.TestMode)
	r := router.SetupRoutes()

	t.Cleanup(func() {
		database.DropTables()
	})

	shelter := mock.Shelters()[0]
	shelterCreated, err := database.CreateUser(shelter)
	require.NoError(t, err)

	tests := []struct {
		name       string
		id         uint64
		wantStatus int
	}{
		{
			name:       "should return status 200",
			id:         shelterCreated.ID,
			wantStatus: http.StatusOK,
		},
		{
			name:       "should return status 404 when shelter not found",
			id:         999,
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/shelters/%d", tt.id), nil)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			if tt.wantStatus == http.StatusOK {
				var shelterResponse schemas.UserResponse
				err = json.Unmarshal(w.Body.Bytes(), &shelterResponse)
				require.NoError(t, err)

				assert.Equal(t, shelterCreated.ID, shelterResponse.ID)
				assert.Equal(t, shelterCreated.Name, shelterResponse.Name)
				assert.Equal(t, shelterCreated.Email, shelterResponse.Email)
				assert.Equal(t, shelterCreated.Phone, shelterResponse.Phone)
				assert.Equal(t, shelterCreated.Photo, shelterResponse.Photo)
				assert.Equal(t, shelterCreated.City, shelterResponse.City)
				assert.Equal(t, shelterCreated.About, shelterResponse.About)
			}
		})
	}
}
