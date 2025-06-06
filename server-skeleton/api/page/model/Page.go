package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Page struct {
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name string    `gorm:"type:varchar(100);not null"`
}

func (p *Page) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}
