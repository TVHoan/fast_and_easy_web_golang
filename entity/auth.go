package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Name     string
	Email    *string
	Age      uint8
	Birthday *time.Time
}
