package tutorhandler

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func TestUpdateTutor(t *testing.T) {
	database.InitTest()
	database.Migrate()
	gin.SetMode(gin.TestMode)
	r := router.SetupRoutes()

	t.Cleanup(func() {
		database.DropTables()
	})

	tutor := mock.Tutors()[0]
	tutor, err := database.CreateUser(tutor)
	require.NoError(t, err)

	tests := []struct {
		name       string
		id         uint64
		reqBody    schemas.UserUpdateRequest
		wantStatus int
	}{
		{
			name:       "should return status 200",
			id:         tutor.ID,
			reqBody:    schemas.UserUpdateRequest{Name: "Tutor One Updated", Email: "tutoroneupdated@email.com"},
			wantStatus: http.StatusOK,
		},
		{
			name:       "should return status 422 when email is invalid or name is empty",
			id:         tutor.ID,
			reqBody:    schemas.UserUpdateRequest{Name: "", Email: "tutoroneupdatedemail.com"},
			wantStatus: http.StatusUnprocessableEntity,
		},
		{
			name:       "should return status 404 when tutor not found",
			id:         999,
			reqBody:    schemas.UserUpdateRequest{Name: "Tutor One Updated", Email: "tutoroneupdated@email.com"},
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody, err := json.Marshal(tt.reqBody)
			require.NoError(t, err)

			req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/tutors/%d", tt.id), bytes.NewBuffer(reqBody))
			require.NoError(t, err)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			if tt.wantStatus == http.StatusOK {
				var tutor schemas.UserResponse
				err = json.Unmarshal(w.Body.Bytes(), &tutor)
				require.NoError(t, err)

				assert.Equal(t, tt.reqBody.Name, tutor.Name)
				assert.Equal(t, tt.reqBody.Email, tutor.Email)
			}
		})
	}
}
