package repository

import (
	"customers-api/domain/dao"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Interfaz
type CustomerRepository interface {
	FindAllCustomer() ([]dao.Customer, error)
	FindCustomerById(id int) (dao.Customer, error)
	Save(customer *dao.Customer) (dao.Customer, error)
	Update(id int, customer *dao.Customer) (dao.Customer, error)
	Delete(id int) error
}

// Implementacion
type CustomerRepositoryImpl struct {
	db *gorm.DB
}

// Inicializacion
func CustomerRepositoryInit(db *gorm.DB) *CustomerRepositoryImpl {
	db.AutoMigrate(&dao.Customer{})
	return &CustomerRepositoryImpl{
		db: db,
	}
}

func (customerRepo CustomerRepositoryImpl) FindAllCustomer() ([]dao.Customer, error) {
	var customers []dao.Customer

	var err = customerRepo.db.Preload("type").Find(&customers).Error
	if err != nil {
		log.Error("Error finding all customers. Error: ", err)
		return nil, err
	}

	return customers, nil
}

func (customerRepo CustomerRepositoryImpl) FindCustomerById(id int) (dao.Customer, error) {
	customer := dao.Customer{
		ID: id,
	}

	err := customerRepo.db.Preload("Type").First(&customer).Error
	if err != nil {
		log.Error("Error finding customer by id. Error", err)
		return dao.Customer{}, err
	}

	return customer, nil

}

func (customerRepo CustomerRepositoryImpl) Save(customer *dao.Customer) (dao.Customer, error) {
	var err = customerRepo.db.Save(customer).Error
	if err != nil {
		log.Error("Error saving customer by id. Error", err)
		return dao.Customer{}, err
	}
	return *customer, nil
}

func (customerRepo CustomerRepositoryImpl) Update(customer *dao.Customer) (dao.Customer, error) {
	err := customerRepo.db.Preload("Type").First(&customer).Error

	if err != nil {
		log.Error("Error finding customer to edit. Error: ", err)
		return dao.Customer{}, err
	}

	if err := customerRepo.db.Updates(customer).Error; err != nil {
		log.Error("Error updating customer. Error: ", err)
		return dao.Customer{}, err
	}
	return *customer, nil
}

func (customerRepo CustomerRepositoryImpl) DeleteCustomerById(id int) error {
	err := customerRepo.db.Delete(&dao.Customer{}, id).Error
	if err != nil {
		log.Error("Error deleting customer. Error: ", err)
		return err
	}
	return nil
}
