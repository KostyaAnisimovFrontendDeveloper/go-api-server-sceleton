package dto

type RequestPageDTO struct {
	ID string `uri:"id" binding:"required,uuid"`
}
