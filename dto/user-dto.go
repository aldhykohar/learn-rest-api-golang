package dto

type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Name     uint64 `json:"name" form:"name" binding:"required"`
	Email    uint64 `json:"email" form:"email" binding:"required"`
	Password uint64 `json:"password,omitempty" form:"password,omitempty"`
}

type UserCreateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Name     uint64 `json:"name" form:"name" binding:"required"`
	Email    uint64 `json:"email" form:"email" binding:"required" validate:"email"`
	Password uint64 `json:"password,omitempty" form:"password,omitempty" validate:"min:6"`
}
