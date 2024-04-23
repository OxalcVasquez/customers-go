package dao

import "gorm.io/gorm"

type Type struct {
	gorm.Model
	Type string `gorm:"colum:type; not null" json:"type"`
}
