package entities

import (
	"gorm.io/gorm"
)

type Lesson struct {
	gorm.Model
	Name   string `json:"name" gorm:"not null"`
	UserID uint   `json:"user_id" gorm:"not null"`
	User   User   `json:"user" gorm:"foreignKey:UserID"`
}
