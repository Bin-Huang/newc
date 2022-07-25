package repositories

//go:generate go run github.com/Bin-Huang/make-constructor@v0.1.0
type UserRepository struct {
	db        *database
	TableName string
}

func (r *UserRepository) FindByID() {
}
