package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type StringArray []string

func (s *StringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan StringArray: not a byte slice")
	}
	return json.Unmarshal(bytes, s)
}

func (s StringArray) Value() (driver.Value, error) {
	return json.Marshal(s)
}
