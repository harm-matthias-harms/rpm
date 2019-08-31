package model

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Preset describes a preset for vital signs
type Preset struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Author     LimitedUser        `json:"author" bson:"author"`
	Editor     LimitedUser        `json:"editor" bson:"editor"`
	CreatedAt  time.Time          `json:"created_at,omitempty" bson:"created_at"`
	EditedAt   time.Time          `json:"edited_at,omitempty" bson:"edited_at"`
	Title      string             `json:"title" bson:"title"`
	VitalSigns VitalSigns         `json:"vital_sign" bson:"vital_sign"`
}

// Validate validates a preset
func (preset *Preset) Validate() error {
	if preset.Author.ID.IsZero() || preset.Author.Username == "" {
		return errors.New("Author not set")
	}
	if preset.CreatedAt.IsZero() {
		return errors.New("Created at not set")
	}
	if preset.Title == "" {
		return errors.New("Title not set")
	}
	if preset.VitalSigns == (VitalSigns{}) {
		return errors.New("Vital signs not set")
	}
	return nil
}
