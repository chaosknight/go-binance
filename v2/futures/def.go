package futures

import (
	// "encoding/json"
	"strconv"
	"strings"
	"time"
)

type (
	JSONFloat64 float64
	JSONInt64   int64
	JSONTime    time.Time
)

func (t JSONTime) String() string { return time.Time(t).String() }

func (t *JSONTime) UnmarshalJSON(s []byte) (err error) {
	r := strings.Replace(string(s), `"`, ``, -1)
	if r == "" {
		return
	}

	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.UnixMilli(q)
	return
}
func (t *JSONTime) MarshalJSON() ([]byte, error) {
	return time.Time(*t).MarshalJSON()
}
func (t *JSONFloat64) UnmarshalJSON(s []byte) (err error) {
	r := strings.Replace(string(s), `"`, ``, -1)
	if r == "" {
		return
	}

	q, err := strconv.ParseFloat(r, 64)
	if err != nil {
		return err
	}
	*(*float64)(t) = q
	return
}
func (t *JSONInt64) UnmarshalJSON(s []byte) (err error) {
	r := strings.Replace(string(s), `"`, ``, -1)
	if r == "" {
		return
	}

	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*int64)(t) = q
	return
}
