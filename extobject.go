package localsystem

import "encoding/json"

// 目标dest对象因为需要返回，所以需要取地址传入
func CopyObject(source interface{}, dest interface{}) {
	sbytes, _ := json.Marshal(source)
	json.Unmarshal(sbytes, dest)
}
