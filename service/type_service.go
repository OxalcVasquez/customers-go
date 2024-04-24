package service

import (
	"customers-api/constant"
	"customers-api/pkg"
	"customers-api/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type TypeService interface {
	GetAllTypes(c *gin.Context)
}

type TypeServiceImpl struct {
	repo repository.TypeRepository
}

func TypeServiceInit(repo repository.TypeRepository) *TypeServiceImpl {
	return &TypeServiceImpl{
		repo: repo,
	}
}

func (customerService TypeServiceImpl) GetAllTypes(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("Start to get all types")

	data, err := customerService.repo.FindAllTypes()

	if err != nil {
		log.Error("Error ocurred find all types. Error", err)
		pkg.PanicException(constant.UknownError)
	}
	//podria solo retornar la data
	c.JSON(http.StatusOK, pkg.BuildReponse(constant.Success, data))
}
