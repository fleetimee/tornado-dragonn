package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username        string   `json:"username" gorm:"unique;not null"`
	Password        string   `json:"password" gorm:"not null"`
	Email           string   `json:"email" gorm:"unique;not null"`
	FirstName       string   `json:"first_name" gorm:"not null"`
	LastName        string   `json:"last_name" gorm:"not null"`
	IsEmailVerified bool     `json:"is_email_verified" gorm:"default:false"`
	Lesson          []Lesson `json:"lesson" gorm:"foreignKey:UserID"`
}
