package pqt

import (
	"database/sql/driver"
	"log"
	"reflect"
)

type NullBytes struct {
	Bytes []byte
	Valid bool
}

// Scan implements the Scanner interface.
func (n *NullBytes) Scan(value interface{}) error {
	log.Printf("NullBytes type: %s", reflect.TypeOf(value))
	panic("NullBytes scan not implemented")
	//n.Time, n.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (n NullBytes) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Bytes, nil
}
