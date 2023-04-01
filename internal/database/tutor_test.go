package database

import (
	"testing"

	"github.com/marcos-nsantos/adopet-backend/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateTutor(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	tutor := mock.Tutors()[0]
	t.Run("should create a tutor", func(t *testing.T) {
		err := CreateTutor(&tutor)
		require.NoError(t, err)
	})

	t.Run("should not create a tutor when email is already in use", func(t *testing.T) {
		err := CreateTutor(&tutor)
		require.Error(t, err)
	})
}

func TestGetTutorByID(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	tutor := mock.Tutors()[0]
	err := CreateTutor(&tutor)
	require.NoError(t, err)

	t.Run("should get a tutor by id", func(t *testing.T) {
		tutorFound, err := GetTutorByID(tutor.ID)
		require.NoError(t, err)
		assert.Equal(t, tutor.ID, tutorFound.ID)
		assert.Equal(t, tutor.Name, tutorFound.Name)
		assert.Equal(t, tutor.Email, tutorFound.Email)
		assert.Empty(t, tutorFound.Password)
		assert.Equal(t, tutor.Phone, tutorFound.Phone)
		assert.Equal(t, tutor.Photo, tutorFound.Photo)
		assert.Equal(t, tutor.City, tutorFound.City)
		assert.Equal(t, tutor.About, tutorFound.About)
		assert.Empty(t, tutorFound.DeletedAt)
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
