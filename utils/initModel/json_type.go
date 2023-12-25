package initModel

import (
	"bytes"
	"database/sql/driver"
	"errors"

	jsoniter "github.com/json-iterator/go"
)

func (j Json) Value() (driver.Value, error) {
	if j.IsNull() {
		return nil, nil
	}
	return string(j), nil
}

func (j *Json) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("invalid Scan Source")
	}
	*j = append((*j)[:], s...)
	return nil
}

func (m Json) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

func (m *Json) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("null point exception")
	}
	*m = append((*m)[:], data...)
	return nil
}

func (j Json) IsNull() bool {
	return len(j) == 0 || string(j) == "null"
}

func (j Json) Equals(j1 Json) bool {
	return bytes.Equal([]byte(j), []byte(j1))
}

// StructToJSON 写入
func StructToJSON(value interface{}) Json {
	bts, _ := jsoniter.Marshal(value)
	return bts
}
