package controller

import (
	"login-service-backend-1/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type LoginController interface {
	Login(ctx *gin.Context)
}

type loginController struct{}

func NewLoginController() *loginController {
	return &loginController{}
}

func (ctrl loginController) Login(ctx *gin.Context) {
	var loginRequest models.LoginRequest
	err := ctx.ShouldBindBodyWith(&loginRequest, binding.JSON)
	if err != nil {
		sendMessageWithStatus(ctx, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	sendMessageWithStatus(ctx, http.StatusOK, "valid request")
}

func sendMessageWithStatus(ctx *gin.Context, httpStatusCode int, message string) {
	responseMap := struct{ Message string }{Message: message}

	ctx.JSON(httpStatusCode,
		responseMap,
	)
}
