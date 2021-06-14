package service

import (
	"credit-report-service-backend-2/repository"
	"errors"
	"log"
)

type RegistrationService interface {
	CreateUser(email, password string) error
}

type registrationService struct {
	repository repository.RegistrationRepository
}

func (service registrationService) CreateUser(email, password string) error {
	err := service.repository.CreateUser(email, password)
	if err != nil {
		log.Printf("error while creating user %v", err)
		return errors.New("something went wrong")
	}

	return nil
}

func NewRegistrationService(registrationRepository repository.RegistrationRepository) RegistrationService {
	return registrationService{repository: registrationRepository}
}
