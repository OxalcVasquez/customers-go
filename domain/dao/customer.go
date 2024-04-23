package dao

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name     string `gorm:"column:name; not null" json:"name"`
	LastName string `gorm:"column:last_name; not null" json:"last_name"`
	Email    string `gorm:"column:email; not null" json:"email"`
	Phone    string `gorm:"column:phone" json:"phone"`
	Status   int    `gorm:"colum:status" json:"status"`
	TypeId   int    `gorm:"colum:type_id; not null" json:"type_id"`
	Type     Type   `gorm:"foreignKey:TypeID; references:ID" json:"type"`
}
