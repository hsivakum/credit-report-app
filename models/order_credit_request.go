package models

type OrderCreditRequest struct {
	ClientKey   string `json:"clientKey" binding:"required"`
	ProductCode string `json:"productCode" binding:"required"`
	ReportKey   string
}
