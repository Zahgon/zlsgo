package ztime

import (
	"time"
)

const (
	timePattern = `(\d{4}[-/\.]\d{1,2}[-/\.]\d{1,2})[:\sT-]*(\d{0,2}:{0,1}\d{0,2}:{0,1}\d{0,2}){0,1}\.{0,1}(\d{0,9})([\sZ]{0,1})([\+-]{0,1})([:\d]*)`
)

var (
	TimeTpl      = "2006-01-02 15:04:05"
	formatKeyTpl = map[byte]string{
		'd': "02",
		'D': "Mon",
		'w': "Monday",
		'N': "Monday",
		'S': "02",
		'l': "Monday",
		'F': "January",
		'm': "01",
		'M': "Jan",
		'n': "1",
		'Y': "2006",
		'y': "06",
		'a': "pm",
		'A': "PM",
		'g': "3",
		'h': "03",
		'H': "15",
		'i': "04",
		's': "05",
		'O': "-0700",
		'P': "-07:00",
		'T': "MST",
		'u': "000000",
		'c': "2006-01-02T15:04:05-07:00",
		'r': "Mon, 02 Jan 06 15:04 MST",
	}
	GetLocationName = func(zone int) string {
		switch zone {
		case 8:
			return "Asia/Shanghai"
		}
		return "UTC"
	}
)

type TimeEngine struct {
	zone *time.Location
}

// Zone eastEightTimeZone
func Zone(zone ...int) *time.Location { _ = "STUB: not implemented"; return nil }

// FormatTlp format template
func FormatTlp(format string) string { _ = "STUB: not implemented"; return "" }

// New new timeEngine
func New(zone ...int) *TimeEngine { _ = "STUB: not implemented"; return nil }

// SetTimeZone SetTimeZone
func (e *TimeEngine) SetTimeZone(zone int) *TimeEngine { _ = "STUB: not implemented"; return nil }

// GetTimeZone GetTimeZone
func (e *TimeEngine) GetTimeZone() *time.Location { _ = "STUB: not implemented"; return nil }

func (e *TimeEngine) In(t time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// FormatTime string format of return time
func (e *TimeEngine) FormatTime(t time.Time, format ...string) string {
	_ = "STUB: not implemented"
	return ""
}

// FormatTimestamp convert UNIX time to time string
func (e *TimeEngine) FormatTimestamp(timestamp int64, format ...string) string {
	_ = "STUB: not implemented"
	return ""
}

// Unix int to time
func (e *TimeEngine) Unix(tt int64) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// UnixMicro int to time
func (e *TimeEngine) UnixMicro(tt int64) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// Parse Parse
func (e *TimeEngine) Parse(str string, format ...string) (t time.Time, err error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func (e *TimeEngine) Week(t time.Time) int { _ = "STUB: not implemented"; return 0 }

// case "Sunday":

// MonthRange gets the start and end UNIX times for the specified year and month
func (e *TimeEngine) MonthRange(year int, month int) (beginTime, endTime int64, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}
