package services

import (
	"log"

	"github.com/Bin-Huang/newc/test/repositories"
)

// EmailService email service for example
//go:generate go run ../../../newc
type EmailService struct {
	baseService
	userRepository *repositories.UserRepository
	logger         *log.Logger
}
