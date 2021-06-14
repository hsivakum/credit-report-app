package service

import (
	"credit-report-service-backend-2/models"
	"credit-report-service-backend-2/repository"
	"encoding/json"
	"log"
)

type QuestionsService interface {
	GetQuestions() (*[]models.QuestionResponse, error)
	SubmitAnswers(request models.SurveySubmitRequest) error
}

type questionsService struct {
	repository repository.QuestionsRepository
}

func (service questionsService) SubmitAnswers(request models.SurveySubmitRequest) error {
	return service.repository.SaveSurveyResults(request)
}

func (service questionsService) GetQuestions() (*[]models.QuestionResponse, error) {
	questions, err := service.repository.FetchQuestions()

	if err != nil {
		log.Printf("error while fetching questions %+v", err)
		return nil, err
	}

	var questionsResponse []models.QuestionResponse
	for _, question := range questions {
		var answers []models.Answer
		err := json.Unmarshal([]byte(question.Answers), &answers)
		if err != nil {
			return nil, err
		}
		questionsResponse = append(questionsResponse, models.QuestionResponse{
			ID:      question.ID,
			Text:    question.Text,
			Answers: answers,
		})
	}

	return &questionsResponse, nil
}

func NewQuestionsService(questionsRepository repository.QuestionsRepository) QuestionsService {
	return questionsService{repository: questionsRepository}
}
