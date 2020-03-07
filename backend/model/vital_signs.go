package model

// VitalSigns are describing the vital signs in a medical case
type VitalSigns struct {
	// Onset of symptoms
	OoS                    string       `json:"oos" form:"oos" bson:"oos"`
	AVPU                   string       `json:"avpu" form:"avpu" bson:"avpu"`
	Mobility               string       `json:"mobility" form:"mobility" bson:"mobility"`
	RespiratoryRate        *int         `json:"respiratoryRate" form:"respiratoryRate" bson:"respiratoryRate"`
	Pulse                  *int         `json:"pulse" form:"pulse" bson:"pulse"`
	Temperature            *float32     `json:"temperature" form:"temperature" bson:"temperature"`
	CapillaryRefill        *float32     `json:"capillaryRefill" form:"capillaryRefill" bson:"capillaryRefill"`
	BloodPressureSystolic  *int         `json:"bloodPressureSystolic" form:"bloodPressureSystolic" bson:"bloodPressureSystolic"`
	BloodPressureDiastolic *int         `json:"bloodPressureDiastolic" form:"bloodPressureDiastolic" bson:"bloodPressureDiastolic"`
	OxygenSaturation       *int         `json:"oxygenSaturation" form:"oxygenSaturation" bson:"oxygenSaturation"`
	Expectations           Expectations `json:"expectations" bson:"Expectations"`
}

// Expectations are
type Expectations struct {
	Foe               string `json:"foe" bson:"foe"`
	TreatmentExpected string `json:"treatmentExpected" bson:"treatmentExpected"`
}
