package tutorhandler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/marcos-nsantos/adopet-backend/internal/schemas"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/mock"
	"github.com/marcos-nsantos/adopet-backend/internal/router"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAllUsers(t *testing.T) {
	database.InitTest()
	database.Migrate()
	gin.SetMode(gin.TestMode)
	r := router.SetupRoutes()

	t.Cleanup(func() {
		database.DropTables()
	})

	users := mock.Tutors()
	database.DB.CreateInBatches(users, len(users))

	tests := []struct {
		name       string
		url        string
		wantStatus int
	}{
		{
			name:       "should return status 200",
			url:        "/tutors",
			wantStatus: http.StatusOK,
		},
		{
			name:       "should return status 200",
			url:        "/tutors?page=1&limit=2",
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tt.url, nil)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)

			if tt.wantStatus == http.StatusOK {
				var reqBody schemas.UsersResponse
				err := json.Unmarshal(w.Body.Bytes(), &reqBody)
				require.NoError(t, err)
				assert.GreaterOrEqual(t, len(reqBody.Users), 2)
				assert.Equal(t, reqBody.Total, len(users))
			}
		})
	}
}
