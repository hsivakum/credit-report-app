package controller

import (
	"credit-report-service-backend-2/models"
	"credit-report-service-backend-2/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

func (controller questionsController) SaveSurveyResults(ctx *gin.Context) {
	var saveRequest models.SurveySubmitRequest
	err := ctx.ShouldBindBodyWith(&saveRequest, binding.JSON)
	if err != nil {
		sendMessageWithStatus(ctx, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	err = controller.service.SubmitAnswers(saveRequest)

	if err != nil{
		log.Println("unable ot save survey results")
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	sendMessageWithStatus(ctx, http.StatusOK, "Thank you for taking the survey")
}

func NewQuestionsController(questionsService service.QuestionsService) *questionsController {
	return &questionsController{service: questionsService}
}
