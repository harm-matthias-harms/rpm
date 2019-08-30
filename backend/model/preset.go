package model

import (
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
