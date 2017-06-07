package yzgo

import (
	"strings"
	"time"
)

//YZTime the youzan timestamp
//Reference: http://blog.charmes.net/2015/08/json-date-management-in-golang_7.html

type YZTime struct {
	time.Time
}

//MarshalJSON for json convert
func (t YZTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Time.Format("2006-01-02 15:04:05") + `"`), nil
}

//UnmarshalJSON for json convert
func (t *YZTime) UnmarshalJSON(buf []byte) error {
	tt, err := time.Parse("2006-01-02 15:04:05", strings.Trim(string(buf), `"`))
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}
