package result

import "github.com/google/uuid"

type PageItemResultDto struct {
	ID   uuid.UUID `json:"id" example:"2a4ced49-0f43-496e-b823-5af77407fd2c"`
	Name string    `json:"name" example:"Some Page"`
}
