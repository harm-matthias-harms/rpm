package model

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Inject describes an medical inject of the exercise
type Inject struct {
	ID                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Author            LimitedUser        `json:"author" bson:"author"`
	Editor            LimitedUser        `json:"editor" bson:"editor"`
	CreatedAt         time.Time          `json:"createdAt" bson:"createdAt"`
	EditedAt          time.Time          `json:"editedAt" bson:"editedAt"`
	ExerciseID        primitive.ObjectID `json:"exerciseID" bson:"exerciseID"`
	Status            string             `json:"status" bson:"status"`
	StartTime         time.Time          `json:"startTime" bson:"startTime"`
	MakeupCenterTitle string             `json:"makeupCenterTitle" bson:"makeupCenterTitle"`
	Team              Team               `json:"team" bson:"team"`
	Roleplayer        struct {
		Fullname    string `json:"fullName" bson:"fullName"`
		Gender      string `json:"gender" bson:"gender"`
		Age         int    `json:"age" bson:"age"`
		Nationality string `json:"nationality" bson:"nationality"`
	} `json:"roleplayer" bson:"roleplayer"`
	MedicalCase MedicalCase `json:"medicalCase" bson:"medicalCase"`
}

// InjectShort describes a minimal inject to display it in a list
type InjectShort struct {
	ID                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Author            LimitedUser        `json:"author" bson:"author"`
	Editor            LimitedUser        `json:"editor" bson:"editor"`
	CreatedAt         time.Time          `json:"createdAt" bson:"createdAt"`
	EditedAt          time.Time          `json:"editedAt" bson:"editedAt"`
	ExerciseID        primitive.ObjectID `json:"exerciseID" bson:"exerciseID"`
	Status            string             `json:"status" bson:"status"`
	StartTime         time.Time          `json:"startTime" bson:"startTime"`
	MakeupCenterTitle string             `json:"makeupCenterTitle" bson:"makeupCenterTitle"`
	Team              Team               `json:"team" bson:"team"`
	Roleplayer        Roleplayer         `json:"roleplayer" bson:"roleplayer"`
	MedicalCase       MedicalCaseShort   `json:"medicalCase" bson:"medicalCase"`
}

// Roleplayer describes the role player who is playing a medical case
type Roleplayer struct {
	Fullname    string `json:"fullName" bson:"fullName"`
	Gender      string `json:"gender" bson:"gender"`
	Age         int    `json:"age" bson:"age"`
	Nationality string `json:"nationality" bson:"nationality"`
}

// InjectQuery is the query fields for the getter
type InjectQuery struct {
	ExerciseID        primitive.ObjectID `query:"exerciseId"`
	Author            string             `query:"author"`
	Team              Team               `query:"team"`
	MakeupCenterTitle string             `query:"makeupCenterTitle"`
	Status            string             `query:"status"`
	Page              int                `query:"page"`
	PageSize          int                `query:"limit"`
}

// InjectList is a list response of inject
type InjectList struct {
	Count   int64         `json:"count"`
	Injects []InjectShort `json:"injects"`
}

// Validate validates an inject for its completeness
func (inject *Inject) Validate() (err error) {
	if inject.Author.ID.IsZero() || inject.Author.Username == "" {
		return errors.New("author not set")
	}
	if inject.CreatedAt.IsZero() {
		return errors.New("created at not set")
	}
	if inject.ExerciseID.IsZero() {
		return errors.New("exercise id is not set")
	}
	if err = inject.Team.Validate(); err != nil {
		return errors.New("team is not set")
	}
	if err = inject.MedicalCase.Validate(); err != nil {
		return errors.New("medical case is not set")
	}
	if inject.Roleplayer == (Roleplayer{}) {
		return errors.New("role player is not set")
	}
	return
}
