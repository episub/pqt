package pqt

import (
	"database/sql/driver"
	"log"
	"reflect"

	"cloud.google.com/go/civil"
)

type NullDate struct {
	Date  civil.Date
	Valid bool
}

// Scan implements the Scanner interface.
func (nd *NullDate) Scan(value interface{}) error {
	if value == nil {
		nd.Valid = false
		return nil
	}
	log.Printf("NullDate type: %s", reflect.TypeOf(value))
	//nd.Time, nd.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (nd NullDate) Value() (driver.Value, error) {
	if !nd.Valid {
		return nil, nil
	}
	return nd.Date.String(), nil
}
