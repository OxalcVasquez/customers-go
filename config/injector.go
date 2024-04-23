// go:build wireinject
//go:build wireinject
// +build wireinject

package config

import (
	"customers-api/controller"
	"customers-api/repository"
	"customers-api/service"

	"github.com/google/wire"
)

var db = wire.NewSet(ConnectToDB)

var customerServiceSet = wire.NewSet(service.CustomerServiceInit,
	wire.Bind(new(service.CustomerService), new(*service.CustomerServiceImpl)),
)

var customerRepoSet = wire.NewSet(repository.CustomerRepositoryInit,
	wire.Bind(new(repository.CustomerRepository), new(*repository.CustomerRepositoryImpl)),
)

var customerCtrlSet = wire.NewSet(controller.CustomerControllerInit,
	wire.Bind(new(controller.CustomerController), new(*controller.CustomerControlerImpl)),
)

var typeRepoSet = wire.NewSet(repository.TypeRepositoryInit,
	wire.Bind(new(repository.TypeRepository), new(*repository.TypeRepositoryImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db, customerCtrlSet, customerServiceSet, customerRepoSet, typeRepoSet)
	return nil
}
