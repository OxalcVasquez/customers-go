package service

import (
	"customers-api/constant"
	"customers-api/domain/dao"
	"customers-api/pkg"
	"customers-api/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type CustomerService interface {
	GetAllCustomer(c *gin.Context)
	GetCustomerById(c *gin.Context)
	CreateCustomer(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	DeleteCustomer(c *gin.Context)
}

type CustomerServiceImpl struct {
	repo repository.CustomerRepository
}

func CustomerServiceInit(repo repository.CustomerRepository) *CustomerServiceImpl {
	return &CustomerServiceImpl{
		repo: repo,
	}
}

func (customerService CustomerServiceImpl) GetAllCustomer(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("Start to get all customers")

	data, err := customerService.repo.FindAllCustomer()

	if err != nil {
		log.Error("Error ocurred find all customers. Error", err)
		pkg.PanicException(constant.UknownError)
	}
	//podria solo retornar la data
	c.JSON(http.StatusOK, pkg.BuildReponse(constant.Success, data))
}

func (customerService CustomerServiceImpl) GetCustomerById(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("Start to get customer by id")
	customerID, _ := strconv.Atoi(c.Param("customerID"))

	data, err := customerService.repo.FindCustomerById(customerID)

	if err != nil {
		log.Error("Error ocurred find customer id. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildReponse(constant.Success, data))
}

func (customerService CustomerServiceImpl) CreateCustomer(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("Start to execute create customer")

	var request dao.Customer
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Error when mapping request create customer. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := customerService.repo.Save(&request)
	if err != nil {
		log.Error("Error saving customer to database. Error", err)
		pkg.PanicException(constant.UknownError)
	}

	c.JSON(http.StatusCreated, pkg.BuildReponse(constant.Success, data))
}

func (customerService CustomerServiceImpl) UpdateCustomer(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("Start to execute pdate customer")

	customerID, _ := strconv.Atoi(c.Param("customerID"))

	var request dao.Customer
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Error when mapping request update customer. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := customerService.repo.FindCustomerById(customerID)
	if err != nil {
		log.Error("Error when get customer from database. Error", err)
		pkg.PanicException(constant.UknownError)
	}

	data.Name = request.Name
	data.LastName = request.LastName
	data.Email = request.Email
	data.Phone = request.Phone
	data.Name = request.Name
	data.Status = request.Status
	data.TypeId = request.TypeId
	customerService.repo.Update(customerID, &data)

	if err != nil {
		log.Error("Error updating customer. Error", err)
		pkg.PanicException(constant.UknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildReponse(constant.Success, data))
}

func (customerService CustomerServiceImpl) DeleteCustomer(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("Start to execute delete customer")

	customerID, _ := strconv.Atoi(c.Param("customerID"))

	err := customerService.repo.Delete(customerID)
	if err != nil {
		log.Error("Error when delete customer from database. Error", err)
		pkg.PanicException(constant.UknownError)
	}

	c.JSON(http.StatusNoContent, pkg.BuildReponse(constant.Success, pkg.Null()))
}
