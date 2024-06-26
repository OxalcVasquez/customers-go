// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package config

import (
	"customers-api/controller"
	"customers-api/repository"
	"customers-api/service"
	"github.com/google/wire"
)

// Injectors from injector.go:

func Init() *Initialization {
	gormDB := ConnectToDB()
	customerRepositoryImpl := repository.CustomerRepositoryInit(gormDB)
	customerServiceImpl := service.CustomerServiceInit(customerRepositoryImpl)
	customerControlerImpl := controller.CustomerControllerInit(customerServiceImpl)
	typeRepositoryImpl := repository.TypeRepositoryInit(gormDB)
	typeServiceImpl := service.TypeServiceInit(typeRepositoryImpl)
	typeControlerImpl := controller.TypeControllerInit(typeServiceImpl)
	initialization := NewInitialization(customerRepositoryImpl, customerServiceImpl, customerControlerImpl, typeRepositoryImpl, typeServiceImpl, typeControlerImpl)
	return initialization
}

// injector.go:

var db = wire.NewSet(ConnectToDB)

var customerServiceSet = wire.NewSet(service.CustomerServiceInit, wire.Bind(new(service.CustomerService), new(*service.CustomerServiceImpl)))

var customerRepoSet = wire.NewSet(repository.CustomerRepositoryInit, wire.Bind(new(repository.CustomerRepository), new(*repository.CustomerRepositoryImpl)))

var customerCtrlSet = wire.NewSet(controller.CustomerControllerInit, wire.Bind(new(controller.CustomerController), new(*controller.CustomerControlerImpl)))

var typeRepoSet = wire.NewSet(repository.TypeRepositoryInit, wire.Bind(new(repository.TypeRepository), new(*repository.TypeRepositoryImpl)))

var typeServiceSet = wire.NewSet(service.TypeServiceInit, wire.Bind(new(service.TypeService), new(*service.TypeServiceImpl)))

var typeControllerSet = wire.NewSet(controller.TypeControllerInit, wire.Bind(new(controller.TypeController), new(*controller.TypeControlerImpl)))
