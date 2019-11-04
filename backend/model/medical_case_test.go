package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMedicalCaseValidate(t *testing.T) {
	mc := MedicalCase{}
	err := mc.Validate()
	assert.EqualError(t, err, "author not set")
	mc.Author.ID = primitive.NewObjectID()
	mc.Author.Username = "username"
	err = mc.Validate()
	assert.EqualError(t, err, "created at not set")
	mc.CreatedAt = time.Now()
	err = mc.Validate()
	assert.EqualError(t, err, "title not set")
	mc.Title = "title"
	err = mc.Validate()
	assert.NoError(t, err)
}
