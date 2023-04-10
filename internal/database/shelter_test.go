package database

import (
	"testing"

	"github.com/marcos-nsantos/adopet-backend/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateShelter(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	shelter := mock.Shelters()[0]
	t.Run("should create a shelter", func(t *testing.T) {
		shelterCreated, err := CreateShelter(shelter)
		require.NoError(t, err)

		assert.NotEmpty(t, shelterCreated.ID)
		assert.Equal(t, shelter.Name, shelterCreated.Name)
		assert.Equal(t, shelter.Email, shelterCreated.Email)
		assert.Equal(t, shelter.Password, shelterCreated.Password)
		assert.Equal(t, shelter.Phone, shelterCreated.Phone)
		assert.Equal(t, shelter.Photo, shelterCreated.Photo)
		assert.Equal(t, shelter.City, shelterCreated.City)
		assert.Equal(t, shelter.About, shelterCreated.About)
		assert.NotEmpty(t, shelterCreated.CreatedAt)
	})
}

func TestGetShelterByID(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	shelter := mock.Shelters()[0]
	shelterCreated, err := CreateShelter(shelter)
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

	t.Run("should return error when shelter is not found", func(t *testing.T) {
		_, err := GetShelterByID(999)
		require.Error(t, err)
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

func TestUpdateShelter(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	shelter := mock.Shelters()[0]
	shelterCreated, err := CreateShelter(shelter)
	require.NoError(t, err)

	t.Run("should update a shelter", func(t *testing.T) {
		shelterCreated.Name = "New Name"
		shelterCreated.Email = "newshelter@email.com"

		err := UpdateShelter(&shelterCreated)
		require.NoError(t, err)

		result, err := GetShelterByID(shelterCreated.ID)
		require.NoError(t, err)

		assert.Equal(t, shelterCreated.ID, result.ID)
		assert.Equal(t, shelterCreated.Name, result.Name)
		assert.Equal(t, shelterCreated.Email, result.Email)
	})
}

func TestDeleteShelter(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	shelter := mock.Shelters()[0]
	shelterCreated, err := CreateShelter(shelter)
	require.NoError(t, err)

	t.Run("should delete a shelter", func(t *testing.T) {
		err := DeleteShelter(shelterCreated.ID)
		require.NoError(t, err)

		_, err = GetShelterByID(shelterCreated.ID)
		assert.Error(t, err)
	})
}

func TestGetIDAndPasswordByEmailFromShelter(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	shelter := mock.Shelters()[0]
	shelterCreated, err := CreateShelter(shelter)
	require.NoError(t, err)

	t.Run("should get id and password from shelter", func(t *testing.T) {
		id, password, err := GetIDAndPasswordByEmailFromShelter(shelterCreated.Email)
		require.NoError(t, err)

		assert.Equal(t, shelterCreated.ID, id)
		assert.Equal(t, shelterCreated.Password, password)
	})

	t.Run("should return error when shelter is not found", func(t *testing.T) {
		_, _, err := GetIDAndPasswordByEmailFromShelter("othershelter@email.com")
		require.Error(t, err)
	})
}
