package utils

import (
	"encoding/json"
)

// Convert converts between two different types sharing same fields
func Convert(t1 interface{}, t2 interface{}) {
	str, _ := json.Marshal(t1)
	json.Unmarshal(str, t2)
}
