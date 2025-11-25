package model

import(
	"gorm.io/gorm"
)

type Access struct{
	gorm.Model
	FileID string `gorm:"not null"`
	UserID string `gorm:"not null"`
}