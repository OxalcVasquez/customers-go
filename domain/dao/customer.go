package dao

type Customer struct {
	BaseModel
	ID       int    `gorm:"column:id; primary_key; not null" json:"id"`
	Name     string `gorm:"column:name; not null" json:"name"`
	LastName string `gorm:"column:last_name; not null" json:"last_name"`
	Email    string `gorm:"column:email; not null" json:"email"`
	Phone    string `gorm:"column:phone" json:"phone"`
	Status   int    `gorm:"colum:status" json:"status"`
	TypeID   int    `gorm:"colum:type_id; not null" json:"type_id"`
	Type     Type   `gorm:"foreignKey:TypeID; references:ID" json:"type"`
}
