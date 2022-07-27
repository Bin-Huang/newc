// Code generated by github.com/Bin-Huang/make-constructor; DO NOT EDIT.

package services

import (
	"log"

	"github.com/Bin-Huang/make-constructor/test/repositories"
)

// NewBaseService Create a new baseService
func NewBaseService(debugLevel int) *baseService {
	return &baseService{
		debugLevel: debugLevel,
	}
}

// NewEmailService Create a new EmailService
func NewEmailService(baseService baseService, userRepository *repositories.UserRepository, logger *log.Logger) *EmailService {
	return &EmailService{
		baseService:    baseService,
		userRepository: userRepository,
		logger:         logger,
	}
}

// NewUserService Create a new UserService
func NewUserService(baseService baseService, userRepository *repositories.UserRepository, proRepository *repositories.ProRepository, emailService *EmailService, logger *log.Logger, debug bool) *UserService {
	return &UserService{
		baseService:    baseService,
		userRepository: userRepository,
		proRepository:  proRepository,
		emailService:   emailService,
		logger:         logger,
		debug:          debug,
	}
}