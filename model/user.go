package model

import (
	"time"

	"gorm.io/gorm"
)

//User struct
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"unique;not null;size:120" json:"username"`
	Password  string         `gorm:"not null;size:120" json:"password"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
