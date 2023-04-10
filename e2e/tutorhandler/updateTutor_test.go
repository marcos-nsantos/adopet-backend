package tutorhandler

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func TestUpdateTutor(t *testing.T) {
	database.InitTest()
	database.Migrate()
	gin.SetMode(gin.TestMode)
	r := router.SetupRoutes()

	t.Cleanup(func() {
		database.DropTables()
	})

	tutor := mock.Tutors()[0]
	tutor, err := database.CreateTutor(tutor)
	require.NoError(t, err)

	tutorToken, err := auth.GenerateToken(tutor.ID, entity.TutorType)
	require.NoError(t, err)

	tests := []struct {
		name       string
		id         uint64
		reqBody    schemas.TutorUpdateRequest
		token      string
		wantStatus int
	}{
		{
			name: "should return status 200",
			id:   tutor.ID,
			reqBody: schemas.TutorUpdateRequest{
				Name:  "Tutor One Updated",
				Email: "tutoroneupdated@email.com",
				Phone: "99999999999",
				Photo: "https://tutoroneupdatedphoto.com/tutor.jpg",
				City:  "Rio Branco",
				About: "Hi there, I am updated",
			},
			token:      tutorToken,
			wantStatus: http.StatusOK,
		},
		{
			name:       "should return status 422 when body is invalid",
			id:         tutor.ID,
			reqBody:    schemas.TutorUpdateRequest{Name: "", Email: "tutoroneupdatedemail.com"},
			token:      tutorToken,
			wantStatus: http.StatusUnprocessableEntity,
		},
		{
			name: "should return status 404 when tutor not found",
			id:   999,
			reqBody: schemas.TutorUpdateRequest{
				Name:  "Tutor One Updated",
				Email: "tutoroneupdated@email.com",
				Phone: "99999999999",
				Photo: "https://tutoroneupdatedphoto.com/tutor.jpg",
				City:  "Rio Branco",
				About: "Hi there, I am updated",
			},
			token:      tutorToken,
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "should return status 401 when token is not provided",
			id:         tutor.ID,
			reqBody:    schemas.TutorUpdateRequest{},
			wantStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody, err := json.Marshal(tt.reqBody)
			require.NoError(t, err)

			req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/tutors/%d", tt.id), bytes.NewBuffer(reqBody))
			req.Header.Set("Authorization", "Bearer "+tt.token)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			if tt.wantStatus == http.StatusOK {
				var result schemas.TutorResponse
				err = json.Unmarshal(w.Body.Bytes(), &result)
				require.NoError(t, err)

				assert.Equal(t, tt.reqBody.Name, result.Name)
				assert.Equal(t, tt.reqBody.Email, result.Email)
			}
		})
	}
}
