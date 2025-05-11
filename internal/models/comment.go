package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Content   string    `gorm:"not null" json:"content"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"-"`
	MemoryID  uuid.UUID `gorm:"type:uuid;not null" json:"memory_id"`
	Memory    Memory    `gorm:"foreignKey:MemoryID" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New()
	return nil
}
