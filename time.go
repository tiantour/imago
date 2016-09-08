package imago

import (
	"fmt"
	"time"
)

// UnixToString 时间戳转字符串
func (t *itime) UnixToString(n int64) string {
	return time.Unix(n, 0).Format("2006-01-02 15:04:05")
}

// StringToUnix 字符串转时间戳
func (t *itime) StringToUnix(s string) (int64, error) {
	dateTime, err := time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
	if err != nil {
		return 0, err
	}
	n := dateTime.Unix()
	return n, nil
}

// TimeToDate
func (t *itime) TimeToDate(dateTime time.Time) string {
	year, mon, day := dateTime.Date()
	return fmt.Sprintf("%02d-%02d-%02d", year, mon, day)
}

// TimeToYear
func (t *itime) TimeToYear(dateTime time.Time) int {
	return dateTime.Year()
}

// TimeToMonthString
func (t *itime) TimeToMonthString(dateTime time.Time) string {
	return dateTime.Month().String()
}

// TimeToMonth
func (t *itime) TimeToMonth(dateTime time.Time) int {
	month := t.TimeToMonthString(dateTime)
	return monthMap[month]
}

// TimeToDay
func (t *itime) TimeToDay(dateTime time.Time) int {
	return dateTime.Day()
}

// NowToClock 当前时间，字符串
func (t *itime) TimeToClock(dateTime time.Time) string {
	hour, min, sec := dateTime.Clock()
	return fmt.Sprintf("%02d:%02d:%02d", hour, min, sec)
}

// TimeToHour
func (t *itime) TimeToHour(dateTime time.Time) int {
	return dateTime.Hour()
}

// TimeToMinute
func (t *itime) TimeToMinute(dateTime time.Time) int {
	return dateTime.Minute()
}

// TimeToSecond
func (t *itime) TimeToSecond(dateTime time.Time) int {
	return dateTime.Second()
}

// TimeToWeekDayString
func (t *itime) TimeToWeekDayString(dateTime time.Time) string {
	return dateTime.Weekday().String()
}

// TimeToWeekDay
func (t *itime) TimeToWeekDay(dateTime time.Time) int {
	week := t.TimeToWeekDayString(dateTime)
	return weekMap[week]
}

// NowToUnix 当前时间，时间戳
func (t *itime) NowToUnix() int64 {
	return time.Now().Unix()
}

// NowToString
func (t *itime) NowToString() string {
	return t.UnixToString(t.NowToUnix())
}

// NowToDate 当前日期，字符串
func (t *itime) NowToDate() string {
	return t.TimeToDate(time.Now())
}

// NowToYear
func (t *itime) NowToYear() int {
	return t.TimeToYear(time.Now())
}

// NowToMonthString
func (t *itime) NowToMonthString() string {
	return t.TimeToMonthString(time.Now())
}

// NowToMonth
func (t *itime) NowToMonth() int {
	return t.TimeToMonth(time.Now())
}

// NowToDay
func (t *itime) NowToDay() int {
	return t.TimeToDay(time.Now())
}

// NowToClock 当前时间，字符串
func (t *itime) NowToClock() string {
	return t.TimeToClock(time.Now())
}

// NowToHour
func (t *itime) NowToHour() int {
	return t.TimeToHour(time.Now())
}

// NowToMinute
func (t *itime) NowToMinute() int {
	return t.TimeToMinute(time.Now())
}

// NowToSecond
func (t *itime) NowToSecond() int {
	return t.TimeToSecond(time.Now())
}

// NowToWeekDayString
func (t *itime) NowToWeekDayString() string {
	return t.TimeToWeekDayString(time.Now())
}

// NowToWeekDay
func (t *itime) NowToWeekDay() int {
	return t.TimeToWeekDay(time.Now())
}
