package model

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Exercise is a played exercise with different roles and injects
type Exercise struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Author          LimitedUser        `json:"author" bson:"author"`
	CreatedAt       time.Time          `json:"createdAt" bson:"createdAt"`
	EditedAt        time.Time          `json:"editedAt" bson:"editedAt"`
	Title           string             `json:"title" bson:"title"`
	StartTime       time.Time          `json:"startTime" bson:"startTime"`
	EndTime         time.Time          `json:"endTime" bson:"endTime"`
	Teams           []ExerciseTeam     `json:"teams" bson:"teams"`
	RoleplayManager []LimitedUser      `json:"roleplayManager" bson:"roleplayManager"`
	MakeupCenter    []MakeupCenter     `json:"makeupCenter" bson:"makeupCenter"`
}

// ExerciseShort describes the short version of an exercise
type ExerciseShort struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	StartTime time.Time          `json:"startTime" bson:"startTime"`
	EndTime   time.Time          `json:"endTime" bson:"endTime"`
}

// ExerciseTeam is the combination of team and trainer
type ExerciseTeam struct {
	Team    Team        `json:"team" bson:"team"`
	Trainer LimitedUser `json:"trainer" bson:"trainer"`
}

// MakeupCenter is doing the makeup for the roleplayer
type MakeupCenter struct {
	Title   string      `json:"title" bson:"title"`
	Account LimitedUser `json:"account" bson:"account"`
}

// Validate validates an exercise if all necessary fields are filed. (title, start and end time, teams, role-play manager and make-up center)
func (e *Exercise) Validate() (err error) {
	if e.Author.ID.IsZero() || e.Author.Username == "" {
		return errors.New("author not set")
	}
	if e.CreatedAt.IsZero() {
		return errors.New("created at not set")
	}
	if e.Title == "" {
		return errors.New("no title provided")
	}
	if e.StartTime.IsZero() || e.EndTime.IsZero() {
		return errors.New("no start or end time provided")
	}
	if len(e.Teams) == 0 {
		return errors.New("no teams provided")
	}
	if len(e.RoleplayManager) == 0 {
		return errors.New("no role-play manager provided")
	}
	if len(e.MakeupCenter) == 0 {
		return errors.New("no make-up center provided")
	}
	return
}
