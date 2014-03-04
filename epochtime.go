package vk

import (
	"strconv"
	"time"
)

// EpochTime is time in seconds. Use it to successfully parse with json
type EpochTime time.Time

// MarshalJSON unix timestamp strings
func (t EpochTime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

// UnmarshalJSON unix timestamp strings
func (t *EpochTime) UnmarshalJSON(s []byte) error {
	var err error
	var q int64

	if q, err = strconv.ParseInt(string(s), 10, 64); err != nil {
		return err
	}

	*(*time.Time)(t) = time.Unix(q, 0)
	return err
}
