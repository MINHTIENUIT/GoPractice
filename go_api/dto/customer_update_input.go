package dto

type CustomerUpdateInput struct {
	Name string `json:"name" binding:"required"`
}
