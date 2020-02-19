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

// TeamQuery is the query fields for the getter
type TeamQuery struct {
	Title    string `query:"title"`
	Author   string `query:"author"`
	Page     int    `query:"page"`
	PageSize int    `query:"limit"`
}

// TeamsList is a list response of teams
type TeamsList struct {
	Count int64  `json:"count"`
	Teams []Team `json:"teams"`
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
