package service

import (
	"example.com/dto"
	"example.com/model"
)

type CustomerService struct{}

func (cs CustomerService) GetAll(page int, size int) []model.Customer {
	var customers []model.Customer

	model.Db.Scopes(model.Paginate(page, size)).Find(&customers)

	return customers
}

func (cs CustomerService) FindById(id string) (model.Customer, error) {
	var customer model.Customer
	err := model.Db.Where("id = ?", id).First(&customer).Error

	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (cs CustomerService) Create(input dto.CustomerCreateInput) model.Customer {
	var customer = model.Customer{Name: input.Name, Email: input.Email}

	model.Db.Save(&customer)
	return customer
}

func (cs CustomerService) Update(id string, input dto.CustomerUpdateInput) (model.Customer, error) {
	var customer model.Customer
	err := model.Db.Where("id = ?", id).First(&customer).Error

	if err != nil {
		return customer, err
	}

	customer.Name = input.Name

	model.Db.Save(&customer)

	return customer, nil

}

func (cs CustomerService) Delete(id string) error {
	var customer model.Customer
	err := model.Db.Where("id = ?", id).First(&customer).Error

	if err != nil {
		return err
	}

	model.Db.Delete(&customer)

	return nil
}
