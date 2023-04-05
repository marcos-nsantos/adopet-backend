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
