package main

import (
	"errors"
	"fmt"
)

type Status int

const (
	TODO Status = iota
	IN_PROGRESS
	DONE
	UNKNOWN
)

func ParseFromString(s string) (Status, error) {
	switch s {
	case "todo":
		return TODO, nil
	case "in-progress":
		return IN_PROGRESS, nil
	case "done":
		return DONE, nil
	default:
		return UNKNOWN, errors.New(fmt.Sprintf("cannot parse %s to Status type", s))
	}
}

func (s Status) ToString() string {
	switch s {
	case TODO:
		return "todo"
	case IN_PROGRESS:
		return "in progress"
	case DONE:
		return "done"
	default:
		return ""
	}
}

func CheckValidStatus(s string) bool {
	_, err := ParseFromString(s)
	return err == nil
}
