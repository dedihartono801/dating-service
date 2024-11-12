package entity

type PaymentMethodResponse struct {
	Id   int    `json:"id" db:"id" validate:"required"`
	Name string `json:"name" db:"name" validate:"required"`
}
