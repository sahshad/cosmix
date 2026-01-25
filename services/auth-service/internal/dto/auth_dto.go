package dto

type RegisterDTO struct {
	FirstName   string     `json:"first_name" binding:"required"`
	LastName    string     `json:"last_name" binding:"required"`
	Username    string     `json:"user_name" binding:"required,alphanum"`
	Email       string     `json:"email" binding:"required,email"`
	Password    string     `json:"password" binding:"required,min=6"`
}

type LoginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
