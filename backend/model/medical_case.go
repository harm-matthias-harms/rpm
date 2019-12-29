package model

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MedicalCase describes a medical case for the injects
type MedicalCase struct {
	ID               primitive.ObjectID `json:"id" form:"id" bson:"_id,omitempty"`
	Author           LimitedUser        `json:"author" form:"author" bson:"author"`
	Editor           LimitedUser        `json:"editor" form:"editor" bson:"editor"`
	CreatedAt        time.Time          `json:"createdAt" form:"createdAt" bson:"createdAt"`
	EditedAt         time.Time          `json:"editedAt" form:"editedAt" bson:"editedAt"`
	Title            string             `json:"title" form:"title" bson:"title"`
	MakeUp           string             `json:"makeup" form:"makeup" bson:"makeup"`
	OtherInformation string             `json:"otherInformation" form:"otherInformation" bson:"otherInformation"`
	GeneralInformation struct {
		Surgical      bool   `json:"surgical" form:"surgical" bson:"surgical"`
		Hospilisation bool   `json:"hospilisation" form:"hospilisation" bson:"hospilisation"`
		USAR          bool   `json:"usar" form:"usar" bson:"usar"`
		Medicav       bool   `json:"medivac" form:"medivac" bson:"medivac"`
		Triage        string `json:"triage" form:"triage" bson:"triage"`
		ShortSummary  string `json:"shortSummary" form:"shortSummary" bson:"shortSummary"`
		Age           string `json:"age" form:"age" bson:"age"`
		Gender        string `json:"gender" form:"gender" bson:"gender"`
	} `json:"generalInformation" form:"generalInformation" bson:"generalInformation"`
	MedicalHistroy struct {
		Problems         string `json:"problems" form:"problems" bson:"problems"`
		Vaccinations     string `json:"vaccinations" form:"vaccinations" bson:"vaccinations"`
		Allergies        string `json:"allergies" form:"allergies" bson:"allergies"`
		Medication       string `json:"medication" form:"medication" bson:"medication"`
		ImplantedDevices string `json:"implantedDevices" form:"implantedDevices" bson:"implantedDevices"`
	} `json:"medicalHistory" form:"medicalHistory" bson:"medicalHistory"`
	Expectations struct {
		GeneralStatus string `json:"generalStatus" form:"generalStatus" bson:"generalStatus"`
		OnExamination string `json:"onExamination" form:"onExamination" bson:"onExamination"`
		Expectations  string `json:"expectations" form:"expectations" bson:"expectations"`
	} `json:"expectations" form:"expectations" bson:"expectations"`
	VitalSigns []NestedVitalSigns `json:"vitalSigns" form:"vitalSigns" bson:"vitalSigns"`
	Files      []MedicalCaseFile  `json:"files" bson:"files"`
}

//NestedVitalSigns enable evolving cases
type NestedVitalSigns struct {
	Title  string     `json:"title" form:"title" bson:"title"`
	Reason string     `jso:"reason" form:"reason" bson:"reason"`
	Data   VitalSigns `json:"data" form:"data" bson:"data"`
	Childs []NestedVitalSigns
}

// MedicalCaseFile describes a document stored at a medical case
type MedicalCaseFile struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
	Size int64              `json:"size" bson:"size"`
}

// MedicalCaseShort serves as the short Version of medical cases for lists.
type MedicalCaseShort struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Author LimitedUser        `json:"author" bson:"author"`
	Title  string             `json:"title" bson:"title"`
}

// MedicalCaseQuery is the query fields for the getter
type MedicalCaseQuery struct {
	Title    string `query:"title"`
	Author   string `query:"author"`
	Page     int    `query:"page"`
	PageSize int    `query:"limit"`
}

// MedicalCaseList is a list response of presets
type MedicalCaseList struct {
	Count        int64              `json:"count"`
	MedicalCases []MedicalCaseShort `json:"medicalCases"`
}

// Validate validates the medical case
func (mc *MedicalCase) Validate() error {
	if mc.Author.ID.IsZero() || mc.Author.Username == "" {
		return errors.New("author not set")
	}
	if mc.CreatedAt.IsZero() {
		return errors.New("created at not set")
	}
	if mc.Title == "" {
		return errors.New("title not set")
	}
	return nil
}
