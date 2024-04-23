package controller

import (
	"customers-api/service"

	"github.com/gin-gonic/gin"
)

type CustomerController interface {
	GetAllCustomer(c *gin.Context)
	GetCustomerById(c *gin.Context)
	CreateCustomer(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	DeleteCustomer(c *gin.Context)
}

type CustomerControlerImpl struct {
	service service.CustomerService
}

func CustomerControllerInit(customerService service.CustomerService) *CustomerControlerImpl {
	return &CustomerControlerImpl{
		service: customerService,
	}
}

func (customerController CustomerControlerImpl) GetAllCustomer(c *gin.Context) {
	customerController.service.GetAllCustomer(c)
}

func (customerController CustomerControlerImpl) GetCustomerById(c *gin.Context) {
	customerController.service.GetCustomerById(c)
}

func (customerController CustomerControlerImpl) CreateCustomer(c *gin.Context) {
	customerController.service.CreateCustomer(c)
}

func (customerController CustomerControlerImpl) UpdateCustomer(c *gin.Context) {
	customerController.service.UpdateCustomer(c)
}

func (customerController CustomerControlerImpl) DeleteCustomer(c *gin.Context) {
	customerController.service.DeleteCustomer(c)
}
