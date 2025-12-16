package model

import(
	"gorm.io/gorm"
)

type Access struct {
    gorm.Model
    FileID string `gorm:"not null;index:idx_file_user,unique"`
    UserID string `gorm:"not null;index:idx_file_user,unique"`
}