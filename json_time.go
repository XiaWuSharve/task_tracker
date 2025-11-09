package main

import (
	"fmt"
	"strings"
	"time"
)

type JSONTime time.Time

type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	quotedString := fmt.Sprintf("\"%s\"", t.ToString())
	return []byte(quotedString), nil
}

func (t *JSONTime) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	parsedTime, err := time.Parse(time.DateTime, s)
	if err == nil {
		*t = JSONTime(parsedTime)
	}
	return err
}

func (t *JSONTime) ToString() string {
	s := time.Time(*t).Format(time.DateTime)
	return s
}

func JSONTimeNow() JSONTime {
	return JSONTime(time.Now())
}
