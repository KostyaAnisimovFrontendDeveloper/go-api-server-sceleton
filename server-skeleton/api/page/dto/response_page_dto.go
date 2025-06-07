package dto

import "github.com/google/uuid"

type ResponsePageDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
