package controller

import (
	"github.com/gin-gonic/gin"
)

type Customer struct {
	Router string `json:"Router"`
}

func (a Customer) GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"messages": "GetAll",
	})
}

func (a Customer) GetById(c *gin.Context) {
	c.JSON(200, gin.H{
		"messages": "GetById",
	})
}

func (a Customer) Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"messages": "Create",
	})
}

func (a Customer) Update(c *gin.Context) {
	c.JSON(200, gin.H{
		"messages": "Update",
	})
}

func (a Customer) Delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"messages": "Delete",
	})
}

func (a Customer) Query(c *gin.Context) {

}

func NewCustomerCtl() Customer {
	return Customer{Router: "customer"}
}