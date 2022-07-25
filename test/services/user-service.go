package services

import (
	"log"

	"github.com/Bin-Huang/make-constructor/test/repositories"
)

// UserService a user service for example
//go:generate go run github.com/Bin-Huang/make-constructor@v0.2.0
type UserService struct {
	baseService

	userRepository *repositories.UserRepository
	proRepository  *repositories.ProRepository

	emailService *EmailService

	logger *log.Logger

	debug bool
}
