package service

import (
	"credit-report-service-backend-2/models"
	"credit-report-service-backend-2/repository"
	"credit-report-service-backend-2/utils"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type RegistrationService interface {
	CreateUser(request models.CreateUserRequest) (*models.CreateUserResponse, error)
}

type registrationService struct {
	repository repository.RegistrationRepository
}

func (service registrationService) CreateUser(request models.CreateUserRequest) (*models.CreateUserResponse, error) {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Println("unable to create uuid")
		return nil, errors.New("unable to create uuid")
	}
	clientKey := utils.RandStringBytes(27)
	authToken := strings.ToUpper(string(out))

	fmt.Print("this is auth token", authToken)
	err = service.repository.CreateUser(request, clientKey, authToken)
	if err != nil {
		log.Printf("error while creating user %v", err)
		return nil, errors.New("something went wrong")
	}

	response := models.CreateUserResponse{
		ClientKey: clientKey,
		UserId:    clientKey,
		AuthToken: authToken,
	}

	return &response, nil
}

func NewRegistrationService(registrationRepository repository.RegistrationRepository) RegistrationService {
	return registrationService{repository: registrationRepository}
}
