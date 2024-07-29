package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id         uint   `gorm:"primaryKey"`
	Email      string `gorm:"unique"`
	Password   string
	Name       string
	LastName   string
	IsVerified bool
}

type UserToken struct {
	gorm.Model
	Id     uint   `gorm:"primaryKey"`
	UserId string `gorm:"unique"`
	Token  string
}
