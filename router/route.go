package router

import (
	"customers-api/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" ///gin-swagger middleware
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		customer := api.Group("/customers")
		customer.GET("", init.CustomerController.GetAllCustomer)
		customer.POST("", init.CustomerController.CreateCustomer)
		customer.GET("/:customerID", init.CustomerController.GetCustomerById)
		customer.PUT("", init.CustomerController.UpdateCustomer)
		customer.DELETE("/:customerID", init.CustomerController.DeleteCustomer)
	}

	return router
}
