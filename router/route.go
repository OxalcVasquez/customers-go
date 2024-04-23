package router

import (
	"customers-api/config"

	"github.com/gin-gonic/gin"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		customer := api.Group("/user")
		customer.GET("", init.CustomerController.GetAllCustomer)
		customer.POST("", init.CustomerController.CreateCustomer)
		customer.GET("/:customerID", init.CustomerController.GetCustomerById)
		customer.PUT("/:customerID", init.CustomerController.UpdateCustomer)
		customer.DELETE("/:customerID", init.CustomerController.DeleteCustomer)
	}

	return router
}
