package dao

type Type struct {
	BaseModel
	ID   int    `gorm:"column:id; primary_key; not null" json:"id"`
	Type string `gorm:"colum:type; not null" json:"type"`
}
