package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {
	password := "secretPassword123"
	hash, err := Hash(password)
	require.NoError(t, err)
	assert.NotEqual(t, password, hash)
}

func TestCompare(t *testing.T) {
	password := "secretPassword123"
	hash, err := Hash(password)
	require.NoError(t, err)

	t.Run("should return nil when password is correct", func(t *testing.T) {
		err = Compare(hash, password)
		assert.NoError(t, err)
	})

	t.Run("should return error when password is incorrect", func(t *testing.T) {
		err = Compare(hash, "wrongPassword")
		assert.Error(t, err)
	})
}
