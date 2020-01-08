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
	CreatedAt  time.Time          `json:"createdAt" bson:"createdAt"`
	EditedAt   time.Time          `json:"editedAt" bson:"editedAt"`
	Title      string             `json:"title" bson:"title"`
	VitalSigns VitalSigns         `json:"vitalSigns" bson:"vitalSigns"`
}

// PresetQuery is the query fields for the getter
type PresetQuery struct {
	Title    string `query:"title"`
	Author   string `query:"author"`
	Page     int    `query:"page"`
	PageSize int    `query:"limit"`
}

// PresetShort describes a short preset for a list of presets
type PresetShort struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Author LimitedUser        `json:"author" bson:"author"`
	Title  string             `json:"title" bson:"title"`
}

// PresetsList is a list response of presets
type PresetsList struct {
	Count   int64         `json:"count"`
	Presets []PresetShort `json:"presets"`
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
