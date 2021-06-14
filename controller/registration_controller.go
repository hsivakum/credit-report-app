package controller

import (
	"credit-report-service-backend-2/models"
	"credit-report-service-backend-2/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
)

type registrationController struct {
	service service.RegistrationService
}

func (controller registrationController) CreateUser(ctx *gin.Context) {
	var createUserRequest models.CreateUserRequest

	err := ctx.ShouldBindBodyWith(&createUserRequest, binding.JSON)
	if err != nil {
		sendMessageWithStatus(ctx, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	response, err := controller.service.CreateUser(createUserRequest)

	if err != nil {
		log.Printf("unable to create user %v", err)
		sendMessageWithStatus(ctx, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func NewRegistrationService(registrationService service.RegistrationService) *registrationController {
	return &registrationController{service: registrationService}
}
