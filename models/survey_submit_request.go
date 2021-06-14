package models

type SurveySubmitRequest struct {
	AppKey    string            `json:"appKey" binding:"required"`
	Answer    map[string]string `json:"answers" binding:"required"`
	ClientKey string            `json:"clientKey" binding:"required"`
	AuthToken string            `json:"authToken" binding:"required"`
}
