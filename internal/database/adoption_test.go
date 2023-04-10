package database

import (
	"testing"

	"github.com/marcos-nsantos/adopet-backend/internal/entity"
	"github.com/marcos-nsantos/adopet-backend/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdopt(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	shelter := mock.Shelters()[0]
	shelterCreated, err := CreateShelter(shelter)
	require.NoError(t, err)

	pet := mock.Pet()[0]
	pet.ShelterID = shelterCreated.ID
	petCreated, err := CreatePet(pet)
	require.NoError(t, err)

	tutor := mock.Tutors()[0]
	tutorCreated, err := CreateTutor(tutor)
	require.NoError(t, err)

	t.Run("should adopt a pet", func(t *testing.T) {
		adoption := entity.Adoption{
			PetID:   petCreated.ID,
			TutorID: tutorCreated.ID,
		}

		err := Adopt(&adoption)
		require.NoError(t, err)

		assert.NotEmpty(t, adoption.ID)
		assert.Equal(t, petCreated.ID, adoption.PetID)
		assert.Equal(t, tutorCreated.ID, adoption.TutorID)
		assert.NotEmpty(t, adoption.CreatedAt)
	})
}
