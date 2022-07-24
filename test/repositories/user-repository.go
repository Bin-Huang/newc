package repositories

//go:generate go run ../../../make-constructor
type UserRepository struct {
	db        *database
	TableName string
}

func (r *UserRepository) FindByID() {
}
