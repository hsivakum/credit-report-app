package repository

import (
	"credit-report-service-backend-2/models"
	"github.com/jmoiron/sqlx"
	"log"
)

type QuestionsRepository interface {
	FetchQuestions() ([]models.Question, error)
	SaveSurveyResults(request models.SurveySubmitRequest) error
}

type questionsRepository struct {
	db *sqlx.DB
}

const (
	SelectQuestions     = "SELECT Q.ID, Q.Question, JSON_ARRAYAGG(JSON_OBJECT('id', A.ID, 'text', A.`Option`)) as `Answers` from Answers A, Questions Q WHERE A.QuestionID = Q.ID GROUP BY Q.ID, Q.Question"
	InsertSurveyResults = "INSERT INTO UserSurveys(QuestionID, AnswerID, UserID)VALUES(?,?, (SELECT ID from Users where APP_KEY = ? AND USER_ID = ? AND AUTH_TOKEN = ?))"
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

func (repository questionsRepository) SaveSurveyResults(request models.SurveySubmitRequest) error {
	for questionID, answerID := range request.Answer {
		_, err := repository.db.Exec(InsertSurveyResults, questionID, answerID, request.AppKey, request.ClientKey, request.AuthToken)
		if err != nil {
			log.Println("unable to store survey results")
			return err
		}
	}
	return nil
}

func NewQuestionsRepository(db *sqlx.DB) QuestionsRepository {
	return questionsRepository{db: db}
}
