package services

import (
	"log"

	"github.com/Bin-Huang/make-constructor/test/repositories"
)

//go:generate go run ../../../make-constructor
type UserService struct {
	baseService

	userRepository *repositories.UserRepository
	proRepository  *repositories.ProRepository

	emailService *EmailService

	logger *log.Logger

	debug bool
}
