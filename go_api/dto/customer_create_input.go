package dto

type CustomerCreateInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}
