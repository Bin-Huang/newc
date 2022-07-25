package services

import (
	"log"

	"github.com/Bin-Huang/make-constructor/test/repositories"
)

//go:generate go run github.com/Bin-Huang/make-constructor@v0.1.0
type EmailService struct {
	baseService
	userRepository *repositories.UserRepository
	logger         *log.Logger
}
