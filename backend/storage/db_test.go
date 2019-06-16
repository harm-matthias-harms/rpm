package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMongoRunning(t *testing.T) {
	setMongoDatabase()
	assert.NotEqual(t, nil, mongoSession)
}
