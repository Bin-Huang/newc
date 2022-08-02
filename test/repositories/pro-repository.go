package repositories

// ProRepository a repository for example
//go:generate go run ../../../make-constructor
type ProRepository struct {
	db        *database
	TableName string
	version   int
}
