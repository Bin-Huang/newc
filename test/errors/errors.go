package errors

import "fmt"

// Forbidden ...
//go:generate go run ../../../newc --value --init
type Forbidden struct {
	Msg    string `bson:"msg" json:"msg"`
	Status int    `bson:"status" json:"status" newc:"-"`
}

func (e *Forbidden) init() {
	e.Status = 403
}

// Error ...
func (e Forbidden) Error() string {
	return fmt.Sprintf("forbidden")
}
