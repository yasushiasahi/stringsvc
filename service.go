package main

import (
	"errors"
	"strings"
)

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type stringService struct{}

// Uppercase ...
func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

// Count ...
func (stringService) Count(s string) int {
	return len(s)
}

// ErrEmpty -
var ErrEmpty = errors.New("empty string")
