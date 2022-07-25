package repositories

import "time"

// database the internal database client for example
//go:generate go run github.com/Bin-Huang/make-constructor@v0.3.0
type database struct {
	DSN     string
	Timeout time.Duration
	debug   bool
}
