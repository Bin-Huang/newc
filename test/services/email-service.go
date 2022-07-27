package services

import (
	"log"

	"github.com/Bin-Huang/make-constructor/test/repositories"
)

// EmailService email service for example
//go:generate go run github.com/Bin-Huang/make-constructor@v0.5.0
type EmailService struct {
	baseService
	userRepository *repositories.UserRepository
	logger         *log.Logger
}
