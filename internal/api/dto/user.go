package dto

type UserRegisterRequest struct {
	Email    string `json:"email" binding:"required" validate:"required,email"`
	Username string `json:"username" binding:"required" validate:"required,min=2,max=32"`
	Password string `json:"password" binding:"required" validate:"required,min=8,max=32"`
}

type UserRegisterResponse struct {
	ID int64 `json:"id"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required" validate:"required,email"`
	Password string `json:"password" binding:"required" validate:"required,min=8,max=32"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type Token struct {
	Token string `json:"token"`
}
