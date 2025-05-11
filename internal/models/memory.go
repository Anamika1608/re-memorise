package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Memory struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key" json:"id"`
	Title       string     `gorm:"not null" json:"title"`
	Description string     `json:"description"`
	ImagePath   string     `gorm:"not null" json:"image_path"`
	UserID      uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
	User        User       `gorm:"foreignKey:UserID" json:"-"`
	FolderID    uuid.UUID  `gorm:"type:uuid;not null" json:"folder_id"`
	Folder      Folder     `gorm:"foreignKey:FolderID" json:"-"`
	VoiceID     *uuid.UUID `gorm:"type:uuid" json:"voice_id,omitempty"`
	Voice       *Voice     `gorm:"foreignKey:VoiceID" json:"voice,omitempty"`
	Comments    []Comment  `gorm:"foreignKey:MemoryID" json:"comments,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Date        time.Time  `json:"date"` 
}

func (m *Memory) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.New()
	return nil
}


