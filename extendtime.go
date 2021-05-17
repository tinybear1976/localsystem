package localsystem

import (
	"fmt"
	"strconv"
	"time"
)

//第一时间为字符串，第二时间为time，t2-t1=相差秒,第一个时间应该小于第二个时间
func DiffSecByStr(firsttime string, endtime time.Time) int {
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", firsttime, time.Local)
	sub := endtime.Sub(t1)
	s2 := strconv.FormatFloat(sub.Seconds(), 'f', 0, 64)
	sec, err := strconv.Atoi(s2)
	if err != nil {
		sec = -1
	}
	return sec
}

//第一时间为字符串，第二时间为字符串，第三个为日期格式，t2-t1=相差秒,第一个时间应该小于第二个时间，如果出现错误返回-1
func DiffSecByStrWithFormat(ts1, ts2, formattemp string) int {
	t1, _ := time.ParseInLocation(formattemp, ts1, time.Local)
	t2, _ := time.ParseInLocation(formattemp, ts2, time.Local)
	sub := t2.Sub(t1)
	s2 := strconv.FormatFloat(sub.Seconds(), 'f', 0, 64)
	sec, err := strconv.Atoi(s2)
	if err != nil {
		sec = -1
	}
	return sec
}

//第一时间为字符串，第二时间为字符串，第三个为日期格式，t2-t1=相差秒,第一个时间应该小于第二个时间，如果出现错误返回-1
func DiffSecByStrWithFormatErr(ts1, ts2, formattemp string) (int, error) {
	t1, err := time.ParseInLocation(formattemp, ts1, time.Local)
	if err != nil {
		return -1, err
	}
	t2, err := time.ParseInLocation(formattemp, ts2, time.Local)
	if err != nil {
		return -1, err
	}
	sub := t2.Sub(t1)
	s2 := strconv.FormatFloat(sub.Seconds(), 'f', 0, 64)
	sec, err := strconv.Atoi(s2)
	if err != nil {
		sec = -1
	}
	return sec, nil
}

/*seconds为计算秒数,如果为负数，函数会自动现将其调整为正数;
  模板参数"d:hh:mm:ss"为全格式，可以逐级减项，"d:hh:mm" "hh:mm:ss" "hh:mm" "mm:ss"，最少两项
*/
func SecondsToString(seconds int, template string) string {
	if seconds < 0 {
		seconds = -seconds
	}

	var day, hour, minute, second int
	if seconds > 0 {
		secondsPerMinute := 60
		secondsPerHour := secondsPerMinute * 60
		secondsPerDay := secondsPerHour * 24

		day = seconds / secondsPerDay
		tmp := seconds % secondsPerDay
		hour = tmp / secondsPerHour
		tmp = tmp % secondsPerHour
		minute = tmp / secondsPerMinute
		second = tmp % secondsPerMinute
	}
	switch template {
	case "d:hh:mm:ss":
		return fmt.Sprintf("%d:%02d:%02d:%02d", day, hour, minute, second)
	case "d:hh:mm":
		return fmt.Sprintf("%d:%02d:%02d", day, hour, minute)
	case "hh:mm:ss":
		return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
	case "hh:mm":
		return fmt.Sprintf("%02d:%02d", hour, minute)
	case "mm:ss":
		return fmt.Sprintf("%02d:%02d", minute, second)
	}
	return ""
}

/*将秒数转换成字符串格式，指定返回模板的最高位为完整计算数，可能会超过2位，
  比如mm:ss则表示所有的描述最大单位以分钟计数，
  seconds为计算秒数,如果为负数，函数会自动现将其调整为正数;
  模板参数"dd:hh:mm:ss"为全格式，可以逐级减项，"dd:hh:mm" "hh:mm:ss" "hh:mm" "mm:ss"，最少两项
  如果模板截断最低位ss，结果将忽略秒数部分
*/
func SecondsToStringTimelong(seconds int, template string) string {
	if seconds < 0 {
		seconds = -seconds
	}

	switch template {
	case "d:hh:mm:ss":
		var day, hour, minute, second int
		secondsPerMinute := 60
		secondsPerHour := secondsPerMinute * 60
		secondsPerDay := secondsPerHour * 24
		day = seconds / secondsPerDay
		tmp := seconds % secondsPerDay
		hour = tmp / secondsPerHour
		tmp = tmp % secondsPerHour
		minute = tmp / secondsPerMinute
		second = tmp % secondsPerMinute
		tmpl := "%d:%02d:%02d:%02d"
		return fmt.Sprintf(tmpl, day, hour, minute, second)
	case "d:hh:mm":
		var day, hour, minute int
		secondsPerMinute := 60
		secondsPerHour := secondsPerMinute * 60
		secondsPerDay := secondsPerHour * 24
		day = seconds / secondsPerDay
		tmp := seconds % secondsPerDay
		hour = tmp / secondsPerHour
		tmp = tmp % secondsPerHour
		minute = tmp / secondsPerMinute
		tmpl := "%d:%02d:%02d"
		return fmt.Sprintf(tmpl, day, hour, minute)
	case "h:mm:ss":
		var hour, minute, second int
		secondsPerMinute := 60
		secondsPerHour := secondsPerMinute * 60

		hour = seconds / secondsPerHour
		tmp := seconds % secondsPerHour
		minute = tmp / secondsPerMinute
		second = tmp % secondsPerMinute
		tmpl := "%d:%02d:%02d"
		return fmt.Sprintf(tmpl, hour, minute, second)
	case "h:mm":
		var hour, minute int
		secondsPerMinute := 60
		secondsPerHour := secondsPerMinute * 60

		hour = seconds / secondsPerHour
		tmp := seconds % secondsPerHour
		minute = tmp / secondsPerMinute
		tmpl := "%d:%02d"
		return fmt.Sprintf(tmpl, hour, minute)
	case "m:ss":
		var minute, second int
		secondsPerMinute := 60

		minute = seconds / secondsPerMinute
		second = seconds % secondsPerMinute
		tmpl := "%d:%02d"
		return fmt.Sprintf(tmpl, minute, second)
	}
	return ""
}

//字符串转时间戳，格式字符串，实际时间
func String2Timestamp(layout, value string) (time.Time, error) {
	t, err := time.ParseInLocation(layout, value, time.Local)
	return t, err
}
