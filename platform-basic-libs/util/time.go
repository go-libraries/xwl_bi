package util

import "time"

const (
	TimeFormat     = "2006-01-02 15:04:05"
	TimeFormatDay  = "20060102"
	TimeFormatDay2 = "2006-01-02"
	TimeFormatDay3 = "2006/01/02"
	TimeFormatDay4  = "2006.01.02_15"
)

/**
 * 二个时间戳是否同一天
 * @return true 是 false 不是今天
 */
func IsSameDay(oldDay, anotherDay int64) bool {
	tm := time.Unix(oldDay, 0)
	tmAnother := time.Unix(anotherDay, 0)
	if tm.Format(TimeFormatDay2) == tmAnother.Format(TimeFormatDay2) {
		return true
	}
	return false
}

/**字符串->时间对象*/
func Str2Time(formatTimeStr,timeFormat string) time.Time {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeFormat, formatTimeStr, loc) //使用模板在对应时区转化为time.time类型
	return theTime
}