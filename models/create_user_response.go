package models

type CreateUserResponse struct {
	ClientKey string `json:"clientKey"`
	UserId    string `json:"userId"`
	AuthToken string `json:"authToken"`
}
