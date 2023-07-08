package entity

import (
	"gorm.io/gorm"
)

type User struct {
	ID       string    `gorm:"type:varchar(50);primary_key:true"`
	Name     string    `gorm:"type:varchar(100)"`
	Email    string    `gorm:"type:varchar(100);unique"`
	Password string    `gorm:"type:varchar(100)"`
	gorm.Model
}

func (c *User) TableName() string {
	return "users"
}
