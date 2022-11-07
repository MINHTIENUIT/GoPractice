package controller

import (
	"net/http"
	"strconv"

	"example.com/dto"
	"example.com/service"
	"github.com/gin-gonic/gin"
)

type Customer struct {
	Router string `json:"Router"`
}

var s service.CustomerService

func (a Customer) GetAll(c *gin.Context) {

	pageS, _ := c.GetQuery("page")
	page, err := strconv.Atoi(pageS)
	if err != nil {
		page = 0
	}

	sizeS, _ := c.GetQuery("size")
	size, err := strconv.Atoi(sizeS)
	if err != nil {
		size = 10
	}

	var cArrays = s.GetAll(page, size)
	c.JSON(http.StatusOK, dto.CreateListCustomerResponse(cArrays))
}

func (a Customer) GetById(c *gin.Context) {
	var customer, err = s.FindById(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, dto.CreateCustomerReponse(customer))
}

func (a Customer) Create(c *gin.Context) {

	var input dto.CustomerCreateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "BadRequest",
			"body":    input,
		})
		return
	}

	customer := s.Create(input)

	c.JSON(http.StatusCreated, dto.CreateCustomerReponse(customer))
}

func (a Customer) Update(c *gin.Context) {
	var input dto.CustomerUpdateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "BadRequest",
			"body":    input,
		})
		return
	}

	customer, err := s.Update(c.Param("id"), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, dto.CreateCustomerReponse(customer))

}

func (a Customer) Delete(c *gin.Context) {
	err := s.Delete(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (a Customer) Query(c *gin.Context) {

}

func NewCustomerCtl() Customer {
	return Customer{Router: "customer"}
}
