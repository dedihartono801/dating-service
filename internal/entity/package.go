package entity

type PackageResponse struct {
	Id    int    `json:"id" db:"id" validate:"required"`
	Name  string `json:"name" db:"name" validate:"required"`
	Price int    `json:"price" db:"price" validate:"required"`
}
