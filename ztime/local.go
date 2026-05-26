package ztime

import (
	"database/sql/driver"
	"time"
)

type LocalTime struct {
	time.Time
}

func (t LocalTime) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (t LocalTime) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

func (t LocalTime) String() string { _ = "STUB: not implemented"; return "" }

func (t LocalTime) Format(layout string) string { _ = "STUB: not implemented"; return "" }

func (t *LocalTime) Scan(v interface{}) error { _ = "STUB: not implemented"; return nil }
