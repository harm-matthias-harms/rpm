package model

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MedicalCase describes a medical case for the injects
type MedicalCase struct {
	ID        primitive.ObjectID `json:"id" form:"id" bson:"_id,omitempty"`
	Author    LimitedUser        `json:"author" form:"author" bson:"author"`
	Editor    LimitedUser        `json:"editor" form:"editor" bson:"editor"`
	CreatedAt time.Time          `json:"createdAt" form:"createdAt" bson:"createdAt"`
	EditedAt  time.Time          `json:"editedAt" form:"editedAt" bson:"editedAt"`
	Approved  bool               `json:"approved" form:"approved" bson:"approved"`
	Title     string             `json:"title" form:"title" bson:"title"`
	General   struct {
		Discipline []string `json:"discipline" form:"discipline" bson:"discipline"`
		Context    []string `json:"context" form:"context" bson:"context"`
		Scenario   []string `json:"scenario" form:"scenario" bson:"scenario"`
	} `json:"general" form:"general" bson:"general"`
	Patient struct {
		Type   string   `json:"type" form:"type" bson:"type"`
		Triage string   `json:"triage" form:"triage" bson:"triage"`
		Gender []string `json:"gender" form:"gender" bson:"gender"`
		Age    string   `json:"age" form:"age" bson:"age"`
	} `json:"patient" form:"patient" bson:"patient"`
	Medical struct {
		Signs      string `json:"signs" form:"signs" bson:"signs"`
		Allergies  string `json:"allergies" form:"allergies" bson:"allergies"`
		Medication string `json:"medication" form:"medication" bson:"medication"`
		Past       string `json:"past" form:"past" bson:"past"`
		Loi        string `json:"loi" form:"loi" bson:"loi"`
		Events     string `json:"events" form:"events" bson:"events"`
	} `json:"medical" form:"medical" bson:"medical"`
	MakeUp struct {
		MakeUp string `json:"makeup" form:"makeup" bson:"makeup"`
		Acting string `json:"acting" form:"acting" bson:"acting"`
	} `json:"makeup" form:"makeup" bson:"makeup"`
	VitalSigns []NestedVitalSigns `json:"vitalSigns" form:"vitalSigns" bson:"vitalSigns"`
	Files      []MedicalCaseFile  `json:"files" bson:"files"`
}

//NestedVitalSigns enable evolving cases
type NestedVitalSigns struct {
	Title  string             `json:"title" form:"title" bson:"title"`
	Data   VitalSigns         `json:"data" form:"data" bson:"data"`
	Childs []NestedVitalSigns `json:"childs" form:"childs" bson:"childs"`
}

// MedicalCaseFile describes a document stored at a medical case
type MedicalCaseFile struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
	Size int64              `json:"size" bson:"size"`
}

// MedicalCaseShort serves as the short Version of medical cases for lists.
type MedicalCaseShort struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Author   LimitedUser        `json:"author" bson:"author"`
	Title    string             `json:"title" bson:"title"`
	Approved bool               `json:"approved" bson:"approved"`
	General  struct {
		Discipline string   `json:"discipline" form:"discipline" bson:"discipline"`
		Context    []string `json:"context" form:"context" bson:"context"`
		Scenario   []string `json:"scenario" form:"scenario" bson:"scenario"`
	} `json:"general" form:"general" bson:"general"`
}

// MedicalCaseQuery is the query fields for the getter
type MedicalCaseQuery struct {
	Title    string `query:"title"`
	Author   string `query:"author"`
	Page     int    `query:"page"`
	PageSize int    `query:"limit"`
}

// MedicalCaseList is a list response of medical cases
type MedicalCaseList struct {
	Count        int64              `json:"count"`
	MedicalCases []MedicalCaseShort `json:"medicalCases"`
}

// Validate validates the medical case
func (mc *MedicalCase) Validate() error {
	if mc.Title == "" {
		return errors.New("title not set")
	}
	return nil
}
