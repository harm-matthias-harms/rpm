package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	assert.Equal(t, "test", GetEnv("SOME_RANDOM_ENV", "test"))
	os.Setenv("SOME_RANDOM_ENV", "1")
	assert.Equal(t, "1", GetEnv("SOME_RANDOM_ENV", "test"))
}
