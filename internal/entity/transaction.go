package entity

type TransactionRequest struct {
	PaymentMethodId int    `json:"payment_method_id" validate:"required"`
	Amount          int    `json:"amount" validate:"required"`
	Currency        string `json:"currency" validate:"required"`
	PackageTypeId   int    `json:"package_type_id" validate:"required"`
}

type PackageType struct {
	Name  string `json:"name" validate:"required"`
	Price int    `json:"price" validate:"required"`
}
