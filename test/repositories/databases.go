package repositories

import "time"

// database the internal database client for example
//go:generate go run ../../../make-constructor
type database struct {
	DSN     string
	Timeout time.Duration
	debug   bool
}
