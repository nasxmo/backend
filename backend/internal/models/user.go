package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	OryID      string `gorm:"uniqueIndex"`
	Email      string `gorm:"uniqueIndex"`
	FirstName  string
	LastName   string
	ProfilePic string
}

func (u *User) TableName() string {
	return "users"
}