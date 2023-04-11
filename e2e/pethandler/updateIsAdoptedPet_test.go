package pethandler

import (
	"bytes"
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

func TestUpdateIsAdoptedPet(t *testing.T) {
	database.Init()
	database.Migrate()
	gin.SetMode(gin.TestMode)
	r := router.SetupRoutes()

	t.Cleanup(func() {
		database.DropTables()
	})

	tutor := mock.Tutors()[0]
	tutor, err := database.CreateTutor(tutor)
	require.NoError(t, err)

	pet := mock.Pet()[0]
	petCreated, err := database.CreatePet(pet)
	require.NoError(t, err)

	tutorToken, err := auth.GenerateToken(tutor.ID, entity.TutorType)
	require.NoError(t, err)

	tests := []struct {
		name       string
		id         uint64
		reqBody    schemas.UpdateIsAdoptPetRequests
		token      string
		wantStatus int
	}{
		{
			name:       "should return status 204",
			id:         petCreated.ID,
			reqBody:    schemas.UpdateIsAdoptPetRequests{IsAdopted: !pet.IsAdopted},
			token:      tutorToken,
			wantStatus: http.StatusNoContent,
		},
		{
			name:       "should return status 404 when pet is not found",
			id:         999,
			reqBody:    schemas.UpdateIsAdoptPetRequests{IsAdopted: !pet.IsAdopted},
			token:      tutorToken,
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody, err := json.Marshal(tt.reqBody)
			require.NoError(t, err)

			req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("/pets/%d/adopted", tt.id), bytes.NewReader(reqBody))
			req.Header.Set("Authorization", "Bearer "+tt.token)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}
