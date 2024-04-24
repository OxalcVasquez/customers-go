package repository

import (
	"customers-api/domain/dao"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TypeRepository interface {
	FindAllTypes() ([]dao.Type, error)
}

type TypeRepositoryImpl struct {
	db *gorm.DB
}

func TypeRepositoryInit(db *gorm.DB) *TypeRepositoryImpl {
	db.AutoMigrate(&dao.Type{}, &dao.Type{})
	return &TypeRepositoryImpl{
		db: db,
	}
}

func (r TypeRepositoryImpl) FindAllTypes() ([]dao.Type, error) {
	var types []dao.Type

	var err = r.db.Find(&types).Error

	if err != nil {
		log.Error("Error getting all types. Error ", err)
		return nil, err
	}

	return types, nil

}
