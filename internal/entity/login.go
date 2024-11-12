package entity

type LoginResponse struct {
	Token     string `json:"token"`
	ExpiredAt string `json:"expired_at"`
}

type GetUserDetailResponse struct {
	Id         string `json:"id" db:"id"`
	Email      string `json:"email" db:"email"`
	Password   string `json:"password" db:"password"`
	Gender     string `json:"gender" db:"gender"`
	IsVerified bool   `json:"is_verified" db:"is_verified"`
	IsPremium  bool   `json:"is_premium" db:"is_premium"`
}

type GetUserListResponse struct {
	Id             string  `json:"id" db:"id"`
	FirstName      string  `json:"first_name" db:"first_name"`
	LastName       string  `json:"last_name" db:"last_name"`
	ProfilePicture *string `json:"profile_picture" db:"profile_picture"`
	Bio            string  `json:"bio" db:"bio"`
	Location       string  `json:"location" db:"location"`
	Age            int     `json:"age" db:"age"`
	Gender         string  `json:"gender" db:"gender"`
	IsVerified     bool    `json:"is_verified" db:"is_verified"`
	IsPremium      bool    `json:"is_premium" db:"is_premium"`
}
