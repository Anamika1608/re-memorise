package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Folder struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	UserID      uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID" json:"-"`
	ShareCode   string    `gorm:"unique" json:"share_code"`
	IsPublic    bool      `gorm:"default:false" json:"is_public"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Memories    []Memory  `gorm:"foreignKey:FolderID" json:"memories,omitempty"`
}

func (f *Folder) BeforeCreate(tx *gorm.DB) error {
	f.ID = uuid.New()
	return nil
}