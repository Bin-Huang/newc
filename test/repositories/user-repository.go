package repositories

import (
	"errors"

	config2 "github.com/Bin-Huang/make-constructor/test/config2"
)

// UserRepository the user repository for example
//go:generate go run ../../../make-constructor --init
type UserRepository struct {
	conf      *config2.Config
	db        *database
	TableName string
}

func (r *UserRepository) init() {
	r.TableName = "foo"
}

// FindByID find something by id
func (r *UserRepository) FindByID() error {
	return errors.New("no found")
}
