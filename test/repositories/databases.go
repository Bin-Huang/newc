package repositories

import "time"

//go:generate go run github.com/Bin-Huang/make-constructor@v0.1.0
type database struct {
	DSN     string
	Timeout time.Duration
	debug   bool
}
