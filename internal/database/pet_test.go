package database

import (
	"testing"

	"github.com/marcos-nsantos/adopet-backend/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreatePet(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	t.Run("should create a pet", func(t *testing.T) {
		pet := mock.Pet()[0]
		result, err := CreatePet(pet)
		require.NoError(t, err)

		assert.Equal(t, pet.Name, result.Name)
		assert.Equal(t, pet.Description, result.Description)
		assert.Equal(t, pet.IsAdopt, result.IsAdopt)
		assert.Equal(t, pet.Age, result.Age)
		assert.Equal(t, pet.Photo, result.Photo)
		assert.Equal(t, pet.UF, result.UF)
		assert.Equal(t, pet.City, result.City)
	})
}

func TestGetPetByID(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	pet := mock.Pet()[0]
	result, err := CreatePet(pet)
	require.NoError(t, err)

	t.Run("should get a pet by id", func(t *testing.T) {
		result, err := GetPetByID(result.ID)
		require.NoError(t, err)

		assert.Equal(t, pet.Name, result.Name)
		assert.Equal(t, pet.Description, result.Description)
		assert.Equal(t, pet.IsAdopt, result.IsAdopt)
		assert.Equal(t, pet.Age, result.Age)
		assert.Equal(t, pet.Photo, result.Photo)
		assert.Equal(t, pet.UF, result.UF)
		assert.Equal(t, pet.City, result.City)
	})

	t.Run("should return error when pet is not found", func(t *testing.T) {
		_, err := GetPetByID(0)
		require.Error(t, err)
	})

	t.Run("should return error when pet is deleted", func(t *testing.T) {
		DB.Delete(&result)
		_, err := GetPetByID(result.ID)
		require.Error(t, err)
	})
}

func TestGetAllPets(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	pets := mock.Pet()
	DB.CreateInBatches(pets, len(pets))

	t.Run("should get all pets", func(t *testing.T) {
		result, total, err := GetAllPets(1, 10)
		require.NoError(t, err)

		assert.Equal(t, len(pets), total)
		assert.Equal(t, len(pets), len(result))
	})

	t.Run("should get all pets with limit of 2", func(t *testing.T) {
		result, total, err := GetAllPets(1, 2)
		require.NoError(t, err)

		assert.Equal(t, len(pets), total)
		assert.Equal(t, 2, len(result))
	})
}
