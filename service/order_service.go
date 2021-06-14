package service

import (
	"credit-report-service-backend-2/models"
	"credit-report-service-backend-2/repository"
	"credit-report-service-backend-2/utils"
)

type CreditOrderService interface {
	OrderCredit(request models.OrderCreditRequest) (*models.OrderCreditResponse, error)
}

type creditOrderService struct {
	repository repository.CreditOrderRepository
}

func (service creditOrderService) OrderCredit(request models.OrderCreditRequest) (*models.OrderCreditResponse, error) {
	request.ReportKey = utils.RandStringBytes(10)
	return service.repository.OrderCredit(request)
}

func NewCreditOrderService(orderRepository repository.CreditOrderRepository) CreditOrderService {
	return creditOrderService{repository: orderRepository}
}
