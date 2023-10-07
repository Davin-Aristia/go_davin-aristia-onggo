package dto

type UserResponse struct {
	ID 		 uint `json:"id" form:"id"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
