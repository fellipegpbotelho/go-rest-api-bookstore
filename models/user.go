package models

import (
	_ "github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID uint `json:"id" gorm:"primary_key"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Age uint8 `json:"age"`
	Email string `json:"email"`
	Password string `json:"password" gorm:"unique"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
