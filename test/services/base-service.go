package services

//go:generate go run ../../../make-constructor
type baseService struct {
	debugLevel int
}
