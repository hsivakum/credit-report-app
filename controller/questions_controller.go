package controller

import (
	"credit-report-service-backend-2/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type questionsController struct {
	service service.QuestionsService
}

func (controller questionsController) GetQuestions(ctx *gin.Context) {
	response, err := controller.service.GetQuestions()

	if err != nil {
		log.Printf("unable to get questions user %v", err)
		sendMessageWithStatus(ctx, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func NewQuestionsController(questionsService service.QuestionsService) *questionsController {
	return &questionsController{service: questionsService}
}
