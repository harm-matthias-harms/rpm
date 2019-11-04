package model

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MedicalCase describes a medical case for the injects
type MedicalCase struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Author           LimitedUser        `json:"author" bson:"author"`
	Editor           LimitedUser        `json:"editor" bson:"editor"`
	CreatedAt        time.Time          `json:"createdAt" bson:"createdAt"`
	EditedAt         time.Time          `json:"editedAt" bson:"editedAt"`
	Title            string             `json:"title" bson:"title"`
	MakeUp           string             `json:"makeup" bson:"makeup"`
	OtherInformation string             `json:"other_information" bson:"other_information"`
	// TODO Documents
	GeneralInformation struct {
		Surgical      bool   `json:"surgical" bson:"surgical"`
		Hospilisation bool   `json:"hospilisation" bson:"hospilisation"`
		USAR          bool   `json:"usar" bson:"usar"`
		Medicav       bool   `json:"medivac" bson:"medivac"`
		Triage        string `json:"triage" bson:"triage"`
		ShortSummary  string `json:"short_summary" bson:"short_summary"`
		Age           string `json:"age" bson:"age"`
		Gender        string `json:"gender" bson:"gender"`
	} `json:"general_information" bson:"general_information"`
	MedicalHistroy struct {
		Problems         string `json:"problems" bson:"problems"`
		Vaccinations     string `json:"vaccinations" bson:"vaccinations"`
		Allergies        string `json:"allergies" bson:"allergies"`
		Medication       string `json:"medication" bson:"medication"`
		ImplantedDevices string `json:"implanted_devices" bson:"implanted_devices"`
	} `json:"medical_history" bson:"medical_history"`
	Expectations struct {
		GeneralStatus string `json:"general_status" bson:"general_status"`
		OnExamination string `json:"on_examination" bson:"on_examination"`
		Expectations  string `json:"expectations" bson:"expectations"`
	} `json:"expectations" bson:"expectations"`
	// TODO VitalSigns graph
	VitalSigns []struct {
		Title      string     `json:"title" bson:"title"`
		Prepending string     `json:"prepending" bson:"prepending"`
		Reason     string     `jso:"reason" bson:"reason"`
		Data       VitalSigns `json:"data" bson:"data"`
	} `json:"vital_signs" bson:"vital_signs"`
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
	Count   int64         `json:"count"`
	MedicalCases []MedicalCaseShort `json:"medical_cases"`
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
