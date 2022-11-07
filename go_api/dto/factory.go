package dto

import (
	"example.com/model"
)

func CreateCustomerReponse(c model.Customer) CustomerResponse {
	return CustomerResponse{Id: c.ID, Name: c.Name, Email: c.Email}
}

func CreateListCustomerResponse(cArray []model.Customer) []CustomerResponse {
	rsArray := []CustomerResponse{}
	for _, c := range cArray {
		rsArray = append(rsArray, CreateCustomerReponse(c))
	}

	return rsArray
}
