package repository

import (
	"credit-report-service-backend-2/models"
	"github.com/jmoiron/sqlx"
	"log"
)

type QuestionsRepository interface {
	FetchQuestions() ([]models.Question, error)
}

type questionsRepository struct {
	db *sqlx.DB
}

const (
	SelectQuestions = "SELECT Q.ID, Q.Question, JSON_ARRAYAGG(JSON_OBJECT('id', A.ID, 'text', A.`Option`)) as `Answers` from Answers A, Questions Q WHERE A.QuestionID = Q.ID GROUP BY Q.ID, Q.Question"
)

func (repository questionsRepository) FetchQuestions() ([]models.Question, error) {
	var questions []models.Question
	err := repository.db.Select(&questions, SelectQuestions)

	if err != nil {
		log.Printf("unable to fetch questions")
		return nil, err
	}

	return questions, nil
}

func NewQuestionsRepository(db *sqlx.DB) QuestionsRepository {
	return questionsRepository{db: db}
}
