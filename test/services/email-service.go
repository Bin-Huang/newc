package services

import (
	"log"

	"github.com/Bin-Huang/make-constructor/test/repositories"
)

//go:generate go run ../../../make-constructor
type EmailService struct {
	baseService
	userRepository *repositories.UserRepository
	logger         *log.Logger
}
