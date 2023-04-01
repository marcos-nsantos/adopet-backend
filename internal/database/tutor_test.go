package database

import (
	"testing"

	"github.com/marcos-nsantos/adopet-backend/internal/mock"
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
