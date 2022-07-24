package repositories

//go:generate go run ../../../make-constructor
type ProRepository struct {
	db        *database
	TableName string
	version   int
}
