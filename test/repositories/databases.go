package repositories

import "time"

//go:generate go run ../../../make-constructor
type database struct {
	DSN     string
	Timeout time.Duration
	debug   bool
}
