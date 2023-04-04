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
