package repository

import (
	"credit-report-service-backend-2/models"
	"credit-report-service-backend-2/utils/time"
	"database/sql"
	"log"
)

type RegistrationRepository interface {
	CreateUser(request models.CreateUserRequest, userId, authToken string) error
}

type registrationRepository struct {
	db *sql.DB
}

const (
	InsertUser = "INSERT INTO Users(APP_KEY, USER_ID, FIRSTNAME, LASTNAME, DOB, SSN, STATE, STREET, CITY, ZIP, AUTH_TOKEN) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
)

func (repository registrationRepository) CreateUser(request models.CreateUserRequest, userId, authToken string) error {
	address := request.Address
	_, err := repository.db.Exec(InsertUser, request.AppKey, userId, request.FirstName, request.LastName, time.GetTimeFromString(request.DOB), request.SSN, address.State, address.State, address.City, address.Zip, authToken)

	if err != nil {
		log.Printf("unable to create user")
		return err
	}

	return nil
}

func NewRegistrationRepository(db *sql.DB) RegistrationRepository {
	return registrationRepository{
		db: db,
	}
}
