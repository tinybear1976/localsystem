package localsystem

import "encoding/json"

func CopyObject(source interface{}, dest interface{}) {
	sbytes, _ := json.Marshal(source)
	json.Unmarshal(sbytes, dest)
}
