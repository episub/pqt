package pqt

import (
	"database/sql/driver"
	"fmt"
	"log"
	"reflect"
	"time"

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

	ts := reflect.TypeOf(value).String()
	if ts != "time.Time" {
		return fmt.Errorf("Unexpected type for NullDate of %s", ts)
	}

	v, _ := value.(time.Time)
	nd.Date = civil.DateOf(v)
	nd.Valid = true

	log.Printf("Time: %s", v.Format(time.RFC3339))

	return nil
}

// Value implements the driver Valuer interface.
func (nd NullDate) Value() (driver.Value, error) {
	if !nd.Valid {
		return nil, nil
	}
	return nd.Date.String(), nil
}
