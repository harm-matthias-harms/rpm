package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedicalCaseValidate(t *testing.T) {
	mc := MedicalCase{}

	err := mc.Validate()
	assert.EqualError(t, err, "title not set")

	mc.Title = "title"
	err = mc.Validate()
	assert.NoError(t, err)
}
