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
	assert.EqualError(t, err, "empty or none vital signs exist")
	mc.VitalSigns = append(mc.VitalSigns, struct {
		Title      string     `json:"title" bson:"title"`
		Prepending string     `json:"prepending" bson:"prepending"`
		Reason     string     `jso:"reason" bson:"reason"`
		Data       VitalSigns `json:"data" bson:"data"`
	}{Title: "title", Data: VitalSigns{OoS: "oos"}})
	err = mc.Validate()
	assert.NoError(t, err)
}

func TestMedicalCaseVSPresent(t *testing.T) {
	mc := MedicalCase{}
	assert.Equal(t, false, mc.vsPresent())
	mc.VitalSigns = append(mc.VitalSigns, struct {
		Title      string     `json:"title" bson:"title"`
		Prepending string     `json:"prepending" bson:"prepending"`
		Reason     string     `jso:"reason" bson:"reason"`
		Data       VitalSigns `json:"data" bson:"data"`
	}{Title: "title"})
	assert.Equal(t, false, mc.vsPresent())
	mc.VitalSigns[0].Data.OoS = "oos"
	assert.Equal(t, true, mc.vsPresent())
}
