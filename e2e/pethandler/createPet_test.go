package pethandler

import (
	"bytes"
	"encoding/json"
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

func TestCreatePet(t *testing.T) {
	database.InitTest()
	database.Migrate()
	gin.SetMode(gin.TestMode)
	r := router.SetupRoutes()

	t.Cleanup(func() {
		database.DropTables()
	})

	shelter := mock.Shelters()[0]
	shelterCreated, err := database.CreateUser(shelter)
	require.NoError(t, err)
	pet := mock.Pet()[0]

	tests := []struct {
		name       string
		reqBody    schemas.PetCreateRequests
		wantStatus int
	}{
		{
			name: "should create a pet and return status 201",
			reqBody: schemas.PetCreateRequests{
				Name:        pet.Name,
				Description: pet.Description,
				Photo:       pet.Photo,
				Age:         pet.Age,
				IsAdopt:     pet.IsAdopt,
				UF:          pet.UF,
				City:        pet.City,
				UserID:      shelterCreated.ID,
			},
			wantStatus: http.StatusCreated,
		},
		{
			name: "should return status 422 when name is empty",
			reqBody: schemas.PetCreateRequests{
				Name:        "",
				Description: pet.Description,
				Photo:       pet.Photo,
				Age:         pet.Age,
				IsAdopt:     pet.IsAdopt,
				UF:          pet.UF,
				City:        pet.City,
				UserID:      shelterCreated.ID,
			},
			wantStatus: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBody, err := json.Marshal(tt.reqBody)
			require.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, "/pets", bytes.NewBuffer(jsonBody))
			require.NoError(t, err)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}
