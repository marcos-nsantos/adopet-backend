package adoptionhandler

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/entity"
	"github.com/marcos-nsantos/adopet-backend/internal/mock"
	"github.com/marcos-nsantos/adopet-backend/internal/router"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeleteAdoption(t *testing.T) {
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

	adoptionCreated := entity.Adoption{
		PetID:   petCreated.ID,
		TutorID: tutorCreated.ID,
	}
	err = database.Adopt(&adoptionCreated)
	require.NoError(t, err)

	tests := []struct {
		name string
		id   uint64
	}{
		{
			name: "should delete adoption",
			id:   adoptionCreated.ID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/adoptions/%d", tt.id), bytes.NewBuffer([]byte{}))
			require.NoError(t, err)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusNoContent, w.Code)
		})
	}
}
