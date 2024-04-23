package repository

import (
	"customers-api/domain/dao"

	"gorm.io/gorm"
)

type TypeRepository interface {
	FindAllTypes()
}

type TypeRepositoryImpl struct {
	db *gorm.DB
}

func TypeRepositoryInit(db *gorm.DB) *TypeRepositoryImpl {
	db.AutoMigrate(&dao.Type{}, &dao.Customer{})
	return &TypeRepositoryImpl{
		db: db,
	}
}

func (r TypeRepositoryImpl) FindAllTypes() {

}
