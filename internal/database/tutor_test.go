package database

import (
	"testing"

	"github.com/marcos-nsantos/adopet-backend/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTutorByID(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	tutor := mock.Tutors()[0]
	tutorCreated, err := CreateUser(tutor)
	require.NoError(t, err)

	t.Run("should get a tutor by id", func(t *testing.T) {
		tutorFound, err := GetTutorByID(tutorCreated.ID)
		require.NoError(t, err)
		assert.Equal(t, tutorCreated.ID, tutorFound.ID)
		assert.Equal(t, tutorCreated.Name, tutorFound.Name)
		assert.Equal(t, tutorCreated.Email, tutorFound.Email)
		assert.Empty(t, tutorFound.Password)
		assert.Equal(t, tutorCreated.Phone, tutorFound.Phone)
		assert.Equal(t, tutorCreated.Photo, tutorFound.Photo)
		assert.Equal(t, tutorCreated.City, tutorFound.City)
		assert.Equal(t, tutorCreated.About, tutorFound.About)
	})

	t.Run("should not get a tutor by id when tutor does not exist", func(t *testing.T) {
		_, err := GetTutorByID(0)
		require.Error(t, err)
	})
}

func TestGetAllTutors(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	tutors := mock.Tutors()
	DB.CreateInBatches(tutors, len(tutors))

	t.Run("should get all tutors", func(t *testing.T) {
		tutorsFound, total, err := GetAllTutors(1, 10)
		require.NoError(t, err)
		assert.Equal(t, len(tutors), total)
		assert.Len(t, tutorsFound, len(tutors))
	})

	t.Run("should get all tutors with pagination", func(t *testing.T) {
		tutorsFound, total, err := GetAllTutors(1, 2)
		require.NoError(t, err)
		assert.Equal(t, len(tutors), total)
		assert.Len(t, tutorsFound, 2)
	})
}

func TestUpdateTutor(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	tutors := mock.Tutors()
	DB.CreateInBatches(tutors, len(tutors))

	t.Run("should update a tutor", func(t *testing.T) {
		tutor := tutors[0]
		tutor.Name = "Tutor One Updated"
		tutor.Email = "tutoroneupdate@email.com"
		tutor.Phone = "99999999999"
		tutor.Photo = "https://www.alura.com.br"
		tutor.City = "Rio de Janeiro"
		tutor.About = "I am a tutor updated"

		err := UpdateTutor(&tutor)
		require.NoError(t, err)

		tutorFound, err := GetTutorByID(tutor.ID)
		require.NoError(t, err)

		assert.Equal(t, tutor.ID, tutorFound.ID)
		assert.Equal(t, tutor.Name, tutorFound.Name)
		assert.Equal(t, tutor.Email, tutorFound.Email)
		assert.Equal(t, tutor.Phone, tutorFound.Phone)
		assert.Equal(t, tutor.Photo, tutorFound.Photo)
		assert.Equal(t, tutor.City, tutorFound.City)
		assert.Equal(t, tutor.About, tutorFound.About)
	})

	t.Run("should not update a tutor when email is already in use", func(t *testing.T) {
		tutor := tutors[0]
		tutor.Email = tutors[1].Email

		err := UpdateTutor(&tutor)
		require.Error(t, err)
	})
}

func TestDeleteTutor(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	tutor, err := CreateUser(mock.Tutors()[0])
	require.NoError(t, err)

	t.Run("should delete a tutor", func(t *testing.T) {
		err := DeleteTutor(tutor.ID)
		require.NoError(t, err)

		_, err = GetTutorByID(tutor.ID)
		assert.Error(t, err)
	})
}
