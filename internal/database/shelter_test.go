package database

import (
	"github.com/marcos-nsantos/adopet-backend/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetShelterByID(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	shelter := mock.Shelters()[0]
	shelterCreated, err := CreateUser(shelter)
	require.NoError(t, err)

	t.Run("should get a shelter by id", func(t *testing.T) {
		shelterFound, err := GetShelterByID(shelterCreated.ID)
		require.NoError(t, err)

		assert.Equal(t, shelterCreated.ID, shelterFound.ID)
		assert.Equal(t, shelterCreated.Name, shelterFound.Name)
		assert.Equal(t, shelterCreated.Email, shelterFound.Email)
		assert.Empty(t, shelterFound.Password)
		assert.Equal(t, shelterCreated.Phone, shelterFound.Phone)
		assert.Equal(t, shelterCreated.Photo, shelterFound.Photo)
		assert.Equal(t, shelterCreated.City, shelterFound.City)
		assert.Equal(t, shelterCreated.About, shelterFound.About)
	})

	t.Run("should not get a shelter by id when shelter does not exist", func(t *testing.T) {
		_, err := GetShelterByID(0)
		require.Error(t, err)
	})

	t.Run("should an error when user is not a shelter", func(t *testing.T) {
		tutor := mock.Tutors()[1]
		tutorCreated, err := CreateUser(tutor)
		require.NoError(t, err)

		_, err = GetShelterByID(tutorCreated.ID)
		assert.Error(t, err)
	})
}

func TestGetAllShelters(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	shelters := mock.Shelters()
	DB.CreateInBatches(shelters, len(shelters))

	t.Run("should get all shelters", func(t *testing.T) {
		sheltersFound, total, err := GetAllShelters(1, 10)
		require.NoError(t, err)

		assert.Equal(t, len(shelters), total)
		assert.Equal(t, len(shelters), len(sheltersFound))
	})

	t.Run("should get all shelters with pagination", func(t *testing.T) {
		sheltersFound, total, err := GetAllShelters(1, 2)
		require.NoError(t, err)

		assert.Equal(t, len(shelters), total)
		assert.Equal(t, 2, len(sheltersFound))
	})
}
