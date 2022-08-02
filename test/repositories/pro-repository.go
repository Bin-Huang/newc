package repositories

import (
	"github.com/Bin-Huang/make-constructor/test/config"
	"github.com/Bin-Huang/make-constructor/test/pkg/errors"
)

// ProRepository a repository for example
//go:generate go run ../../../make-constructor
type ProRepository struct {
	conf      config.Config
	db        *database
	TableName string
	version   int
}

// FindByID find something by id
func (r *ProRepository) FindByID() errors.NoFound {
	return errors.NoFound{}
}
