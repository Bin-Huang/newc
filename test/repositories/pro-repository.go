package repositories

import "github.com/Bin-Huang/make-constructor/test/config"

// ProRepository a repository for example
//go:generate go run ../../../make-constructor
type ProRepository struct {
	conf      config.Config
	db        *database
	TableName string
	version   int
}
