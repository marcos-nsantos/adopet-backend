package database

import (
	"github.com/marcos-nsantos/adopet-backend/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateUser(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	user := mock.Tutors()[0]
	t.Run("should create a user", func(t *testing.T) {
		userCreated, err := CreateUser(user)
		require.NoError(t, err)

		assert.NotEmpty(t, userCreated.ID)
		assert.Equal(t, user.Name, userCreated.Name)
		assert.Equal(t, user.Email, userCreated.Email)
		assert.Equal(t, user.Password, userCreated.Password)
		assert.Equal(t, user.Type, userCreated.Type)
		assert.Equal(t, user.Phone, userCreated.Phone)
		assert.Equal(t, user.Photo, userCreated.Photo)
		assert.Equal(t, user.City, userCreated.City)
		assert.Equal(t, user.About, userCreated.About)
		assert.NotEmpty(t, userCreated.CreatedAt)
	})

	t.Run("should not create a user when email is already in use", func(t *testing.T) {
		_, err := CreateUser(user)
		require.Error(t, err)
	})
}

func TestUpdateUser(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	users := mock.Tutors()
	DB.CreateInBatches(users, len(users))

	t.Run("should update a users", func(t *testing.T) {
		user := users[0]
		user.Name = "User One Updated"
		user.Email = "useroneupdate@email.com"
		user.Phone = "99999999999"
		user.Photo = "https://www.alura.com.br"
		user.City = "Rio de Janeiro"
		user.About = "I am a user updated"

		err := UpdateUser(&user)
		require.NoError(t, err)

		userFound, err := GetTutorByID(user.ID)
		require.NoError(t, err)

		assert.Equal(t, user.ID, userFound.ID)
		assert.Equal(t, user.Name, userFound.Name)
		assert.Equal(t, user.Email, userFound.Email)
		assert.Equal(t, user.Phone, userFound.Phone)
		assert.Equal(t, user.Photo, userFound.Photo)
		assert.Equal(t, user.City, userFound.City)
		assert.Equal(t, user.About, userFound.About)
	})

	t.Run("should not update a user when email is already in use", func(t *testing.T) {
		user := users[0]
		user.Email = users[1].Email

		err := UpdateUser(&user)
		require.Error(t, err)
	})
}

func TestDeleteUser(t *testing.T) {
	Init()
	Migrate()

	t.Cleanup(func() {
		DropTables()
	})

	user, err := CreateUser(mock.Tutors()[0])
	require.NoError(t, err)

	t.Run("should delete a user", func(t *testing.T) {
		err := DeleteUser(user.ID)
		require.NoError(t, err)

		_, err = GetTutorByID(user.ID)
		assert.Error(t, err)
	})
}
