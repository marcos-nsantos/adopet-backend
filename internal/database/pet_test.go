package database

import (
	"testing"

	"github.com/marcos-nsantos/adopet-backend/internal/entity"
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

	user := mock.Shelters()[0]
	shelterCreated, err := CreateShelter(user)
	require.NoError(t, err)

	t.Run("should create a pet", func(t *testing.T) {
		pet := mock.Pet()[0]
		pet.ShelterID = shelterCreated.ID
		result, err := CreatePet(pet)
		require.NoError(t, err)

		assert.Equal(t, pet.Name, result.Name)
		assert.Equal(t, pet.Description, result.Description)
		assert.Equal(t, pet.IsAdopted, result.IsAdopted)
		assert.Equal(t, pet.Age, result.Age)
		assert.Equal(t, pet.Photo, result.Photo)
		assert.Equal(t, pet.UF, result.UF)
		assert.Equal(t, pet.City, result.City)
		assert.Equal(t, pet.ShelterID, result.ShelterID)
	})
}

func TestGetPetByID(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	user := mock.Shelters()[0]
	shelterCreated, err := CreateShelter(user)
	require.NoError(t, err)
	pet := mock.Pet()[0]
	pet.ShelterID = shelterCreated.ID
	result, err := CreatePet(pet)
	require.NoError(t, err)

	t.Run("should get a pet by id", func(t *testing.T) {
		result, err := GetPetByID(result.ID)
		require.NoError(t, err)

		assert.Equal(t, pet.Name, result.Name)
		assert.Equal(t, pet.Description, result.Description)
		assert.Equal(t, pet.IsAdopted, result.IsAdopted)
		assert.Equal(t, pet.Age, result.Age)
		assert.Equal(t, pet.Photo, result.Photo)
		assert.Equal(t, pet.UF, result.UF)
		assert.Equal(t, pet.City, result.City)
		assert.Equal(t, pet.ShelterID, result.ShelterID)
	})

	t.Run("should return error when pet is not found", func(t *testing.T) {
		_, err := GetPetByID(0)
		require.Error(t, err)
		assert.ErrorIs(t, err, entity.ErrPetNotFound)
	})

	t.Run("should return error when pet is deleted", func(t *testing.T) {
		DB.Delete(&result)
		_, err := GetPetByID(result.ID)
		require.Error(t, err)
		assert.ErrorIs(t, err, entity.ErrPetNotFound)
	})

	t.Run("should not return a pet it is adopted", func(t *testing.T) {
		result.IsAdopted = true
		DB.Save(&result)
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

	_, err := CreateShelter(mock.Shelters()[0])
	require.NoError(t, err)
	pets := mock.Pet()
	DB.CreateInBatches(pets, len(pets))

	t.Run("should get all pets", func(t *testing.T) {
		result, total, err := GetAllPets(1, 10)
		require.NoError(t, err)
		assert.NotEqual(t, 0, total)
		assert.NotEmpty(t, result)
	})

	t.Run("should get all pets with limit of 2", func(t *testing.T) {
		result, total, err := GetAllPets(1, 2)
		require.NoError(t, err)

		assert.NotEqual(t, len(pets), total)
		assert.Equal(t, 2, len(result))
	})
}

func TestUpdatePet(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	shelter := mock.Shelters()[0]
	shelterCreated, err := CreateShelter(shelter)
	pet := mock.Pet()[0]
	pet.ShelterID = shelterCreated.ID
	result, err := CreatePet(pet)
	require.NoError(t, err)

	t.Run("should update a pet", func(t *testing.T) {
		result.Name = "Test"
		err := UpdatePet(result)
		require.NoError(t, err)

		pet, err := GetPetByID(result.ID)
		require.NoError(t, err)

		assert.Equal(t, result.Name, pet.Name)
	})

	t.Run("should return error when pet is not found", func(t *testing.T) {
		result.ID = 999
		err = UpdatePet(result)
		require.Error(t, err)
		assert.ErrorIs(t, err, entity.ErrPetNotFound)
	})
}

func TestUpdateIsAdoptedPet(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	shelter := mock.Shelters()[0]
	shelterCreated, err := CreateShelter(shelter)
	pet := mock.Pet()[0]
	pet.ShelterID = shelterCreated.ID
	result, err := CreatePet(pet)
	require.NoError(t, err)

	t.Run("should update a pet", func(t *testing.T) {
		pet := entity.Pet{ID: result.ID, IsAdopted: true}
		err := UpdateIsAdoptedPet(pet)
		assert.NoError(t, err)
	})

	t.Run("should return error when pet is not found", func(t *testing.T) {
		pet := entity.Pet{ID: 999, IsAdopted: true}
		err := UpdateIsAdoptedPet(pet)
		assert.Error(t, err)
		assert.ErrorIs(t, err, entity.ErrPetNotFound)
	})
}

func TestDeletePet(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	shelter := mock.Shelters()[0]
	shelterCreated, err := CreateShelter(shelter)
	pet := mock.Pet()[0]
	pet.ShelterID = shelterCreated.ID
	result, err := CreatePet(pet)
	require.NoError(t, err)

	t.Run("should delete a pet", func(t *testing.T) {
		err := DeletePet(result.ID)
		require.NoError(t, err)

		_, err = GetPetByID(result.ID)
		require.Error(t, err)
	})

	t.Run("should return error when pet is not found", func(t *testing.T) {
		err := DeletePet(999)
		require.Error(t, err)
		assert.ErrorIs(t, err, entity.ErrPetNotFound)
	})
}
