package page

import (
	"github.com/google/uuid"
	"time"
)

// ============================== Request DTO ==========================================================================

type RequestFilterPageDto struct {
	Name          []string          `form:"names[]"`
	Limit         int               `form:"limit"`
	Cursor        string            `form:"cursor"`
	LastTimestamp string            `json:"lastTimestamp"`
	Orders        map[string]string `json:"orders[]"`
}

type RequestPageDTO struct {
	Name      string `form:"name" example:"Some Page"`
	CreatedAt string `form:"created_at" example:"2022-01-01T00:00:00Z"`
	UpdatedAt string `form:"updated_at" example:"2022-01-01T00:00:00Z"`
	DeletedAt string `form:"deleted_at" example:"2022-01-01T00:00:00Z"`
}

type RequestPageIdDTO struct {
	ID string `uri:"id" binding:"required,uuid" example:"987fbc97-4bed-5078-9f07-9141ba07c9f3"`
}

// ============================== Response DTO =========================================================================

type ErrorResponseDto struct {
	Message string `json:"message"`
}

type SuccessResponseDto struct {
	Message string `json:"message"`
}

type PageItemResultDto struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type ResultListDTO struct {
	List          []PageItemResultDto `json:"list"`
	Cursor        uuid.UUID           `json:"cursor"`
	LastTimestamp time.Time           `json:"lastTimestamp"`
	Total         int64               `json:"total"`
}
