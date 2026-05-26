// Package ztime provides time related operations
package ztime

import (
	"sync/atomic"
	"time"
)

var inlay = New()

// Now format current time
func Now(format ...string) string { _ = "STUB: not implemented"; return "" }

// Time With the time zone of the time
func Time(realTime ...bool) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// SetTimeZone SetTimeZone
func SetTimeZone(zone int) *TimeEngine { _ = "STUB: not implemented"; return nil }

// GetTimeZone getTimeZone
func GetTimeZone() *time.Location { _ = "STUB: not implemented"; return nil }

// FormatTime format time
func FormatTime(t time.Time, format ...string) string { _ = "STUB: not implemented"; return "" }

// FormatTimestamp format timestamp
func FormatTimestamp(timestamp int64, format ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func Week(t time.Time) int { _ = "STUB: not implemented"; return 0 }

func MonthRange(year int, month int) (beginTime, endTime int64, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// Parse string to time
func Parse(str string, format ...string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// Unix int to time
func Unix(tt int64) time.Time {
	_ = "STUB: not implemented"
	return *

	// UnixMicro int to time
	new(time.Time)
}

func UnixMicro(tt int64) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// In time to time
func In(tt time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

var clock int64

func init() {
	atomic.StoreInt64(&clock, time.Now().UnixNano()/1000)

	go func() {
		const updateInterval = 10 * time.Millisecond
		const microsecondsPerUpdate = int64(10000)

		ticker := time.NewTicker(updateInterval)
		defer ticker.Stop()

		for {
			atomic.StoreInt64(&clock, time.Now().UnixNano()/1000)
			for i := 0; i < 10; i++ {
				<-ticker.C
				atomic.AddInt64(&clock, microsecondsPerUpdate)
			}
			<-ticker.C
		}
	}()
}

// Clock The current microsecond timestamp has an accuracy of 100ms
func Clock() int64 { _ = "STUB: not implemented"; return 0 }
