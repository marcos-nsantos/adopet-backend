package adoptionhandler

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

func TestCreateAdoption(t *testing.T) {
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

	pet := mock.Pet()[0]
	pet.ShelterID = shelterCreated.ID
	petCreated, err := database.CreatePet(pet)
	require.NoError(t, err)

	tutor := mock.Tutors()[0]
	tutorCreated, err := database.CreateTutor(tutor)
	require.NoError(t, err)

	tutorToken, err := auth.GenerateToken(shelterCreated.ID, entity.TutorType)
	require.NoError(t, err)

	tests := []struct {
		name  string
		url   string
		token string
		want  int
	}{
		{
			name:  "should adopt a pet",
			url:   fmt.Sprintf("/adoptions/%d/%d", petCreated.ID, tutorCreated.ID),
			token: tutorToken,
			want:  http.StatusCreated,
		},
		{
			name: "should not adopt a pet when token is not provided",
			url:  fmt.Sprintf("/adoptions/%d/%d", petCreated.ID, tutorCreated.ID),
			want: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, tt.url, nil)
			req.Header.Set("Authorization", "Bearer "+tt.token)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.want, w.Code)
			if tt.want == http.StatusCreated {
				var result schemas.AdoptionResponse
				err = json.Unmarshal(w.Body.Bytes(), &result)
				require.NoError(t, err)

				assert.NotEmpty(t, result.ID)
				assert.Equal(t, petCreated.ID, result.PetID)
				assert.Equal(t, tutorCreated.ID, result.TutorID)
				assert.NotEmpty(t, result.CreatedAt)
			}
		})
	}
}
