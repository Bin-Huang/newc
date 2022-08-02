package repositories

// UserRepository the user repository for example
//go:generate go run ../../../make-constructor
type UserRepository struct {
	db        *database
	TableName string
}

// FindByID find something by id
func (r *UserRepository) FindByID() {
}
