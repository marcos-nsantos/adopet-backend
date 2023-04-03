package database

import (
	"github.com/marcos-nsantos/adopet-backend/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateTutor(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	user := mock.Tutors()[0]
	t.Run("should create a user", func(t *testing.T) {
		tutorCreated, err := CreateUser(user)
		require.NoError(t, err)

		assert.NotEmpty(t, tutorCreated.ID)
		assert.Equal(t, user.Name, tutorCreated.Name)
		assert.Equal(t, user.Email, tutorCreated.Email)
		assert.Equal(t, user.Password, tutorCreated.Password)
		assert.Equal(t, user.Type, tutorCreated.Type)
		assert.Equal(t, user.Phone, tutorCreated.Phone)
		assert.Equal(t, user.Photo, tutorCreated.Photo)
		assert.Equal(t, user.City, tutorCreated.City)
		assert.Equal(t, user.About, tutorCreated.About)
		assert.NotEmpty(t, tutorCreated.CreatedAt)
	})

	t.Run("should not create a user when email is already in use", func(t *testing.T) {
		_, err := CreateUser(user)
		require.Error(t, err)
	})
}
