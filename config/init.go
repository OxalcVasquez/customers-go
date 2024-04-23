package config

import (
	"customers-api/controller"
	"customers-api/repository"
	"customers-api/service"
)

type Initialization struct {
	customerRepo       repository.CustomerRepository
	customerService    service.CustomerService
	CustomerController controller.CustomerController
	TypeRepo           repository.TypeRepository
}

func NewInitialization(customerRepo repository.CustomerRepository,
	customerService service.CustomerService,
	customerController controller.CustomerController,
	typeRepo repository.TypeRepository) *Initialization {
	return &Initialization{
		customerRepo:       customerRepo,
		customerService:    customerService,
		CustomerController: customerController,
		TypeRepo:           typeRepo,
	}
}
