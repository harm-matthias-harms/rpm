package model

// VitalSigns are describing the vital signs in a medical case
type VitalSigns struct {
	// Onset of symptoms
	OoS      string `json:"oos" bson:"oos"`
	AVPU     string `json:"avpu" bson:"avpu"`
	Mobility string `json:"mobility" bson:"mobility"`
	// Respiratory rate
	RespiratoryRate        int     `json:"respiratoryRate" bson:"respiratoryRate"`
	Pulse                  int     `json:"pulse" bson:"pulse"`
	Temperature            float32 `json:"temperature" bson:"temperature"`
	CapillaryRefill        float32 `json:"capillaryRefill" bson:"capillaryRefill"`
	BloodPressureSystolic  int     `json:"bloodPressureSystolic" bson:"bloodPressureSystolic"`
	BloodPressureDiastolic int     `json:"bloodPressureDiastolic" bson:"bloodPressureDiastolic"`
	OxygenSaturation       int     `json:"oxygenSaturation" bson:"oxygenSaturation"`
	Weight                 float32 `json:"weight" bson:"weight"`
	Height                 int     `json:"height" bson:"height"`
}
