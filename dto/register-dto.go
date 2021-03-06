package dto

// RegisterDTO is used when client post form /register url
type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"min=1"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:6"`
}
