package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMongoRunning(t *testing.T) {
	err := SetMongoDatabase()
	assert.NoError(t, err)
	assert.NotEqual(t, nil, MongoSession)
}
