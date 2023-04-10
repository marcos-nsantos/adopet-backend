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
		tutorCreated, err := CreateTutor(tutor)
		require.NoError(t, err)

		assert.NotEmpty(t, tutorCreated.ID)
		assert.Equal(t, tutor.Name, tutorCreated.Name)
		assert.Equal(t, tutor.Email, tutorCreated.Email)
		assert.Equal(t, tutor.Password, tutorCreated.Password)
		assert.Equal(t, tutor.Phone, tutorCreated.Phone)
		assert.Equal(t, tutor.Photo, tutorCreated.Photo)
		assert.Equal(t, tutor.City, tutorCreated.City)
		assert.Equal(t, tutor.About, tutorCreated.About)
		assert.NotEmpty(t, tutorCreated.CreatedAt)
	})

	t.Run("should not create a tutor with an existing email", func(t *testing.T) {
		_, err := CreateTutor(tutor)
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
	tutorCreated, err := CreateTutor(tutor)
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

	t.Run("should return error when tutor not found", func(t *testing.T) {
		_, err := GetTutorByID(999)
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

	tutor := mock.Tutors()[0]
	tutorCreated, err := CreateTutor(tutor)
	require.NoError(t, err)

	t.Run("should update a tutor", func(t *testing.T) {
		tutorCreated.Name = "New Name"
		tutorCreated.Email = "newtutor@email.com"

		err := UpdateTutor(&tutorCreated)
		require.NoError(t, err)

		result, err := GetTutorByID(tutorCreated.ID)
		require.NoError(t, err)

		assert.Equal(t, tutorCreated.ID, result.ID)
		assert.Equal(t, tutorCreated.Name, result.Name)
		assert.Equal(t, tutorCreated.Email, result.Email)
	})
}

func TestDeleteTutor(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	tutor := mock.Tutors()[0]
	tutorCreated, err := CreateTutor(tutor)
	require.NoError(t, err)

	t.Run("should delete a tutor", func(t *testing.T) {
		err := DeleteTutor(tutorCreated.ID)
		require.NoError(t, err)

		_, err = GetTutorByID(tutorCreated.ID)
		assert.Error(t, err)
	})
}

func TestGetIDAndPasswordByEmailFromTutor(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	tutor := mock.Tutors()[0]
	tutorCreated, err := CreateTutor(tutor)
	require.NoError(t, err)

	t.Run("should get id and password by email", func(t *testing.T) {
		id, password, err := GetIDAndPasswordByEmailFromTutor(tutorCreated.Email)
		require.NoError(t, err)

		assert.Equal(t, tutorCreated.ID, id)
		assert.Equal(t, tutorCreated.Password, password)
	})

	t.Run("should return error when tutor not found", func(t *testing.T) {
		_, _, err := GetIDAndPasswordByEmailFromTutor("otheremail@email.com")
		require.Error(t, err)
	})
}
