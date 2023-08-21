package entitys

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	UserName string
	PassWord string
	Name     string
	Email    *string
	Age      uint8
	Birthday *time.Time
}
