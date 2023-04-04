package shelterhandler

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

func TestUpdateShelter(t *testing.T) {
	database.InitTest()
	database.Migrate()
	gin.SetMode(gin.TestMode)
	r := router.SetupRoutes()

	t.Cleanup(func() {
		database.DropTables()
	})

	shelter := mock.Shelters()[0]
	shelter, err := database.CreateUser(shelter)
	require.NoError(t, err)

	tests := []struct {
		name       string
		id         uint64
		reqBody    schemas.UserUpdateRequest
		wantStatus int
	}{
		{
			name: "should return status 200",
			id:   shelter.ID,
			reqBody: schemas.UserUpdateRequest{
				Name:  "Shelter One Updated",
				Email: "shelteroneupdated@email.com",
				Phone: "123456729",
				Photo: "https://cdn.dribbble.com/userupload/2624051/file/original-0c2e27a535ca15358be82cb68805de49.png?compress=1&resize=752x",
				About: "Shelter One Updated",
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "should return status 422 when email is invalid or name is empty",
			id:   shelter.ID,
			reqBody: schemas.UserUpdateRequest{
				Name:  "",
				Email: "shelteroneupdatedemail.com",
				Phone: "123456739",
				Photo: "https://cdn.dribbble.com/userupload/2624051/file/original-0c2e27a535ca15358be82cb68805de49.png?compress=1&resize=752x",
				About: "Shelter One Updated",
			},
			wantStatus: http.StatusUnprocessableEntity,
		},
		{
			name: "should return status 404 when shelter is not found",
			id:   999,
			reqBody: schemas.UserUpdateRequest{
				Name:  "Shelter One Updated",
				Email: "shelteroneupdated@email.com",
				Phone: "123456739",
				Photo: "https://cdn.dribbble.com/userupload/2624051/file/original-0c2e27a535ca15358be82cb68805de49.png?compress=1&resize=752x",
				About: "Shelter One Updated",
			},
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody, err := json.Marshal(tt.reqBody)
			require.NoError(t, err)

			req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/shelters/%d", tt.id), bytes.NewBuffer(reqBody))
			require.NoError(t, err)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			if tt.wantStatus == http.StatusOK {
				var shelterResp schemas.UserResponse
				err = json.Unmarshal(w.Body.Bytes(), &shelterResp)
				require.NoError(t, err)

				assert.Equal(t, tt.reqBody.Name, shelterResp.Name)
				assert.Equal(t, tt.reqBody.Email, shelterResp.Email)
				assert.Equal(t, tt.reqBody.Phone, shelterResp.Phone)
				assert.Equal(t, tt.reqBody.Photo, shelterResp.Photo)
				assert.Equal(t, tt.reqBody.About, shelterResp.About)
			}
		})
	}
}
