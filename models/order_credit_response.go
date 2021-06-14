package models

type OrderCreditResponse struct {
	DisplayToken string `json:"displayToken" db:"DisplayToken"`
	ReportKey    string `json:"reportKey" db:"ReportKey"`
}
