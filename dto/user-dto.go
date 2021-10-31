package dto

// UserUpdateDTO is used by client when PUT update profile
type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Name     uint64 `json:"name" form:"name" binding:"required"`
	Email    uint64 `json:"email" form:"email" binding:"required,email"`
	Password uint64 `json:"password,omitempty" form:"password,omitempty" validate:"min=6"`
}

// UserCreateDTO is used when register a user
// type UserCreateDTO struct {
// 	ID       uint64 `json:"id" form:"id" binding:"required"`
// 	Name     uint64 `json:"name" form:"name" binding:"required"`
// 	Email    uint64 `json:"email" form:"email" binding:"required" validate:"email"`
// 	Password uint64 `json:"password,omitempty" form:"password,omitempty" validate:"min:6" binding:"reqired"`
// }
