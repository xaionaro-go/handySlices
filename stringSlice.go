package handySlices

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type StringSlice []string

func (t *StringSlice) Scan(src interface{}) (err error) {
	var srcB []byte

	switch srcTyped := src.(type) {
	case string:
		srcB = []byte(srcTyped)
	case []uint8:
		srcB = []byte(srcTyped)
	default:
		err = fmt.Errorf("don't know how to covert %T (\"%v\") to handySlices.StringSlice", src, src)
		return
	}

	return json.Unmarshal(srcB, &t)
}
func (t StringSlice) Value() (driver.Value, error) {
	return json.Marshal(t)
}

