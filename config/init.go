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
	typeRepo           repository.TypeRepository
	typeService        service.TypeService
	TypeController     controller.TypeController
}

func NewInitialization(customerRepo repository.CustomerRepository,
	customerService service.CustomerService,
	customerController controller.CustomerController,
	typeRepo repository.TypeRepository, typeService service.TypeService, typeController controller.TypeController) *Initialization {
	return &Initialization{
		customerRepo:       customerRepo,
		customerService:    customerService,
		CustomerController: customerController,
		typeRepo:           typeRepo,
		typeService:        typeService,
		TypeController:     typeController,
	}
}
