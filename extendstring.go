package localsystem

import (
	"strconv"
	"strings"
)

// 将字符串数组连接成字符串。
// data  字符串数组, separator 间隔符(一般为,), delimiter 定界符（如果不需要可以为空字符串）
func ConcatStrings(data []string, separator string, delimiter string) string {
	if len(data) <= 0 {
		return ""
	}
	var sb strings.Builder
	for i := 0; i < len(data); i++ {
		tmp := separator + delimiter + data[i] + delimiter
		sb.WriteString(tmp)
	}
	if len(sb.String()) < 2 {
		return ""
	}
	return sb.String()[1:]
}

// 将整数切片连接成字符串。
// data  int数组, separator 间隔符(一般为,), delimiter 定界符（如果不需要可以为空字符串）
func ConcatStringFromInts(data []int, separator string, delimiter string) string {
	if len(data) <= 0 {
		return ""
	}
	var sb strings.Builder
	for i := 0; i < len(data); i++ {
		tmp := separator + delimiter + strconv.Itoa(data[i]) + delimiter
		sb.WriteString(tmp)
	}
	if len(sb.String()) < 2 {
		return ""
	}
	return sb.String()[1:]
}

// 字符串是否在字符串数组中存在
func InStrings(target string, str_array []string) bool {
	for _, element := range str_array {
		if target == element {
			return true
		}
	}
	return false
}

// 获得字符串结尾N个字符
func GetLastRune(str string, amount int) string {
	r := []rune(str)
	if len(r) >= amount {
		return string(r[len(r)-amount:])
	}
	return str
}

// 移除字符串结尾N个字符
func RemoveLastRune(str string, amount int) string {
	r := []rune(str)
	if len(r) >= amount {
		return string(r[:len(r)-amount])
	}
	return ""
}

// float64转换成字符串
func Float64toString(f64 float64) string {
	return strconv.FormatFloat(f64, 'f', -1, 64)
}
