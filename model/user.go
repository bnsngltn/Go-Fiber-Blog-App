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

//UserJSON struct
//Used to show the model of a user in endpoints
//Also used in swagger documentation
//Yes I am aware of the swagger ignore tag
type UserJSON struct {
	ID       int    `default:"2" json:"id"`
	Username string `json:"username"`
}

//JSON func
//Hides the fields I don't want to be shown in responses.
func (u User) JSON() UserJSON {
	return UserJSON{
		ID:       int(u.ID),
		Username: u.Username,
	}
}

//CreateUserStruct struct
//Used for swagger docs only
type CreateUserStruct struct {
	Username string `default:"admin" json:"username"`
	Password string `default:"123456" json:"password"`
}
