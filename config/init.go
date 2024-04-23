package config

import "customers-api/repository"

type Initialization struct {
	customerRepo       repository.CustomerRepository
	customerService    repository.CustomerRepository
	customerController repository.CustomerRepository
	typeRepo           repository.TypeRepository
}

func NewInitialization(customerRepo repository.CustomerRepository, customerService repository.CustomerRepository,
	customerController repository.CustomerRepository, typeRepo repository.TypeRepository) *Initialization {
	return &Initialization{
		customerRepo:       customerRepo,
		customerService:    customerService,
		customerController: customerController,
		typeRepo:           typeRepo,
	}
}
