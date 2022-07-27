package repositories

// ProRepository a repository for example
//go:generate go run github.com/Bin-Huang/make-constructor@v0.5.0
type ProRepository struct {
	db        *database
	TableName string
	version   int
}
