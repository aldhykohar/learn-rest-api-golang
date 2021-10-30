package dto

type BookUpdateDTO struct {
	ID    uint64 `json:"id" form:"id" binding:"required"`
	Title string `json:"title" form:"title" binding:"required"`
}
