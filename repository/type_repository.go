package repository

import (
	"customers-api/domain/dao"

	"gorm.io/gorm"
)

type RoleRepository interface {
	FindAllRole()
}

type TypeRepositoryImpl struct {
	db *gorm.DB
}

func RoleRepositoryInit(db *gorm.DB) *TypeRepositoryImpl {
	db.AutoMigrate(&dao.Type{}, &dao.Customer{})
	return &TypeRepositoryImpl{
		db: db,
	}
}

func (r TypeRepositoryImpl) FindAllTypes() {

}
