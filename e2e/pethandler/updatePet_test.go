package pethandler

import (
	"bytes"
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

func TestUpdatePet(t *testing.T) {
	database.InitTest()
	database.Migrate()
	gin.SetMode(gin.TestMode)
	r := router.SetupRoutes()

	t.Cleanup(func() {
		database.DropTables()
	})

	pet := mock.Pet()[0]
	pet, err := database.CreatePet(pet)
	require.NoError(t, err)

	tests := []struct {
		name       string
		id         uint64
		reqBody    schemas.PetUpdateRequests
		wantStatus int
	}{
		{
			name: "should return status 200",
			id:   pet.ID,
			reqBody: schemas.PetUpdateRequests{
				Name:        "Pet One Updated",
				Description: "Description of pet one updated",
				Age:         1,
				Photo:       "https://somephoto.com/photo.png",
				UF:          "SP",
				City:        "São Paulo",
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "should return status 422 when name is empty",
			id:   pet.ID,
			reqBody: schemas.PetUpdateRequests{
				Name:        "",
				Description: "Description of pet one updated",
				Age:         1,
				Photo:       "https://somephoto.com/photo.png",
				UF:          "SP",
				City:        "São Paulo",
			},
			wantStatus: http.StatusUnprocessableEntity,
		},
		{
			name: "should return status 404 when pet is not found",
			id:   999,
			reqBody: schemas.PetUpdateRequests{
				Name:        "Pet One Updated",
				Description: "Description of pet one updated",
				Age:         1,
				Photo:       "https://somephoto.com/photo.png",
				UF:          "SP",
				City:        "São Paulo",
			},
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody, err := json.Marshal(tt.reqBody)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/pets/%d", tt.id), bytes.NewReader(reqBody))
			require.NoError(t, err)

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}
