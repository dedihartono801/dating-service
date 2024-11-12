package entity

type SignupRequest struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Gender      string `json:"gender" validate:"required"`
	Age         int    `json:"age" validate:"required"`
	DateOfBirth string `json:"date_of_birth" validate:"required"`
	Bio         string `json:"bio"`
	Location    string `json:"location"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SwipeRequest struct {
	TargetUserId int    `json:"target_user_id" validate:"required"`
	SwipeType    string `json:"swipe_type" validate:"required"`
}
