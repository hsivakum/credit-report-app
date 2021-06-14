package controller

import (
	"credit-report-service-backend-2/models"
	"credit-report-service-backend-2/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
)

type orderController struct {
	service service.CreditOrderService
}

func(controller orderController) OrderCredit(ctx *gin.Context) {
	var orderCreditRequest models.OrderCreditRequest
	err := ctx.ShouldBindBodyWith(&orderCreditRequest, binding.JSON)
	if err != nil {
		sendMessageWithStatus(ctx, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	response, err := controller.service.OrderCredit(orderCreditRequest)
	if err != nil{
		log.Println("unable ot save survey results")
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func NewOrderController(orderService service.CreditOrderService) *orderController {
	return &orderController{service: orderService}
}
