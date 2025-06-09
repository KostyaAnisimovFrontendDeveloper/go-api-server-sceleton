package request

type RequestPageIdDTO struct {
	ID string `uri:"id" binding:"required,uuid" example:"987fbc97-4bed-5078-9f07-9141ba07c9f3"`
}
