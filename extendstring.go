package localsystem

import "strings"

func ConcatInStr(data []string) string {
	if len(data) <= 0 {
		return ""
	}
	var sb strings.Builder
	for i := 0; i < len(data); i++ {
		tmp := ",'" + data[i] + "'"
		sb.WriteString(tmp)
	}
	if len(sb.String()) < 2 {
		return ""
	}
	return sb.String()[1:]
}

func InStrings(target string, str_array []string) bool {
	for _, element := range str_array {
		if target == element {
			return true
		}
	}
	return false
}
