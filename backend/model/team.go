package model

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Team describes a team taking part in the exercises
type Team struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Author    LimitedUser        `json:"author" bson:"author"`
	Editor    LimitedUser        `json:"editor" bson:"editor"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	EditedAt  time.Time          `json:"editedAt" bson:"editedAt"`
	Title     string             `json:"title" bson:"title"`
	Type      string             `json:"type" bson:"type"`
	Medivac   bool               `json:"medivac" bson:"medivac"`
}

// Validate validates a team
func (team *Team) Validate() error {
	if team.Author.ID.IsZero() || team.Author.Username == "" {
		return errors.New("author not set")
	}
	if team.CreatedAt.IsZero() {
		return errors.New("created at not set")
	}
	if team.Title == "" {
		return errors.New("title not set")
	}
	return nil
}
