package models

type QuestionsResponse struct {
	AuthToken string     `json:"authToken"`
	Provider  string     `json:"provider"`
	Questions []QuestionsResponse `json:"questions"`
}

type Question struct {
	ID      string `json:"id" db:"ID"`
	Text    string `json:"text" db:"Question"`
	Answers string `json:"answers" db:"Answers"`
}

type Answer struct {
	ID   int `json:"id"`
	Text string `json:"text"`
}

type QuestionResponse struct {
	ID      string `json:"id"`
	Text    string `json:"text"`
	Answers []Answer `json:"answers"`
}
