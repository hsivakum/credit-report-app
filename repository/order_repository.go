package repository

import (
	"credit-report-service-backend-2/models"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type CreditOrderRepository interface {
	OrderCredit(request models.OrderCreditRequest) (*models.OrderCreditResponse, error)
}

type creditOrderRepository struct {
	db *sqlx.DB
}

const (
	InsertOrder = "INSERT INTO CreditOrders(UserID, ProductCode, ReportKey) VALUES ((SELECT ID FROM Users WHERE USER_ID = ?), ?, ?)"
)

func (repository creditOrderRepository) OrderCredit(request models.OrderCreditRequest) (*models.OrderCreditResponse, error) {
	result, err := repository.db.Exec(InsertOrder, request.ClientKey, request.ProductCode, request.ReportKey)
	if err != nil {
		return nil, err
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &models.OrderCreditResponse{
		DisplayToken: strconv.FormatInt(insertId, 10),
		ReportKey:    request.ReportKey,
	}, nil
}

func NewCreditOrderRepository(db *sqlx.DB) CreditOrderRepository {
	return creditOrderRepository{
		db: db,
	}
}
