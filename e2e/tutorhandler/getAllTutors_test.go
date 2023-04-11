package tutorhandler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/marcos-nsantos/adopet-backend/internal/auth"
	"github.com/marcos-nsantos/adopet-backend/internal/entity"
	"github.com/marcos-nsantos/adopet-backend/internal/schemas"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/mock"
	"github.com/marcos-nsantos/adopet-backend/internal/router"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAllUsers(t *testing.T) {
	database.Init()
	database.Migrate()
	gin.SetMode(gin.TestMode)
	r := router.SetupRoutes()

	t.Cleanup(func() {
		database.DropTables()
	})

	users := mock.Tutors()
	database.DB.CreateInBatches(users, len(users))

	tutorToken, err := auth.GenerateToken(users[0].ID, entity.TutorType)
	require.NoError(t, err)

	tests := []struct {
		name       string
		url        string
		token      string
		wantStatus int
	}{
		{
			name:       "should return status 200",
			url:        "/tutors",
			token:      tutorToken,
			wantStatus: http.StatusOK,
		},
		{
			name:       "should return status 200",
			url:        "/tutors?page=1&limit=2",
			token:      tutorToken,
			wantStatus: http.StatusOK,
		},
		{
			name:       "should return status 401 when token is not provided",
			url:        "/tutors",
			wantStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tt.url, nil)
			req.Header.Set("Authorization", "Bearer "+tt.token)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)

			if tt.wantStatus == http.StatusOK {
				var result schemas.TutorsResponse
				err := json.Unmarshal(w.Body.Bytes(), &result)
				require.NoError(t, err)

				assert.GreaterOrEqual(t, len(result.Tutors), 2)
				assert.Equal(t, result.Total, len(users))
			}
		})
	}
}
