package models

type CreateUserRequest struct {
	AppKey    string  `json:"appKey" binding:"required"`
	FirstName string  `json:"firstName" binding:"required"`
	LastName  string  `json:"lastName"  binding:"required"`
	SSN       string  `json:"ssn" binding:"ValidSSN"`
	DOB       string  `json:"dob" binding:"ValidDOB"`
	Address   Address `json:"address" binding:"required"`
}

type Address struct {
	State  string `json:"state" binding:"required"`
	Street string `json:"street" binding:"required"`
	City   string `json:"city" binding:"required"`
	Zip    string `json:"zip" binding:"required"`
}
