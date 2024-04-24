package dto

type CreateCustomerRequest struct {
	Name     string `validate:"required,min=1,max=200" json:"name"`
	LastName string `validate:"required,min=1,max=200" json:"last_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Status   bool   `json:"status"`
	TypeId   int    `json:"type_id"`
}
