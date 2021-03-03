package model

import (
	"time"

	"gorm.io/gorm"
)

//Post struct
type Post struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserID    int    `gorm:"not null" json:"user_id"`
	Title     string `gorm:"unique;not null;size:120" json:"title"`
	Body      string `gorm:"not null" json:"body"`
	CreatedAt time.Time
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
