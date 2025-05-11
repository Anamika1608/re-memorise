package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Voice struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	FilePath  string    `gorm:"not null" json:"file_path"`
	Duration  int       `json:"duration"`// sec
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (v *Voice) BeforeCreate(tx *gorm.DB) error {
	v.ID = uuid.New()
	return nil
}
