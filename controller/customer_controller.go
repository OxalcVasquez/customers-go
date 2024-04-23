package controller

import (
	// _ "customers-api/domain/dto/"
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

// GetAllCustomer godoc
//
//	@Summary      Lista clientes
//	@Description  get all customers
//	@Tags         customers
//	@Accept       json
//	@Produce      application/json
//	@Success      200  {array}   dto.ApiResponse
//	@Failure      400
//	@Failure      500
//	@Router       /customers [get]
func (customerController CustomerControlerImpl) GetAllCustomer(c *gin.Context) {
	customerController.service.GetAllCustomer(c)
}

// GetCustomerById godoc
//
//	@Summary      Obtiene un cliente
//	@Description  Lista un cliente por su id
//	@Tags         customers
//	@Accept       json
//	@Produce      application/json
//
// @Param   ID     path    int     true        "ID del cliente"
//
//	@Success      200  {object}   dto.ApiResponse
//	@Failure      400
//	@Failure      500
//	@Router       /customers/{ID} [get]
func (customerController CustomerControlerImpl) GetCustomerById(c *gin.Context) {
	customerController.service.GetCustomerById(c)
}

// CreateCustomer godoc
//
//	@Summary      Add a client
//	@Description  Add a client :
//	@Tags         customers
//	@Accept       application/json
//	@Produce      json
//
// @Param   customer body dto.CreateCustomerRequest true "Create customer"
//
//	@Success      201  {object}   dto.ApiResponse
//	@Failure      400
//	@Failure      500
//	@Router       /customers/ [post]
func (customerController CustomerControlerImpl) CreateCustomer(c *gin.Context) {
	customerController.service.CreateCustomer(c)
}

// UpdateCustomer godoc
//
//	@Summary      Update a customer
//	@Description  Update a customer
//	@Tags         customers
//	@Accept       application/json
//	@Produce      json
//
// @Param   customer body dto.UpdateCustomerRequest true "Update customer"
//
//	@Success      200  {object}   dto.ApiResponse
//	@Failure      400
//	@Failure      500
//	@Router       /customers/ [put]
func (customerController CustomerControlerImpl) UpdateCustomer(c *gin.Context) {
	customerController.service.UpdateCustomer(c)
}

// DeleteCustomer godoc
//
//	@Summary      Elimina un cliente
//	@Description  Elimina un cliente por su id
//	@Tags         customers
//	@Accept       json
//	@Produce      application/json
//
// @Param   ID     path    int     true        "ID del cliente"
//
//	@Success      204  {object}   dto.ApiResponse
//	@Failure      400
//	@Failure      500
//	@Router       /customers/{ID} [delete]
func (customerController CustomerControlerImpl) DeleteCustomer(c *gin.Context) {
	customerController.service.DeleteCustomer(c)
}
