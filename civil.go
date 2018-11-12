package pqt

import (
	"database/sql/driver"
	"fmt"
	"time"

	"cloud.google.com/go/civil"
)

type Date struct {
	civil.Date
}

// Scan implements the Scanner interface.
func (d *Date) Scan(value interface{}) error {
	var when time.Time
	var valid bool
	when, valid = value.(time.Time)
	if !valid {
		return fmt.Errorf("Not a date object")
	}

	d.Date = civil.DateOf(when)

	return nil
}

// Value implements the driver Valuer interface.
func (d Date) Value() (driver.Value, error) {
	return d.Date.String(), nil
}

type NullDate struct {
	Date  civil.Date
	Valid bool
}

// Scan implements the Scanner interface.
func (nd *NullDate) Scan(value interface{}) error {
	var when time.Time
	when, nd.Valid = value.(time.Time)
	if !nd.Valid {
		return nil
	}

	nd.Date = civil.DateOf(when)
	nd.Valid = true

	return nil
}

// Value implements the driver Valuer interface.
func (nd NullDate) Value() (driver.Value, error) {
	if !nd.Valid {
		return nil, nil
	}
	return nd.Date.String(), nil
}
