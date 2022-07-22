package pkg1

// F ...
func F(a interface{}) {
}

// Service ...
//go:generate go run ../../../make-constructor
type Service struct {
	Name string
}

// PostService ...
//go:generate go run ../../../make-constructor
type PostService struct {
	Service
	Version int
}

// AgeService ...
//go:generate go run ../../../make-constructor
type AgeService struct {
	Service
	Age int
}
