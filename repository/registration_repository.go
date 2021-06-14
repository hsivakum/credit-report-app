package repository

import (
	"database/sql"
	"log"
)

type RegistrationRepository interface {
	CreateUser(email, password string) error
}

type registrationRepository struct {
	db *sql.DB
}

const (
	InsertUser = "INSERT INTO Users(EMAIL, PASSWORD) VALUES (?, SHA1(?))"
)

func (repository registrationRepository) CreateUser(email, password string) error {
	_, err := repository.db.Exec(InsertUser, email, password)

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
