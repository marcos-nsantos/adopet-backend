package tutorhandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/handler/tutorhandler"
	"github.com/marcos-nsantos/adopet-backend/internal/mock"
	"github.com/marcos-nsantos/adopet-backend/internal/router"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTutorByID(t *testing.T) {
	database.Init()
	database.Migrate()
	gin.SetMode(gin.TestMode)
	r := router.SetupRoutes()

	t.Cleanup(func() {
		database.DropTables()
	})

	tutor := mock.Tutors()[0]
	tutorCreated, err := database.CreateUser(tutor)
	require.NoError(t, err)

	tests := []struct {
		name       string
		id         uint64
		wantStatus int
	}{
		{
			name:       "should return status 200",
			id:         tutorCreated.ID,
			wantStatus: http.StatusOK,
		},
		{
			name:       "should return status 404 when tutor not found",
			id:         999,
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/tutors/%d", tt.id), nil)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			if tt.wantStatus == http.StatusOK {
				var tutor tutorhandler.UserResponse
				err = json.Unmarshal(w.Body.Bytes(), &tutor)
				require.NoError(t, err)

				assert.Equal(t, tutorCreated.ID, tutor.ID)
				assert.Equal(t, tutorCreated.Name, tutor.Name)
				assert.Equal(t, tutorCreated.Email, tutor.Email)
				assert.Equal(t, tutorCreated.Phone, tutor.Phone)
				assert.Equal(t, tutorCreated.Photo, tutor.Photo)
				assert.Equal(t, tutorCreated.City, tutor.City)
				assert.Equal(t, tutorCreated.About, tutor.About)
			}
		})
	}
}
