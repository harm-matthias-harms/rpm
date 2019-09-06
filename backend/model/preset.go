package model

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Preset describes a preset for vital signs
type Preset struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Author     LimitedUser        `json:"author" bson:"author"`
	Editor     LimitedUser        `json:"editor" bson:"editor"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	EditedAt   time.Time          `json:"edited_at" bson:"edited_at"`
	Title      string             `json:"title" bson:"title"`
	VitalSigns VitalSigns         `json:"vital_signs" bson:"vital_signs"`
}

// Validate validates a preset
func (preset *Preset) Validate() error {
	if preset.Author.ID.IsZero() || preset.Author.Username == "" {
		return errors.New("author not set")
	}
	if preset.CreatedAt.IsZero() {
		return errors.New("created at not set")
	}
	if preset.Title == "" {
		return errors.New("title not set")
	}
	if preset.VitalSigns == (VitalSigns{}) {
		return errors.New("vital signs not set")
	}
	return nil
}
