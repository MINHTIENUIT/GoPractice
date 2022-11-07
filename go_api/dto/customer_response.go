package dto

type CustomerResponse struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
