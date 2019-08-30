package model

// VitalSigns are describing the vital signs in a medical case
type VitalSigns struct {
	// Onset of symptoms
	OoS      string `json:"oos" bson:"oos"`
	AVPU     string `json:"avpu" bson:"avpu"`
	Mobility string `json:"mobility" bson:"mobility"`
	// Respiratory rate
	RespRate               int     `json:"resp_rate" bson:"resp_rate"`
	Pulse                  int     `json:"pulse" bson:"pulse"`
	Temperature            float32 `json:"temperature" bson:"temperature"`
	CapillaryRefill        float32 `json:"capillary_refill" bson:"capillary_refill"`
	BloodPressureSystolic  int     `json:"blood_pressure_systolic" bson:"blood_pressure_systolic"`
	BloodPressureDiastolic int     `json:"blood_pressure_diastolic" bson:"blood_pressure_diastolic"`
	OxygenSaturation       int     `json:"oxygen_saturation" bson:"oxygen_saturation"`
	Weight                 float32 `json:"weight" bson:"weight"`
	Height                 int     `json:"height" bson:"height"`
}
