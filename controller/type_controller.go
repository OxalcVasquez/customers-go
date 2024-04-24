package controller

import (
	// _ "customers-api/domain/dto/"
	"customers-api/service"

	"github.com/gin-gonic/gin"
)

type TypeController interface {
	GetAllTypes(c *gin.Context)
}

type TypeControlerImpl struct {
	service service.TypeService
}

func TypeControllerInit(typeService service.TypeService) *TypeControlerImpl {
	return &TypeControlerImpl{
		service: typeService,
	}
}

// GetAllCustomer godoc
//
//	@Summary      Lista tipos de clientes
//	@Description  get all tipos de clientes
//	@Tags         types customers
//	@Accept       json
//	@Produce      application/json
//	@Success      200  {array}   dto.ApiResponse
//	@Failure      400
//	@Failure      500
//	@Router       /types [get]
func (typeController TypeControlerImpl) GetAllTypes(c *gin.Context) {
	typeController.service.GetAllTypes(c)
}
